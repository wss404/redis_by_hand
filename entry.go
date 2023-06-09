package main

import (
	"redis_by_hand/datastructure/hashtable"
	"redis_by_hand/datastructure/zset"
	"redis_by_hand/serialization"
	"unsafe"
)

// Entry 侵入式数据结构, 将hashtable节点结构嵌入到有效载荷数据中
type Entry struct {
	Node    hashtable.HNode
	HeapIdx int32
	Key     string
	Val     string
	Type_   uint32
	ZSet    *zset.ZSet
}

func EntryEq(l *hashtable.HNode, r *hashtable.HNode) bool {
	le := (*Entry)(unsafe.Pointer(l))
	re := (*Entry)(unsafe.Pointer(r))
	return le != nil && re != nil && l.HCode == r.HCode && le.Key == re.Key
}

func EntryKey(h *hashtable.HNode, arg *[]byte) {
	serialization.SerializeStr(
		arg,
		&((*Entry)(unsafe.Pointer(h)).Key),
	)
}
