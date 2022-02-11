/*
 * @Author: Tperam
 * @Date: 2022-02-09 22:42:00
 * @LastEditTime: 2022-02-11 22:22:36
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\lock_bit.go
 */
package bitoperand

import "sync"

type LockBit struct {
	bits  []int64
	locks []sync.RWMutex
}

func NewLockBit(size uint) *LockBit {
	if size == 0 {
		size = DefaultSize
	}
	return &LockBit{
		bits:  make([]int64, size/64+1),
		locks: make([]sync.RWMutex, size/64+1),
	}
}

func (b *LockBit) Set(bits uint) {
	b.locks[bits/64].Lock()
	b.bits[bits/64] ^= (1 << (bits & 63))
	b.locks[bits/64].Unlock()
}
func (b *LockBit) Get(bits uint) bool {
	return b.bits[bits/64]&(1<<(bits&63)) == (1 << (bits & 63))
}

func (b *LockBit) GetBackingSliceInt() []int64 {
	return b.bits
}
