/*
 * @Author: Tperam
 * @Date: 2022-02-09 01:14:39
 * @LastEditTime: 2022-02-09 01:20:48
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\bitgroup.go
 */
package bitoperand

type BitGroup struct {
	bitMeaning map[string]uint
	bitRole    map[string]int
	bits       []Bits
}
