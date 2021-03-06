GO_BUILD=go build
GO_GEN=go generate
GO_TEST=go test
GOPATH=$(realpath ../../../..)

# this converts rotate instructions from "ro[lr] <reg>" -> "ro[lr] <reg>, 1" for yasm compatibility
PERL_FIXUP_ROTATE=perl -i -pe 's/(ro[rl]\s+\w{2,3})$$/\1, 1/'

C2GOASM=c2goasm -a -f
CC=clang
C_FLAGS=-target x86_64-unknown-none -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables \
	-fno-exceptions -fno-rtti -O3 -fno-builtin -ffast-math -fno-jump-tables -I_lib
ASM_FLAGS_AVX2=-mavx2 -mfma -mllvm -force-vector-width=32
ASM_FLAGS_SSE4=-msse4

GO_SOURCES  := $(shell find . -path ./_lib -prune -o -name '*.go' -not -name '*_test.go')
ALL_SOURCES := $(shell find . -path ./_lib -prune -o -name '*.go' -name '*.s' -not -name '*_test.go')

INTEL_SOURCES := \
	float64_avx2_amd64.s float64_sse4_amd64.s \
	int64_avx2_amd64.s int64_sse4_amd64.s \
	uint64_avx2_amd64.s uint64_sse4_amd64.s

.PHONEY: assembly

assembly: $(INTEL_SOURCES)

generate: ../bin/tmpl
	../bin/tmpl -i -data=float64.tmpldata type.go.tmpl=float64.go type_amd64.go.tmpl=float64_amd64.go type_noasm.go.tmpl=float64_noasm.go type_test.go.tmpl=float64_test.go
	../bin/tmpl -i -data=float64.tmpldata -d arch=avx2 type_simd_amd64.go.tmpl=float64_avx2_amd64.go
	../bin/tmpl -i -data=float64.tmpldata -d arch=sse4 type_simd_amd64.go.tmpl=float64_sse4_amd64.go
	../bin/tmpl -i -data=int64.tmpldata type.go.tmpl=int64.go type_amd64.go.tmpl=int64_amd64.go type_noasm.go.tmpl=int64_noasm.go type_test.go.tmpl=int64_test.go
	../bin/tmpl -i -data=int64.tmpldata -d arch=avx2 type_simd_amd64.go.tmpl=int64_avx2_amd64.go
	../bin/tmpl -i -data=int64.tmpldata -d arch=sse4 type_simd_amd64.go.tmpl=int64_sse4_amd64.go
	../bin/tmpl -i -data=uint64.tmpldata type.go.tmpl=uint64.go type_amd64.go.tmpl=uint64_amd64.go type_noasm.go.tmpl=uint64_noasm.go type_test.go.tmpl=uint64_test.go
	../bin/tmpl -i -data=uint64.tmpldata -d arch=avx2 type_simd_amd64.go.tmpl=uint64_avx2_amd64.go
	../bin/tmpl -i -data=uint64.tmpldata -d arch=sse4 type_simd_amd64.go.tmpl=uint64_sse4_amd64.go

_lib/float64_avx2.s: _lib/float64.c
	$(CC) -S $(C_FLAGS) $(ASM_FLAGS_AVX2) $^ -o $@ ; $(PERL_FIXUP_ROTATE) $@

_lib/float64_sse4.s: _lib/float64.c
	$(CC) -S $(C_FLAGS) $(ASM_FLAGS_SSE4) $^ -o $@ ; $(PERL_FIXUP_ROTATE) $@

float64_avx2_amd64.s: _lib/float64_avx2.s
	$(C2GOASM) -a -f $^ $@

float64_sse4_amd64.s: _lib/float64_sse4.s
	$(C2GOASM) -a -f $^ $@

_lib/int64_avx2.s: _lib/int64.c
	$(CC) -S $(C_FLAGS) $(ASM_FLAGS_AVX2) $^ -o $@ ; $(PERL_FIXUP_ROTATE) $@

_lib/int64_sse4.s: _lib/int64.c
	$(CC) -S $(C_FLAGS) $(ASM_FLAGS_SSE4) $^ -o $@ ; $(PERL_FIXUP_ROTATE) $@

int64_avx2_amd64.s: _lib/int64_avx2.s
	$(C2GOASM) -a -f $^ $@

int64_sse4_amd64.s: _lib/int64_sse4.s
	$(C2GOASM) -a -f $^ $@

_lib/uint64_avx2.s: _lib/uint64.c
	$(CC) -S $(C_FLAGS) $(ASM_FLAGS_AVX2) $^ -o $@ ; $(PERL_FIXUP_ROTATE) $@

_lib/uint64_sse4.s: _lib/uint64.c
	$(CC) -S $(C_FLAGS) $(ASM_FLAGS_SSE4) $^ -o $@ ; $(PERL_FIXUP_ROTATE) $@

uint64_avx2_amd64.s: _lib/uint64_avx2.s
	$(C2GOASM) -a -f $^ $@

uint64_sse4_amd64.s: _lib/uint64_sse4.s
	$(C2GOASM) -a -f $^ $@

