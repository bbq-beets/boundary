// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authtoken

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hashicorp/boundary/internal/auth/password"
	"github.com/hashicorp/boundary/internal/authtoken/store"
	"github.com/hashicorp/boundary/internal/db"
	"github.com/hashicorp/boundary/internal/db/timestamp"
	"github.com/hashicorp/boundary/internal/iam"
	"github.com/hashicorp/boundary/internal/kms"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

// This file contains tests for methods defined in authtoken.go as well as tests which exercise the db
// functionality directly without going through the respository.  Repository centric tests should be
// placed in repository_test.go

func TestAuthToken_DbUpdate(t *testing.T) {
	conn, _ := db.TestSetup(t, "postgres")
	wrapper := db.TestWrapper(t)
	kms := kms.TestKms(t, conn, wrapper)
	iamRepo := iam.TestRepo(t, conn, wrapper)

	org, _ := iam.TestScopes(t, iamRepo)
	am := password.TestAuthMethods(t, conn, org.GetPublicId(), 1)[0]
	acct := password.TestAccount(t, conn, am.GetPublicId(), "name1")

	newAuthTokId, err := NewAuthTokenId()
	require.NoError(t, err)

	type args struct {
		fieldMask []string
		nullMask  []string
		authTok   *store.AuthToken
	}
	future, err := ptypes.TimestampProto(time.Now().Add(time.Hour))
	require.NoError(t, err)

	tests := []struct {
		name    string
		args    args
		want    *AuthToken
		cnt     int
		wantErr bool
	}{
		{
			name: "immutable-authacctid",
			args: args{
				fieldMask: []string{"AuthAccountId"},
				authTok:   &store.AuthToken{AuthAccountId: acct.GetPublicId()},
			},
			wantErr: true,
		},
		{
			name: "immutable-publicid",
			args: args{
				fieldMask: []string{"PublicId"},
				authTok:   &store.AuthToken{PublicId: newAuthTokId},
			},
			wantErr: true,
		},
		{
			name: "update-last-access-time",
			args: args{
				fieldMask: []string{"ApproximateLastAccessTime"},
				authTok:   &store.AuthToken{ApproximateLastAccessTime: &timestamp.Timestamp{Timestamp: future}},
			},
			cnt: 1,
		},
		{
			name: "nullify-last-access-time",
			args: args{
				nullMask: []string{"ApproximateLastAccessTime"},
				authTok:  &store.AuthToken{},
			},
			cnt: 1,
		},
		{
			name: "update-expiration",
			args: args{
				fieldMask: []string{"ExpirationTime"},
				authTok:   &store.AuthToken{ExpirationTime: &timestamp.Timestamp{Timestamp: future}},
			},
			cnt: 1,
		},
		{
			name: "update-status",
			args: args{
				fieldMask: []string{"Status"},
				authTok:   &store.AuthToken{Status: string(IssuedStatus)},
			},
			cnt: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			w := db.New(conn)

			org, _ := iam.TestScopes(t, iamRepo)
			authTok := TestAuthToken(t, conn, kms, org.GetPublicId())
			proto.Merge(authTok.AuthToken, tt.args.authTok)

			err := authTok.encrypt(context.Background(), wrapper)
			require.NoError(t, err)
			cnt, err := w.Update(context.Background(), authTok, tt.args.fieldMask, tt.args.nullMask)
			if tt.wantErr {
				t.Logf("Got error :%v", err)
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
			assert.Equal(tt.cnt, cnt)
		})
	}
}

func TestAuthToken_DbCreate(t *testing.T) {
	conn, _ := db.TestSetup(t, "postgres")
	rootWrapper := db.TestWrapper(t)
	kms := kms.TestKms(t, conn, rootWrapper)
	org, _ := iam.TestScopes(t, iam.TestRepo(t, conn, rootWrapper))
	wrapper, err := kms.GetWrapper(context.Background(), org.GetPublicId(), 1)
	require.NoError(t, err)
	am := password.TestAuthMethods(t, conn, org.GetPublicId(), 1)[0]
	acct := password.TestAccount(t, conn, am.GetPublicId(), "name1")
	createdAuthToken := TestAuthToken(t, conn, kms, org.GetPublicId())

	testAuthTokenId := func() string {
		id, err := NewAuthTokenId()
		require.NoError(t, err)
		return id
	}

	tests := []struct {
		name      string
		in        *store.AuthToken
		wantError bool
	}{
		{
			name: "basic",
			in: &store.AuthToken{
				PublicId:      testAuthTokenId(),
				Token:         "anything",
				AuthAccountId: acct.GetPublicId(),
			},
		},
		{
			name: "duplicate-id",
			in: &store.AuthToken{
				PublicId:      createdAuthToken.GetPublicId(),
				Token:         "duplicate_id_test",
				AuthAccountId: acct.GetPublicId(),
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			at := &AuthToken{AuthToken: tt.in}
			err := at.encrypt(context.Background(), wrapper)
			require.NoError(t, err)
			err = db.New(conn).Create(context.Background(), at)
			if tt.wantError {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}

func TestAuthToken_DbDelete(t *testing.T) {
	conn, _ := db.TestSetup(t, "postgres")
	testAuthTokenId := func() string {
		id, err := NewAuthTokenId()
		require.NoError(t, err)
		return id
	}

	wrapper := db.TestWrapper(t)
	kms := kms.TestKms(t, conn, wrapper)
	org, _ := iam.TestScopes(t, iam.TestRepo(t, conn, wrapper))
	existingAuthTok := TestAuthToken(t, conn, kms, org.GetPublicId())

	tests := []struct {
		name      string
		at        *AuthToken
		wantError bool
		wantCnt   int
	}{
		{
			name:    "basic",
			at:      &AuthToken{AuthToken: &store.AuthToken{PublicId: existingAuthTok.GetPublicId()}},
			wantCnt: 1,
		},
		{
			name:    "delete-nothing",
			at:      &AuthToken{AuthToken: &store.AuthToken{PublicId: testAuthTokenId()}},
			wantCnt: 0,
		},
		{
			name:      "delete-nil",
			at:        nil,
			wantCnt:   0,
			wantError: true,
		},
		{
			name:      "delete-no-public-id",
			at:        &AuthToken{AuthToken: &store.AuthToken{}},
			wantCnt:   0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			cnt, err := db.New(conn).Delete(context.Background(), tt.at)
			assert.Equal(tt.wantCnt, cnt)
			if tt.wantError {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}
