module github.com/parca-dev/parca-agent

go 1.16

require (
	github.com/Microsoft/go-winio v0.4.17 // indirect
	github.com/alecthomas/kong v0.2.17
	github.com/aquasecurity/libbpfgo v0.2.1-libbpf-0.4.0
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.1.2
	github.com/docker/docker v23.0.1+incompatible
	github.com/go-kit/log v0.2.0
	github.com/gogo/googleapis v1.4.0 // indirect
	github.com/google/pprof v0.0.0-20210609004039-a478d1d731e9
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/ianlancetaylor/demangle v0.0.0-20200824232613-28f6c0f3b639
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/minio/highwayhash v1.0.2
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/oklog/run v1.1.0
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417
	github.com/parca-dev/parca v0.0.0-20210831075758-4d575344697c
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/prometheus v2.5.0+incompatible
	github.com/stretchr/testify v1.7.0
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22
	google.golang.org/grpc v1.41.0
	gotest.tools/v3 v3.0.3 // indirect
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	k8s.io/cri-api v0.22.2
)

replace github.com/prometheus/prometheus => github.com/prometheus/prometheus v1.8.2-0.20201130085533-a6e18916ab40
