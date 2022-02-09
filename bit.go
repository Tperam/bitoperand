/*
 * @Author: Tperam
 * @Date: 2022-02-08 23:57:35
 * @LastEditTime: 2022-02-09 22:39:27
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\bit.go
 */
package bitoperand

type Bit interface {
	// 设置位，如果原先为真，则将置为假，原先为假则置为真
	Set(bit uint)
	// 返回该位上值
	Get(bit uint) bool
}
