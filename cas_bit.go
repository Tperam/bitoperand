/*
 * @Author: Tperam
 * @Date: 2022-02-10 22:13:34
 * @LastEditTime: 2022-02-10 22:55:04
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\cas_bit.go
 */
package bitoperand

import "sync/atomic"

type CASBit struct {
	bit []int64
}

func NewCASBit(size uint) *CASBit {
	if size == 0 {
		size = DefaultSize
	}
	return &CASBit{
		bit: make([]int64, size/64+1),
	}
}

func (b *CASBit) Set(bit uint) {
	tmp := (1 << (bit & 63))
	for !atomic.CompareAndSwapInt64(&b.bit[bit/64], b.bit[bit/64], b.bit[bit/64]^int64(tmp)) {
	}
}
func (b *CASBit) Get(bit uint) bool {
	return b.bit[bit/64]&(1<<(bit&63)) == (1 << (bit & 63))
}
