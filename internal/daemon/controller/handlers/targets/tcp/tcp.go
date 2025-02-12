// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tcp

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/hashicorp/boundary/internal/daemon/controller/handlers"
	"github.com/hashicorp/boundary/internal/daemon/controller/handlers/targets"
	"github.com/hashicorp/boundary/internal/target"
	"github.com/hashicorp/boundary/internal/target/store"
	"github.com/hashicorp/boundary/internal/target/tcp"
	tcpStore "github.com/hashicorp/boundary/internal/target/tcp/store"
	pb "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/targets"
)

const defaultPortField = "attributes.default_port"

type attribute struct {
	*pb.TcpTargetAttributes
}

func (a *attribute) Options() []target.Option {
	var opts []target.Option
	if a.GetDefaultPort().GetValue() != 0 {
		opts = append(opts, target.WithDefaultPort(a.GetDefaultPort().GetValue()))
	}
	return opts
}

func (a *attribute) Vet() map[string]string {
	badFields := map[string]string{}
	if a.GetDefaultPort() == nil {
		badFields["attributes.default_port"] = "This field is required."
	} else if a.GetDefaultPort().GetValue() == 0 {
		badFields["attributes.default_port"] = "This field cannot be set to zero."
	}
	return badFields
}

func (a *attribute) VetForUpdate(p []string) map[string]string {
	if !handlers.MaskContains(p, defaultPortField) {
		return nil
	}
	badFields := map[string]string{}
	if a.GetDefaultPort() == nil {
		badFields["attributes.default_port"] = "This field is required."
	} else if a.GetDefaultPort().GetValue() == 0 {
		badFields["attributes.default_port"] = "This cannot be set to zero."
	}
	return badFields
}

func newAttribute(m any) targets.Attributes {
	a := &attribute{
		&pb.TcpTargetAttributes{},
	}
	if tcpAttr, ok := m.(*pb.Target_TcpTargetAttributes); ok {
		a.TcpTargetAttributes = tcpAttr.TcpTargetAttributes
	}
	return a
}

func setAttributes(t target.Target, out *pb.Target) error {
	if t == nil {
		return nil
	}

	attrs := &pb.Target_TcpTargetAttributes{
		TcpTargetAttributes: &pb.TcpTargetAttributes{},
	}
	if t.GetDefaultPort() > 0 {
		attrs.TcpTargetAttributes.DefaultPort = &wrappers.UInt32Value{Value: t.GetDefaultPort()}
	}

	out.Attrs = attrs
	return nil
}

func init() {
	var maskManager handlers.MaskManager
	var err error

	if maskManager, err = handlers.NewMaskManager(
		handlers.MaskDestination{&tcpStore.Target{}, &store.TargetAddress{}},
		handlers.MaskSource{&pb.Target{}, &pb.TcpTargetAttributes{}},
	); err != nil {
		panic(err)
	}

	targets.Register(tcp.Subtype, maskManager, newAttribute, setAttributes)
}
