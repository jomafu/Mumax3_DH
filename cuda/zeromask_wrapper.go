package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/jomafu/Mumax3_DH/cuda/cu"
	"github.com/jomafu/Mumax3_DH/timer"
	"sync"
	"unsafe"
)

// CUDA handle for zeromask kernel
var zeromask_code cu.Function

// Stores the arguments for zeromask kernel invocation
type zeromask_args_t struct {
	arg_dst     unsafe.Pointer
	arg_maskLUT unsafe.Pointer
	arg_regions unsafe.Pointer
	arg_N       int
	argptr      [4]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for zeromask kernel invocation
var zeromask_args zeromask_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	zeromask_args.argptr[0] = unsafe.Pointer(&zeromask_args.arg_dst)
	zeromask_args.argptr[1] = unsafe.Pointer(&zeromask_args.arg_maskLUT)
	zeromask_args.argptr[2] = unsafe.Pointer(&zeromask_args.arg_regions)
	zeromask_args.argptr[3] = unsafe.Pointer(&zeromask_args.arg_N)
}

// Wrapper for zeromask CUDA kernel, asynchronous.
func k_zeromask_async(dst unsafe.Pointer, maskLUT unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("zeromask")
	}

	zeromask_args.Lock()
	defer zeromask_args.Unlock()

	if zeromask_code == 0 {
		zeromask_code = fatbinLoad(zeromask_map, "zeromask")
	}

	zeromask_args.arg_dst = dst
	zeromask_args.arg_maskLUT = maskLUT
	zeromask_args.arg_regions = regions
	zeromask_args.arg_N = N

	args := zeromask_args.argptr[:]
	cu.LaunchKernel(zeromask_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("zeromask")
	}
}

// maps compute capability on PTX code for zeromask kernel.
var zeromask_map = map[int]string{0: "",
	30: zeromask_ptx_30,
	32: zeromask_ptx_32,
	35: zeromask_ptx_35,
	37: zeromask_ptx_37,
	50: zeromask_ptx_50,
	52: zeromask_ptx_52,
	53: zeromask_ptx_53,
	60: zeromask_ptx_60,
	61: zeromask_ptx_61,
	62: zeromask_ptx_62,
	70: zeromask_ptx_70,
	72: zeromask_ptx_72,
	75: zeromask_ptx_75}

// zeromask PTX code for various compute capabilities.
const (
	zeromask_ptx_30 = `
.version 6.5
.target sm_30
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<11>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	cvta.to.global.u64 	%rd7, %rd3;
	ld.global.u8 	%r9, [%rd6];
	mul.wide.u32 	%rd8, %r9, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r10, 0;
	st.global.u32 	[%rd12], %r10;

BB0_3:
	ret;
}


`
	zeromask_ptx_32 = `
.version 6.5
.target sm_32
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<12>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	ld.global.nc.u8 	%rs1, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd8, %r10, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r11, 0;
	st.global.u32 	[%rd12], %r11;

BB0_3:
	ret;
}


`
	zeromask_ptx_35 = `
.version 6.5
.target sm_35
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<12>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	ld.global.nc.u8 	%rs1, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd8, %r10, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r11, 0;
	st.global.u32 	[%rd12], %r11;

BB0_3:
	ret;
}


`
	zeromask_ptx_37 = `
.version 6.5
.target sm_37
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<12>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	ld.global.nc.u8 	%rs1, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd8, %r10, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r11, 0;
	st.global.u32 	[%rd12], %r11;

BB0_3:
	ret;
}


`
	zeromask_ptx_50 = `
.version 6.5
.target sm_50
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<12>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	ld.global.nc.u8 	%rs1, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd8, %r10, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r11, 0;
	st.global.u32 	[%rd12], %r11;

BB0_3:
	ret;
}


`
	zeromask_ptx_52 = `
.version 6.5
.target sm_52
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<12>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	ld.global.nc.u8 	%rs1, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd8, %r10, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r11, 0;
	st.global.u32 	[%rd12], %r11;

BB0_3:
	ret;
}


`
	zeromask_ptx_53 = `
.version 6.5
.target sm_53
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<12>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	ld.global.nc.u8 	%rs1, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd8, %r10, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r11, 0;
	st.global.u32 	[%rd12], %r11;

BB0_3:
	ret;
}


`
	zeromask_ptx_60 = `
.version 6.5
.target sm_60
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<12>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	ld.global.nc.u8 	%rs1, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd8, %r10, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r11, 0;
	st.global.u32 	[%rd12], %r11;

BB0_3:
	ret;
}


`
	zeromask_ptx_61 = `
.version 6.5
.target sm_61
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<12>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	ld.global.nc.u8 	%rs1, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd8, %r10, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r11, 0;
	st.global.u32 	[%rd12], %r11;

BB0_3:
	ret;
}


`
	zeromask_ptx_62 = `
.version 6.5
.target sm_62
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<12>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	ld.global.nc.u8 	%rs1, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd8, %r10, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r11, 0;
	st.global.u32 	[%rd12], %r11;

BB0_3:
	ret;
}


`
	zeromask_ptx_70 = `
.version 6.5
.target sm_70
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<12>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	ld.global.nc.u8 	%rs1, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd8, %r10, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r11, 0;
	st.global.u32 	[%rd12], %r11;

BB0_3:
	ret;
}


`
	zeromask_ptx_72 = `
.version 6.5
.target sm_72
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<12>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	ld.global.nc.u8 	%rs1, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd8, %r10, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r11, 0;
	st.global.u32 	[%rd12], %r11;

BB0_3:
	ret;
}


`
	zeromask_ptx_75 = `
.version 6.5
.target sm_75
.address_size 64

	// .globl	zeromask

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<12>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd6, %rd5, %rd1;
	ld.global.nc.u8 	%rs1, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd8, %r10, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f1, [%rd9];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	mov.u32 	%r11, 0;
	st.global.u32 	[%rd12], %r11;

BB0_3:
	ret;
}


`
)
