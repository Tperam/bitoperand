/*
 * @Author: Tperam
 * @Date: 2022-02-09 01:14:39
 * @LastEditTime: 2022-02-11 22:22:15
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\bitgroup.go
 */
package bitoperand

type BitGroup struct {
	bitMeaning map[string]uint
	bitTable   map[string]int
	bits       []Bit
	bitSize    uint
}

func NewBitGroup(size uint) *BitGroup {

	return &BitGroup{
		bitMeaning: make(map[string]uint, size/64+1),
		bitTable:   make(map[string]int),
		bits:       make([]Bit, 0),
		bitSize:    size,
	}
}
func (bg *BitGroup) SetBitMeaning(key string, index uint) {
	bg.bitMeaning[key] = index
}

func (bg *BitGroup) SetBitTable(tablename string) {
	bg.bitTable[tablename] = len(bg.bits)
	bg.bits = append(bg.bits, NewBaseBit(bg.bitSize))
}

func (bg *BitGroup) SetBit(tablename, key string) {
	index := bg.bitMeaning[key]
	bitsTableIndex := bg.bitTable[tablename]
	bg.bits[bitsTableIndex].Set(index)
}

func (bg *BitGroup) GetBit(tablename, key string) bool {
	index := bg.bitMeaning[key]
	bitsTableIndex := bg.bitTable[tablename]
	return bg.bits[bitsTableIndex].Get(index)
}
