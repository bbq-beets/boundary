# ***
# WARNING: Do not EDIT or MERGE this file, it is generated by packagespec.
# ***

LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_ID             := 00-base-aac1d959b4b3af1495941602b5e51fb6181baded
LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_TYPE           := base
LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_BASE_LAYER     := 
LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_SOURCE_INCLUDE := 
LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_SOURCE_EXCLUDE := 
LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_CACHE_KEY_FILE := .buildcache/cache-keys/base-aac1d959b4b3af1495941602b5e51fb6181baded
LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_ARCHIVE_FILE   := .buildcache/archives/00-base-aac1d959b4b3af1495941602b5e51fb6181baded.tar.gz
$(eval $(call LAYER,$(LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_ID),$(LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_TYPE),$(LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_BASE_LAYER),$(LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_SOURCE_INCLUDE),$(LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_SOURCE_EXCLUDE),$(LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_CACHE_KEY_FILE),$(LAYER_00-base-aac1d959b4b3af1495941602b5e51fb6181baded_ARCHIVE_FILE)))

LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_ID             := 01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d
LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_TYPE           := ui
LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_BASE_LAYER     := 00-base-aac1d959b4b3af1495941602b5e51fb6181baded
LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_SOURCE_INCLUDE := internal/ui/VERSION
LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_SOURCE_EXCLUDE := 
LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_CACHE_KEY_FILE := .buildcache/cache-keys/ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d
LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_ARCHIVE_FILE   := .buildcache/archives/01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d.tar.gz
$(eval $(call LAYER,$(LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_ID),$(LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_TYPE),$(LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_BASE_LAYER),$(LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_SOURCE_INCLUDE),$(LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_SOURCE_EXCLUDE),$(LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_CACHE_KEY_FILE),$(LAYER_01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d_ARCHIVE_FILE)))

LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_ID             := 02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e
LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_TYPE           := go-modules
LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_BASE_LAYER     := 01-ui-5f82d1f0cbf979ee051aa1ba0f7105483fa6947d
LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_SOURCE_INCLUDE := go.mod go.sum
LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_SOURCE_EXCLUDE := 
LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_CACHE_KEY_FILE := .buildcache/cache-keys/go-modules-694e8500194e3b187b35bd9a40229f3bd215710e
LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_ARCHIVE_FILE   := .buildcache/archives/02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e.tar.gz
$(eval $(call LAYER,$(LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_ID),$(LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_TYPE),$(LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_BASE_LAYER),$(LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_SOURCE_INCLUDE),$(LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_SOURCE_EXCLUDE),$(LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_CACHE_KEY_FILE),$(LAYER_02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e_ARCHIVE_FILE)))

LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_ID             := 03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3
LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_TYPE           := copy-source
LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_BASE_LAYER     := 02-go-modules-694e8500194e3b187b35bd9a40229f3bd215710e
LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_SOURCE_INCLUDE := *.go
LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_SOURCE_EXCLUDE := 
LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_CACHE_KEY_FILE := .buildcache/cache-keys/copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3
LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_ARCHIVE_FILE   := .buildcache/archives/03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3.tar.gz
$(eval $(call LAYER,$(LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_ID),$(LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_TYPE),$(LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_BASE_LAYER),$(LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_SOURCE_INCLUDE),$(LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_SOURCE_EXCLUDE),$(LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_CACHE_KEY_FILE),$(LAYER_03-copy-source-dde49480ad145a6d7ce0f55c0d23e6f50c04d8d3_ARCHIVE_FILE)))