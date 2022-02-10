/*
 * @Author: Tperam
 * @Date: 2022-02-09 22:41:26
 * @LastEditTime: 2022-02-10 22:47:26
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\base_bit.go
 */
package bitoperand

type BaseBit struct {
	bit []int64
}

func NewBaseBit(size uint) *BaseBit {
	if size == 0 {
		size = DefaultSize
	}
	return &BaseBit{
		bit: make([]int64, size/64+1),
	}
}

func (b *BaseBit) Set(bit uint) {
	b.bit[bit/64] ^= (1 << (bit & 63))
}
func (b *BaseBit) Get(bit uint) bool {
	return b.bit[bit/64]&(1<<(bit&63)) == (1 << (bit & 63))
}
