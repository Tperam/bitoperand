/*
 * @Author: Tperam
 * @Date: 2022-02-09 22:42:00
 * @LastEditTime: 2022-02-09 22:46:31
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \bitoperand\lock_bit.go
 */
package bitoperand

import "sync"

type LockBit struct {
	lock sync.RWMutex
	bit  Bit
}

// 带锁的bit
// proxy 设计模式
func NewLockBit(b Bit) *LockBit {
	return &LockBit{
		bit:  b,
		lock: sync.RWMutex{},
	}
}

func (b *LockBit) Set(bit uint) {
	b.lock.Lock()
	b.bit.Set(bit)
	b.lock.Unlock()
}

func (b *LockBit) Get(bit uint) bool {
	b.lock.RLock()
	flag := b.bit.Get(bit)
	b.lock.RUnlock()
	return flag
}
