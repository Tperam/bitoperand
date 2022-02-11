/*
 * @Author: Tperam
 * @Date: 2022-02-09 01:27:42
 * @LastEditTime: 2022-02-11 23:10:05
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\bit_test.go
 */
package bitoperand_test

import (
	"bitoperand"
	"testing"
)

func TestB(t *testing.T) {
	b := bitoperand.NewBaseBit(0)
	b.Set(9999)
	t.Log(9999, b.Get(9999))
	b.Set(34)
	t.Log(32, b.Get(32))
	t.Log(34, b.Get(34))
}

func TestBitoperandCompare(t *testing.T) {
	loop := 1000000
	var bi1 bitoperand.Bit = bitoperand.NewBaseBit(65536)
	GoBitSet(bi1, loop, 16)
	var bi2 bitoperand.Bit = bitoperand.NewProxyLockBit(bitoperand.NewBaseBit(65536))
	GoBitSet(bi2, loop, 16)
	var bi3 bitoperand.Bit = bitoperand.NewLockBit(65536)
	GoBitSet(bi3, loop, 16)
	var bi4 bitoperand.Bit = bitoperand.NewCASBit(65536)
	GoBitSet(bi4, loop, 16)
	bi1Int := bi1.GetBackingSliceInt()
	bi2Int := bi2.GetBackingSliceInt()
	bi3Int := bi3.GetBackingSliceInt()
	bi4Int := bi4.GetBackingSliceInt()
	for i := range bi1Int {
		if !(bi1Int[i] == bi2Int[i]) {
			t.Log("b1 与 b2 不相等")
			break
		}
	}
	for i := range bi1Int {
		if !(bi2Int[i] == bi3Int[i]) {
			t.Log("b2 与 b3 不相等")
			break
		}
	}
	for i := range bi1Int {
		if !(bi3Int[i] == bi4Int[i]) {
			t.Log("b3 与 b4 不相等")
			break
		}
	}

}
