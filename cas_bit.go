/*
 * @Author: Tperam
 * @Date: 2022-02-10 22:13:34
 * @LastEditTime: 2022-02-11 22:23:05
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\cas_bit.go
 */
package bitoperand

import "sync/atomic"

// 实际上，对于高性能模块来说，为了使方法能够内联得到一定的速度提升，建议避开 interface接口使用。
type CASBit struct {
	bits []int64
}

func NewCASBit(size uint) *CASBit {
	if size == 0 {
		size = DefaultSize
	}
	return &CASBit{
		bits: make([]int64, size/64+1),
	}
}

func (b *CASBit) Set(bits uint) {
	tmp := (1 << (bits & 63))
	for !atomic.CompareAndSwapInt64(&b.bits[bits/64], b.bits[bits/64], b.bits[bits/64]^int64(tmp)) {
	}
}
func (b *CASBit) Get(bits uint) bool {
	return b.bits[bits/64]&(1<<(bits&63)) == (1 << (bits & 63))
}

func (b *CASBit) GetBackingSliceInt() []int64 {
	return b.bits
}
