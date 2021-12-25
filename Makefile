export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE := auto
LDFLAGS := -s -w

os-archs=darwin:amd64 darwin:arm64 freebsd:386 freebsd:amd64 linux:386 linux:amd64 linux:arm linux:arm64 windows:386 windows:amd64 linux:mips64 linux:mips64le linux:mips:softfloat linux:mipsle:softfloat

all: build

build: app

app:
	rm -rf ./release
	mkdir ./release
	@$(foreach n, $(os-archs),\
		os=$(shell echo "$(n)" | cut -d : -f 1);\
		arch=$(shell echo "$(n)" | cut -d : -f 2);\
		target_suffix=$${os}_$${arch};\
		echo "Build $${os}-$${arch}...";\
		env GOOS=$${os} GOARCH=$${arch} go build -ldflags "$(LDFLAGS)" -o ./release/nsq_auth_$${target_suffix};\
		echo "Build $${os}-$${arch} done";\
	)
	@mv ./release/nsq_auth_windows_386 ./release/nsq_auth_windows_386.exe
	@mv ./release/nsq_auth_windows_amd64 ./release/nsq_auth_windows_amd64.exe