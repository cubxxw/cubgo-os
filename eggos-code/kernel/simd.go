package kernel

const (
	_CPUID_ECX_XSAVE = 1 << 26
	_CPUID_ECX_AVX   = 1 << 28
	_CPUID_EBX_AVX2  = 1 << 5

	_CPUID_FN_STD = 0x00000001
	_CPUID_FN_EXT = 0x80000001
)

//go:nosplit
func sseInit()

//go:nosplit
func avxInit()

//go:nosplit
func cpuid(fn, cx uint32) (eax, ebx, ecx, edx uint32)

//go:nosplit
func simdInit() {
	sseInit()

	// init for avx
	// first check avx function
	_, _, ecx, _ := cpuid(_CPUID_FN_STD, 0)
	if ecx&_CPUID_ECX_XSAVE == 0 {
		return
	}
	if ecx&_CPUID_ECX_AVX == 0 {
		return
	}
	_, ebx, _, _ := cpuid(0x0007, 0)
	if ebx&_CPUID_EBX_AVX2 == 0 {
		return
	}
	// all check passed, init avx
	avxInit()
}
