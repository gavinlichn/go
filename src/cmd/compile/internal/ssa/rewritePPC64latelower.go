// Code generated from _gen/PPC64latelower.rules using 'go generate'; DO NOT EDIT.

package ssa

import "internal/buildcfg"
import "cmd/compile/internal/types"

func rewriteValuePPC64latelower(v *Value) bool {
	switch v.Op {
	case OpPPC64AND:
		return rewriteValuePPC64latelower_OpPPC64AND(v)
	case OpPPC64ISEL:
		return rewriteValuePPC64latelower_OpPPC64ISEL(v)
	case OpPPC64RLDICL:
		return rewriteValuePPC64latelower_OpPPC64RLDICL(v)
	case OpPPC64SETBC:
		return rewriteValuePPC64latelower_OpPPC64SETBC(v)
	case OpPPC64SETBCR:
		return rewriteValuePPC64latelower_OpPPC64SETBCR(v)
	}
	return false
}
func rewriteValuePPC64latelower_OpPPC64AND(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (AND <t> x:(MOVDconst [m]) n)
	// cond: t.Size() <= 2
	// result: (Select0 (ANDCCconst [int64(int16(m))] n))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if x.Op != OpPPC64MOVDconst {
				continue
			}
			m := auxIntToInt64(x.AuxInt)
			n := v_1
			if !(t.Size() <= 2) {
				continue
			}
			v.reset(OpSelect0)
			v0 := b.NewValue0(v.Pos, OpPPC64ANDCCconst, types.NewTuple(typ.Int, types.TypeFlags))
			v0.AuxInt = int64ToAuxInt(int64(int16(m)))
			v0.AddArg(n)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (AND x:(MOVDconst [m]) n)
	// cond: isPPC64ValidShiftMask(m)
	// result: (RLDICL [encodePPC64RotateMask(0,m,64)] n)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if x.Op != OpPPC64MOVDconst {
				continue
			}
			m := auxIntToInt64(x.AuxInt)
			n := v_1
			if !(isPPC64ValidShiftMask(m)) {
				continue
			}
			v.reset(OpPPC64RLDICL)
			v.AuxInt = int64ToAuxInt(encodePPC64RotateMask(0, m, 64))
			v.AddArg(n)
			return true
		}
		break
	}
	// match: (AND x:(MOVDconst [m]) n)
	// cond: m != 0 && isPPC64ValidShiftMask(^m)
	// result: (RLDICR [encodePPC64RotateMask(0,m,64)] n)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if x.Op != OpPPC64MOVDconst {
				continue
			}
			m := auxIntToInt64(x.AuxInt)
			n := v_1
			if !(m != 0 && isPPC64ValidShiftMask(^m)) {
				continue
			}
			v.reset(OpPPC64RLDICR)
			v.AuxInt = int64ToAuxInt(encodePPC64RotateMask(0, m, 64))
			v.AddArg(n)
			return true
		}
		break
	}
	// match: (AND <t> x:(MOVDconst [m]) n)
	// cond: t.Size() == 4 && isPPC64WordRotateMask(m)
	// result: (RLWINM [encodePPC64RotateMask(0,m,32)] n)
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if x.Op != OpPPC64MOVDconst {
				continue
			}
			m := auxIntToInt64(x.AuxInt)
			n := v_1
			if !(t.Size() == 4 && isPPC64WordRotateMask(m)) {
				continue
			}
			v.reset(OpPPC64RLWINM)
			v.AuxInt = int64ToAuxInt(encodePPC64RotateMask(0, m, 32))
			v.AddArg(n)
			return true
		}
		break
	}
	return false
}
func rewriteValuePPC64latelower_OpPPC64ISEL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ISEL [a] x (MOVDconst [0]) z)
	// result: (ISELZ [a] x z)
	for {
		a := auxIntToInt32(v.AuxInt)
		x := v_0
		if v_1.Op != OpPPC64MOVDconst || auxIntToInt64(v_1.AuxInt) != 0 {
			break
		}
		z := v_2
		v.reset(OpPPC64ISELZ)
		v.AuxInt = int32ToAuxInt(a)
		v.AddArg2(x, z)
		return true
	}
	// match: (ISEL [a] (MOVDconst [0]) y z)
	// result: (ISELZ [a^0x4] y z)
	for {
		a := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpPPC64MOVDconst || auxIntToInt64(v_0.AuxInt) != 0 {
			break
		}
		y := v_1
		z := v_2
		v.reset(OpPPC64ISELZ)
		v.AuxInt = int32ToAuxInt(a ^ 0x4)
		v.AddArg2(y, z)
		return true
	}
	return false
}
func rewriteValuePPC64latelower_OpPPC64RLDICL(v *Value) bool {
	v_0 := v.Args[0]
	// match: (RLDICL [em] x:(SRDconst [s] a))
	// cond: (em&0xFF0000)==0
	// result: (RLDICL [mergePPC64RLDICLandSRDconst(em, s)] a)
	for {
		em := auxIntToInt64(v.AuxInt)
		x := v_0
		if x.Op != OpPPC64SRDconst {
			break
		}
		s := auxIntToInt64(x.AuxInt)
		a := x.Args[0]
		if !((em & 0xFF0000) == 0) {
			break
		}
		v.reset(OpPPC64RLDICL)
		v.AuxInt = int64ToAuxInt(mergePPC64RLDICLandSRDconst(em, s))
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValuePPC64latelower_OpPPC64SETBC(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETBC [2] cmp)
	// cond: buildcfg.GOPPC64 <= 9
	// result: (ISELZ [2] (MOVDconst [1]) cmp)
	for {
		if auxIntToInt32(v.AuxInt) != 2 {
			break
		}
		cmp := v_0
		if !(buildcfg.GOPPC64 <= 9) {
			break
		}
		v.reset(OpPPC64ISELZ)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = int64ToAuxInt(1)
		v.AddArg2(v0, cmp)
		return true
	}
	// match: (SETBC [0] cmp)
	// cond: buildcfg.GOPPC64 <= 9
	// result: (ISELZ [0] (MOVDconst [1]) cmp)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		cmp := v_0
		if !(buildcfg.GOPPC64 <= 9) {
			break
		}
		v.reset(OpPPC64ISELZ)
		v.AuxInt = int32ToAuxInt(0)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = int64ToAuxInt(1)
		v.AddArg2(v0, cmp)
		return true
	}
	// match: (SETBC [1] cmp)
	// cond: buildcfg.GOPPC64 <= 9
	// result: (ISELZ [1] (MOVDconst [1]) cmp)
	for {
		if auxIntToInt32(v.AuxInt) != 1 {
			break
		}
		cmp := v_0
		if !(buildcfg.GOPPC64 <= 9) {
			break
		}
		v.reset(OpPPC64ISELZ)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = int64ToAuxInt(1)
		v.AddArg2(v0, cmp)
		return true
	}
	return false
}
func rewriteValuePPC64latelower_OpPPC64SETBCR(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SETBCR [2] cmp)
	// cond: buildcfg.GOPPC64 <= 9
	// result: (ISELZ [6] (MOVDconst [1]) cmp)
	for {
		if auxIntToInt32(v.AuxInt) != 2 {
			break
		}
		cmp := v_0
		if !(buildcfg.GOPPC64 <= 9) {
			break
		}
		v.reset(OpPPC64ISELZ)
		v.AuxInt = int32ToAuxInt(6)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = int64ToAuxInt(1)
		v.AddArg2(v0, cmp)
		return true
	}
	// match: (SETBCR [0] cmp)
	// cond: buildcfg.GOPPC64 <= 9
	// result: (ISELZ [4] (MOVDconst [1]) cmp)
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		cmp := v_0
		if !(buildcfg.GOPPC64 <= 9) {
			break
		}
		v.reset(OpPPC64ISELZ)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = int64ToAuxInt(1)
		v.AddArg2(v0, cmp)
		return true
	}
	// match: (SETBCR [1] cmp)
	// cond: buildcfg.GOPPC64 <= 9
	// result: (ISELZ [5] (MOVDconst [1]) cmp)
	for {
		if auxIntToInt32(v.AuxInt) != 1 {
			break
		}
		cmp := v_0
		if !(buildcfg.GOPPC64 <= 9) {
			break
		}
		v.reset(OpPPC64ISELZ)
		v.AuxInt = int32ToAuxInt(5)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = int64ToAuxInt(1)
		v.AddArg2(v0, cmp)
		return true
	}
	return false
}
func rewriteBlockPPC64latelower(b *Block) bool {
	return false
}
