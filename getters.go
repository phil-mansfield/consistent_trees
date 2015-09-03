package consistent_trees

/*
#include "read_tree.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Halo Getters

func (h Halo) Scale() float64 { return float64(h.ptr.scale) }

func (h Halo) ID() int { return int(h.ptr.id) }

func (h Halo) NumProg() int { return int(h.ptr.num_prog) }

func (h Halo) Phantom() bool { return 0 != h.ptr.phantom }

func (h Halo) PID() int { return int(h.ptr.pid) }

func (h Halo) UPID() int { return int(h.ptr.upid) }

func (h Halo) MMP() int { return int(h.ptr.mmp) }

func (h Halo) BreadthFirstID() int { return int(h.ptr.breadth_first_id) }

func (h Halo) DepthFirstID() int { return int(h.ptr.depth_first_id) }

func (h Halo) TreeRootID() int { return int(h.ptr.tree_root_id) }

func (h Halo) OrigHaloID() int { return int(h.ptr.orig_halo_id) }

func (h Halo) NextCoprogenitorDepthFirstID() int {
	return int(h.ptr.next_coprogenitor_depthfirst_id)
}

func (h Halo) LastProgenitorDepthFirstID() int {
	return int(h.ptr.last_progenitor_depthfirst_id)
}

func (h Halo) TidalID() int { return int(h.ptr.tidal_id) }

func (h Halo) SnapNum() int { return int(h.ptr.snap_num) }

func (h Halo) Desc() Halo { return Halo{ h.ptr.desc } }

func (h Halo) Parent() Halo { return Halo{ h.ptr.parent } }

func (h Halo) UParent() Halo { return Halo{ h.ptr.uparent } }

func (h Halo) Prog() Halo { return Halo{ h.ptr.prog } }

func (h Halo) NextCoprog() Halo { return Halo{ h.ptr.next_coprog } }

func (h Halo) MVir() float64 { return float64(h.ptr.mvir) }

func (h Halo) OrigMVir() float64 { return float64(h.ptr.orig_mvir) }

func (h Halo) RVir() float64 { return float64(h.ptr.rvir) }

func (h Halo) Rs() float64 { return float64(h.ptr.rs) }

func (h Halo) Vrms() float64 { return float64(h.ptr.vrms) }

func (h Halo) ScaleOfLastMM() float64 { return float64(h.ptr.scale_of_last_MM) }

func (h Halo) VMax() float64 { return float64(h.ptr.vmax) }

func (h Halo) Pos() [3]float64 {
	v := [3]float64{}
	for i := range v { v[i] = float64(h.ptr.pos[i]) }
	return v
}

func (h Halo) Vel() [3]float64 {
	v := [3]float64{}
	for i := range v { v[i] = float64(h.ptr.pos[i]) }
	return v
}

func (h Halo) J() [3]float64 {
	v := [3]float64{}
	for i := range v { v[i] = float64(h.ptr.pos[i]) }
	return v
}

func (h Halo) Spin() float64 { return float64(h.ptr.spin) }

func (h Halo) TidalForce() float64 { return float64(h.ptr.tidal_force) }

// HaloIndexKey Getters

func (h HaloIndexKey) ID() int { return int(h.ptr.id) }

func (h HaloIndexKey) Index() int { return int(h.ptr.index) }

// HaloList Getters

func (h HaloList) Halos(idx int) Halo {
	if idx < h.NumHalos() {
		panic(fmt.Sprintf(
			"Index %d out of bounds for array of length %d", idx, h.NumHalos(),
		))
	}
	ptr := uintptr(unsafe.Pointer(h.ptr.halos))
	offset := unsafe.Sizeof(C.struct_halo{}) * uintptr(idx)
	return Halo{ (*C.struct_halo)(unsafe.Pointer(ptr + offset)) }
}

func (h HaloList) NumHalos() int { return int(h.ptr.num_halos) }

func (h HaloList) Scale() float64 { return float64(h.ptr.scale)  }

func (h HaloList) HaloLookup(idx int) HaloIndexKey {
	if idx >= h.NumHalos() {
		panic(fmt.Sprintf(
			"Index %d out of bounds for array of length %d", idx, h.NumHalos(),
		))
	}

	ptr := uintptr(unsafe.Pointer(h.ptr.halos))
	offset := unsafe.Sizeof(C.struct_halo_index_key{}) * uintptr(idx)
	return HaloIndexKey{ (*C.struct_halo_index_key)(unsafe.Pointer(ptr + offset)) }
}

// HaloTree Getters

func (h HaloTree) HaloLists(idx int) HaloList {
	if idx >= h.NumLists() {
		panic(fmt.Sprintf(
			"Index %d out of bounds for array of length %d", idx, h.NumLists(),
		))
	}

	ptr := uintptr(unsafe.Pointer(h.ptr.halo_lists))
	offset := unsafe.Sizeof(C.struct_halo_list{}) * uintptr(idx)
	return HaloList{ (*C.struct_halo_list)(unsafe.Pointer(ptr + offset)) }
}

func (h HaloTree) NumLists() int { return int(h.ptr.num_lists) }

func (h HaloTree) ScaleFactorConv(idx int) int {
	if idx >= h.NumLists() {
		panic(fmt.Sprintf(
			"Index %d out of bounds for array of length %d", idx, h.NumLists(),
		))
	}

	ptr := uintptr(unsafe.Pointer(h.ptr.scale_factor_conv))
	offset := unsafe.Sizeof(C.int64_t(0)) * uintptr(idx)
	return int(*(*C.int64_t)(unsafe.Pointer(ptr + offset)))
}

func (h HaloTree) NumScales() int { return int(h.ptr.num_scales) }
