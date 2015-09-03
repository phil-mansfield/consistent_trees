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

func (hl HaloList) Lookup(id int) Halo {
	return Halo{ C.lookup_halo_in_list(hl.ptr, C.int64_t(id)) }
}

func LookupScale(scale float64) HaloList {
	return HaloList{ C.lookup_scale(C.float(scale)) }
}

func FindClosestScale(scale float64) HaloList {
	return HaloList{ C.find_closest_scale(C.float(scale)) }
}

func ReadTree(filename string) {
	cStr := C.CString(filename)
	defer C.free(unsafe.Pointer(cStr))
	C.read_tree(cStr)
}

func DeleteTree() {
	C.delete_tree()
}
