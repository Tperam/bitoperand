/*
 * @Author: Tperam
 * @Date: 2022-02-09 01:27:42
 * @LastEditTime: 2022-02-09 01:29:50
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
	b := bitoperand.NewBits(0)
	b.Set(9999)
	t.Log(9999, b.Get(9999))
	b.Set(34)
	t.Log(32, b.Get(32))
	t.Log(34, b.Get(34))
}
