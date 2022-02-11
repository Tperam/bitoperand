/*
 * @Author: Tperam
 * @Date: 2022-02-09 22:41:26
 * @LastEditTime: 2022-02-11 22:22:26
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\base_bit.go
 */
package bitoperand

type BaseBit struct {
	bits []int64
}

func NewBaseBit(size uint) *BaseBit {
	if size == 0 {
		size = DefaultSize
	}
	return &BaseBit{
		bits: make([]int64, size/64+1),
	}
}

func (b *BaseBit) Set(bits uint) {
	b.bits[bits/64] ^= (1 << (bits & 63))
}
func (b *BaseBit) Get(bits uint) bool {
	return b.bits[bits/64]&(1<<(bits&63)) == (1 << (bits & 63))
}

func (b *BaseBit) GetBackingSliceInt() []int64 {
	return b.bits
}
