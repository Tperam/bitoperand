/*
 * @Author: Tperam
 * @Date: 2022-02-08 23:57:35
 * @LastEditTime: 2022-02-09 01:27:32
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\bit.go
 */
package bitoperand

type Bits struct {
	bit []int64
}

const DefaultSize = 65536

func NewBits(size uint) *Bits {
	if size == 0 {
		size = DefaultSize
	}
	return &Bits{
		bit: make([]int64, size/64+1),
	}
}

/**
 * @Author: Tperam
 * @description: 单数set导致
 * @param {uint} bit
 * @return {*}
 */
func (b *Bits) Set(bit uint) {
	b.bit[bit/64] ^= (1 << (bit & 63))
}
func (b *Bits) Get(bit uint) bool {
	return b.bit[bit/64]&(1<<(bit&63)) == (1 << (bit & 63))
}
