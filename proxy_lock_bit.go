/*
 * @Author: Tperam
 * @Date: 2022-02-09 22:42:00
 * @LastEditTime: 2022-02-11 22:26:12
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\proxy_lock_bit.go
 */
package bitoperand

import "sync"

type ProxyLockBit struct {
	lock sync.RWMutex
	bit  Bit
}

// 带锁的bit
// proxy 设计模式
func NewProxyLockBit(b Bit) *ProxyLockBit {
	return &ProxyLockBit{
		bit:  b,
		lock: sync.RWMutex{},
	}
}

func (b *ProxyLockBit) Set(bit uint) {
	b.lock.Lock()
	b.bit.Set(bit)
	b.lock.Unlock()
}

func (b *ProxyLockBit) Get(bit uint) bool {
	b.lock.RLock()
	flag := b.bit.Get(bit)
	b.lock.RUnlock()
	return flag
}
func (b *ProxyLockBit) GetBackingSliceInt() []int64 {
	return b.bit.GetBackingSliceInt()
}
