/*
 * @Author: Tperam
 * @Date: 2022-02-08 23:57:35
 * @LastEditTime: 2022-02-09 22:28:28
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\bit.go
 */
package bitoperand

type Bit struct {
	bit []int64
}

const DefaultSize = 65536

func NewBit(size uint) *Bit {
	if size == 0 {
		size = DefaultSize
	}
	return &Bit{
		bit: make([]int64, size/64+1),
	}
}

/**
 * @Author: Tperam
 * @description: 单数set导致
 * @param {uint} bit
 * @return {*}
 */
func (b *Bit) Set(bit uint) {
	b.bit[bit/64] ^= (1 << (bit & 63))
}
func (b *Bit) Get(bit uint) bool {
	return b.bit[bit/64]&(1<<(bit&63)) == (1 << (bit & 63))
}
