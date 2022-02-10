/*
 * @Author: Tperam
 * @Date: 2022-02-10 22:10:02
 * @LastEditTime: 2022-02-10 22:59:46
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\bit_benchmark_test.go
 */
package bitoperand_test

import (
	"bitoperand"
	"sync"
	"testing"
)

func BenchmarkBaseBitSet(b *testing.B) {
	var bi bitoperand.Bit = bitoperand.NewBaseBit(65536)
	BitSet(bi, b.N)
}
func BenchmarkProxyLockBitSet(b *testing.B) {
	var bi bitoperand.Bit = bitoperand.NewProxyLockBit(bitoperand.NewBaseBit(65536))
	BitSet(bi, b.N)
}

func BenchmarkLockBitSet(b *testing.B) {
	var bi bitoperand.Bit = bitoperand.NewLockBit(65536)
	BitSet(bi, b.N)
}

func BenchmarkCASBitSet(b *testing.B) {
	var bi bitoperand.Bit = bitoperand.NewCASBit(65536)
	BitSet(bi, b.N)
}
func BitSet(bi bitoperand.Bit, loop int) {
	for i := 0; i < loop; i++ {
		bi.Set(uint(i) & 65535)
	}
}
func GoBitSet(bi bitoperand.Bit, loop int, goNum int) {
	wg := sync.WaitGroup{}
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			BitSet(bi, loop)
			wg.Done()
		}()
	}
	wg.Wait()
}
func BenchmarkConCurrencyBaseBitSet(b *testing.B) {
	var bi bitoperand.Bit = bitoperand.NewBaseBit(65536)
	GoBitSet(bi, b.N, 16)
}
func BenchmarkConCurrencyProxyLockBitSet(b *testing.B) {
	var bi bitoperand.Bit = bitoperand.NewProxyLockBit(bitoperand.NewBaseBit(65536))
	GoBitSet(bi, b.N, 16)
}
func BenchmarkConCurrencyLockBitSet(b *testing.B) {
	var bi bitoperand.Bit = bitoperand.NewLockBit(65536)
	GoBitSet(bi, b.N, 16)
}

func BenchmarkConCurrencyCASBitSet(b *testing.B) {
	var bi bitoperand.Bit = bitoperand.NewCASBit(65536)
	GoBitSet(bi, b.N, 16)
}
