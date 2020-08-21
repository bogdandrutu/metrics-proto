GOPATH_DIR := $(GOPATH)/src/github.com/bogdandrutu/metrics-proto

# Find all .proto files.
MERGED_PROTO_FILES := $(wildcard merged/proto/*/*/*.proto merged/proto/*/*/*/*.proto)
MERGEDFIXED_PROTO_FILES := $(wildcard mergedfixed/proto/*/*/*.proto mergedfixed/proto/*/*/*/*.proto)
UNMERGED_PROTO_FILES := $(wildcard unmerged/proto/*/*/*.proto unmerged/proto/*/*/*/*.proto)
UNMERGEDFIXED_PROTO_FILES := $(wildcard unmergedfixed/proto/*/*/*.proto unmergedfixed/proto/*/*/*/*.proto)

# Function to execute a command. Note the empty line before endef to make sure each command
# gets executed separately instead of concatenated with previous one.
# Accepts command to execute as first parameter.
define exec-command
$(1)

endef

PROTO_INCLUDES := -I$(GOPATH)/pkg/mod/github.com/gogo/protobuf@v1.3.1/

PROTO_MERGED_DIR := merged/gen
PROTO_MERGEDFIXED_DIR := mergedfixed/gen
PROTO_UNMERGED_DIR := unmerged/gen
PROTO_UNMERGEDFIXED_DIR := unmergedfixed/gen

# Generate ProtoBuf implementation for Go.
.PHONY: gen-merged
gen-merged:
	rm -rf ./$(PROTO_MERGED_DIR)
	mkdir -p ./$(PROTO_MERGED_DIR)
	$(foreach file,$(MERGED_PROTO_FILES),$(call exec-command,protoc -I./ $(PROTO_INCLUDES) --gogofaster_out=plugins=grpc:./merged $(file)))
	mv merged/github.com/bogdandrutu/metrics-proto/merged/gen merged/
	rm -fr merged/github.com

# Generate ProtoBuf implementation for Go.
.PHONY: gen-mergedfixed
gen-mergedfixed:
	rm -rf ./$(PROTO_MERGEDFIXED_DIR)
	mkdir -p ./$(PROTO_MERGEDFIXED_DIR)
	$(foreach file,$(MERGEDFIXED_PROTO_FILES),$(call exec-command,protoc -I./ $(PROTO_INCLUDES) --gogofaster_out=plugins=grpc:./mergedfixed $(file)))
	mv mergedfixed/github.com/bogdandrutu/metrics-proto/mergedfixed/gen mergedfixed/
	rm -fr mergedfixed/github.com

# Generate ProtoBuf implementation for Go.
.PHONY: gen-unmerged
gen-unmerged:
	rm -rf ./$(PROTO_UNMERGED_DIR)
	mkdir -p ./$(PROTO_UNMERGED_DIR)
	$(foreach file,$(UNMERGED_PROTO_FILES),$(call exec-command,protoc -I./ $(PROTO_INCLUDES) --gogofaster_out=plugins=grpc:./unmerged $(file)))
	mv unmerged/github.com/bogdandrutu/metrics-proto/unmerged/gen unmerged/
	rm -fr unmerged/github.com

# Generate ProtoBuf implementation for Go.
.PHONY: gen-unmergedfixed
gen-unmergedfixed:
	rm -rf ./$(PROTO_UNMERGEDFIXED_DIR)
	mkdir -p ./$(PROTO_UNMERGEDFIXED_DIR)
	$(foreach file,$(UNMERGEDFIXED_PROTO_FILES),$(call exec-command,protoc -I./ $(PROTO_INCLUDES) --gogofaster_out=plugins=grpc:./unmergedfixed $(file)))
	mv unmergedfixed/github.com/bogdandrutu/metrics-proto/unmergedfixed/gen unmergedfixed/
	rm -fr unmergedfixed/github.com
