// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DO NOT EDIT (use 'go test -v -update-expected' instead.)
// See cmd/compile/internal/inline/inlheur/testdata/props/README.txt
// for more information on the format of this file.
// <endfilepreamble>

package returns1

import "unsafe"

// returns.go T_simple_allocmem 20 0 1
// ResultFlags
//   0 ResultIsAllocatedMem
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[2]}
// <endfuncpreamble>
func T_simple_allocmem() *Bar {
	return &Bar{}
}

// returns.go T_allocmem_two_returns 30 0 1
// ResultFlags
//   0 ResultIsAllocatedMem
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[2]}
// <endfuncpreamble>
func T_allocmem_two_returns(x int) *Bar {
	// multiple returns
	if x < 0 {
		return new(Bar)
	} else {
		return &Bar{x: 2}
	}
}

// returns.go T_allocmem_three_returns 45 0 1
// ResultFlags
//   0 ResultIsAllocatedMem
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[2]}
// <endfuncpreamble>
func T_allocmem_three_returns(x int) []*Bar {
	// more multiple returns
	switch x {
	case 10, 11, 12:
		return make([]*Bar, 10)
	case 13:
		fallthrough
	case 15:
		return []*Bar{&Bar{x: 15}}
	}
	return make([]*Bar, 0, 10)
}

// returns.go T_return_nil 64 0 1
// ResultFlags
//   0 ResultAlwaysSameConstant
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[8]}
// <endfuncpreamble>
func T_return_nil() *Bar {
	// simple case: no alloc
	return nil
}

// returns.go T_multi_return_nil 75 0 1
// ResultFlags
//   0 ResultAlwaysSameConstant
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[8]}
// <endfuncpreamble>
func T_multi_return_nil(x, y bool) *Bar {
	if x && y {
		return nil
	}
	return nil
}

// returns.go T_multi_return_nil_anomoly 88 0 1
// ResultFlags
//   0 ResultIsConcreteTypeConvertedToInterface
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[4]}
// <endfuncpreamble>
func T_multi_return_nil_anomoly(x, y bool) Itf {
	if x && y {
		var qnil *Q
		return qnil
	}
	var barnil *Bar
	return barnil
}

// returns.go T_multi_return_some_nil 101 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[0]}
// <endfuncpreamble>
func T_multi_return_some_nil(x, y bool) *Bar {
	if x && y {
		return nil
	} else {
		return &GB
	}
}

// returns.go T_mixed_returns 113 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[0]}
// <endfuncpreamble>
func T_mixed_returns(x int) *Bar {
	// mix of alloc and non-alloc
	if x < 0 {
		return new(Bar)
	} else {
		return &GB
	}
}

// returns.go T_mixed_returns_slice 126 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[0]}
// <endfuncpreamble>
func T_mixed_returns_slice(x int) []*Bar {
	// mix of alloc and non-alloc
	switch x {
	case 10, 11, 12:
		return make([]*Bar, 10)
	case 13:
		fallthrough
	case 15:
		return []*Bar{&Bar{x: 15}}
	}
	ba := [...]*Bar{&GB, &GB}
	return ba[:]
}

// returns.go T_maps_and_channels 149 0 1
// ResultFlags
//   0 ResultNoInfo
//   1 ResultNoInfo
//   2 ResultNoInfo
//   3 ResultAlwaysSameConstant
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[0,0,0,8]}
// <endfuncpreamble>
func T_maps_and_channels(x int, b bool) (bool, map[int]int, chan bool, unsafe.Pointer) {
	// maps and channels
	return b, make(map[int]int), make(chan bool), nil
}

// returns.go T_assignment_to_named_returns 158 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[0,0]}
// <endfuncpreamble>
func T_assignment_to_named_returns(x int) (r1 *uint64, r2 *uint64) {
	// assignments to named returns and then "return" not supported
	r1 = new(uint64)
	if x < 1 {
		*r1 = 2
	}
	r2 = new(uint64)
	return
}

// returns.go T_named_returns_but_return_explicit_values 175 0 1
// ResultFlags
//   0 ResultIsAllocatedMem
//   1 ResultIsAllocatedMem
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[2,2]}
// <endfuncpreamble>
func T_named_returns_but_return_explicit_values(x int) (r1 *uint64, r2 *uint64) {
	// named returns ok if all returns are non-empty
	rx1 := new(uint64)
	if x < 1 {
		*rx1 = 2
	}
	rx2 := new(uint64)
	return rx1, rx2
}

// returns.go T_return_concrete_type_to_itf 191 0 1
// ResultFlags
//   0 ResultIsConcreteTypeConvertedToInterface
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[4]}
// <endfuncpreamble>
func T_return_concrete_type_to_itf(x, y int) Itf {
	return &Bar{}
}

// returns.go T_return_concrete_type_to_itfwith_copy 201 0 1
// ResultFlags
//   0 ResultIsConcreteTypeConvertedToInterface
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[4]}
// <endfuncpreamble>
func T_return_concrete_type_to_itfwith_copy(x, y int) Itf {
	b := &Bar{}
	println("whee")
	return b
}

// returns.go T_return_concrete_type_to_itf_mixed 211 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[0]}
// <endfuncpreamble>
func T_return_concrete_type_to_itf_mixed(x, y int) Itf {
	if x < y {
		b := &Bar{}
		return b
	}
	return nil
}

// returns.go T_return_same_func 225 0 1
// ResultFlags
//   0 ResultAlwaysSameInlinableFunc
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[32]}
// <endfuncpreamble>
func T_return_same_func() func(int) int {
	if G < 10 {
		return foo
	} else {
		return foo
	}
}

// returns.go T_return_different_funcs 237 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[0]}
// <endfuncpreamble>
func T_return_different_funcs() func(int) int {
	if G != 10 {
		return foo
	} else {
		return bar
	}
}

// returns.go T_return_same_closure 255 0 1
// ResultFlags
//   0 ResultAlwaysSameInlinableFunc
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[32]}
// <endfuncpreamble>
// returns.go T_return_same_closure.func1 256 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[0]}
// <endfuncpreamble>
func T_return_same_closure() func(int) int {
	p := func(q int) int { return q }
	if G < 10 {
		return p
	} else {
		return p
	}
}

// returns.go T_return_different_closures 278 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[0]}
// <endfuncpreamble>
// returns.go T_return_different_closures.func1 279 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[0]}
// <endfuncpreamble>
// returns.go T_return_different_closures.func2 283 0 1
// ResultFlags
//   0 ResultAlwaysSameConstant
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[8]}
// <endfuncpreamble>
func T_return_different_closures() func(int) int {
	p := func(q int) int { return q }
	if G < 10 {
		return p
	} else {
		return func(q int) int { return 101 }
	}
}

// returns.go T_return_noninlinable 301 0 1
// ResultFlags
//   0 ResultAlwaysSameFunc
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[16]}
// <endfuncpreamble>
// returns.go T_return_noninlinable.func1 302 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[0]}
// <endfuncpreamble>
// returns.go T_return_noninlinable.func1.1 303 0 1
// <endpropsdump>
// {"Flags":0,"ParamFlags":null,"ResultFlags":[]}
// <endfuncpreamble>
func T_return_noninlinable(x int) func(int) int {
	noti := func(q int) int {
		defer func() {
			println(q + x)
		}()
		return q
	}
	return noti
}

type Bar struct {
	x int
	y string
}

func (b *Bar) Plark() {
}

type Q int

func (q *Q) Plark() {
}

func foo(x int) int { return x }
func bar(x int) int { return -x }

var G int
var GB Bar

type Itf interface {
	Plark()
}
