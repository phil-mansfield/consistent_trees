package consistent_trees

/*
#cgo CFLAGS: -m64 -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -O3 -std=c99
#cgo LDFLAGS: -lm
#include <stdlib.h>
#include "read_tree.h"
*/
import "C"

import (
	"unsafe"
)

type Halo struct { ptr *C.struct_halo }
type HaloIndexKey struct { ptr *C.struct_halo_index_key }
type HaloList struct { ptr *C.struct_halo_list }
type HaloTree struct { ptr *C.struct_halo_tree }

func GetHaloTree() HaloTree { return HaloTree{ &C.halo_tree } }
func GetAllHalos() HaloList { return HaloList{ &C.all_halos } }

func LookupHaloInList(hl HaloList, id int) (Halo, bool) {
	h := Halo{ C.lookup_halo_in_list(hl.ptr, C.int64_t(id)) }
	return h, h.ptr != (*C.struct_halo)(nil)
}

func LookupScale(scale float64) (HaloList, bool) {
	hl := HaloList{ C.lookup_scale(C.float(scale)) }
	return hl, hl.ptr != ((*C.struct_halo_list)(nil))
}

func LookupIndex(scale float64) int {
	return int(C.lookup_index(C.float(scale)))
}

func FindClosestScale(scale float64) (HaloList, bool) {
	hl := HaloList{ C.find_closest_scale(C.float(scale)) }
	return hl, hl.ptr != ((*C.struct_halo_list)(nil))
}

func ReadTree(filename string) {
	cStr := C.CString(filename)
	defer C.free(unsafe.Pointer(cStr))
	C.read_tree(cStr)
}

func DeleteTree() {
	C.delete_tree()
}
