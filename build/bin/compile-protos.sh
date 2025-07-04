#!/bin/sh

# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

proto_imports=".:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate:${GOPATH}/src":"${GOPATH}/src/github.com/LABORA-INF-UFG/GUARA-e2t/api"
proto_imports1=".:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate:${GOPATH}/src":"${GOPATH}/src/github.com/LABORA-INF-UFG/GUARA-e2t/api"

go_import_paths="Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types"

# compiling E2AP
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/e2ap/v2/e2ap_commondatatypes.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/e2ap/v2/e2ap_constants.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/e2ap/v2/e2ap_containers.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/e2ap/v2/e2ap_ies.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/e2ap/v2/e2ap_pdu_contents.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/e2ap/v2/e2ap_pdu_descriptions.proto

protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/e2ap/v2/e2ap_commondatatypes.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-commondatatypes/e2ap_commondatatypes.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/e2ap/v2/e2ap_constants.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap_constants/e2ap_constants.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/e2ap/v2/e2ap_containers.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-containers/e2ap_containers.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/e2ap/v2/e2ap_ies.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-ies/e2ap_ies.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/e2ap/v2/e2ap_pdu_contents.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-pdu-contents/e2ap_pdu_contents.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/e2ap/v2/e2ap_pdu_descriptions.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-pdu-descriptions/e2ap_pdu_descriptions.pb.go

# compiling F1AP
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/f1ap/v1/f1ap_commondatatypes.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/f1ap/v1/f1ap_constants.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/f1ap/v1/f1ap_containers.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/f1ap/v1/f1ap_ies.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/f1ap/v1/f1ap_pdu_contents.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/f1ap/v1/f1ap_pdu_descriptions.proto

protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/f1ap/v1/f1ap_commondatatypes.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/f1ap/v1/f1ap_commondatatypes/f1ap_commondatatypes.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/f1ap/v1/f1ap_constants.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/f1ap/v1/f1ap_constants/f1ap_constants.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/f1ap/v1/f1ap_containers.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/f1ap/v1/f1ap_containers/f1ap_containers.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/f1ap/v1/f1ap_ies.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/f1ap/v1/f1ap_ies/f1ap_ies.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/f1ap/v1/f1ap_pdu_contents.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/f1ap/v1/f1ap_pdu_contents/f1ap_pdu_contents.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/f1ap/v1/f1ap_pdu_descriptions.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/f1ap/v1/f1ap_pdu_descriptions/f1ap_pdu_descriptions.pb.go

protoc -I=$proto_imports --gogofaster_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/sim/e2,plugins=grpc:. sim/e2/*.proto

cp -r github.com/LABORA-INF-UFG/GUARA-e2t/* .
rm -rf github.com

# compiling XnAP
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/xnap/v1/xnap_commondatatypes.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/xnap/v1/xnap_constants.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/xnap/v1/xnap_containers.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/xnap/v1/xnap_ies.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/xnap/v1/xnap_pdu_contents.proto
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api --validate_out=lang=go:. --proto_path=api --gogo_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/api:. api/xnap/v1/xnap_pdu_descriptions.proto

protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/xnap/v1/xnap_commondatatypes.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/xnap/v1/xnap-commondatatypes/xnap_commondatatypes.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/xnap/v1/xnap_constants.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/xnap/v1/xnap-constants/xnap_constants.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/xnap/v1/xnap_containers.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/xnap/v1/xnap-containers/xnap_containers.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/xnap/v1/xnap_ies.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/xnap/v1/xnap-ies/xnap_ies.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/xnap/v1/xnap_pdu_contents.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/xnap/v1/xnap-pdu-contents/xnap_pdu_contents.pb.go
protoc -I=$proto_imports1:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=api --go_out=. api/xnap/v1/xnap_pdu_descriptions.proto
protoc-go-inject-tag -input=github.com/LABORA-INF-UFG/GUARA-e2t/api/xnap/v1/xnap-pdu-descriptions/xnap_pdu_descriptions.pb.go

protoc -I=$proto_imports --gogofaster_out=$go_import_paths,import_path=github.com/LABORA-INF-UFG/GUARA-e2t/sim/e2,plugins=grpc:. sim/e2/*.proto

cp -r github.com/LABORA-INF-UFG/GUARA-e2t/* .
rm -rf github.com
