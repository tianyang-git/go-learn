package singleflight

import (
	"sync"
)

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		// 如果请求正在进行中，则等待
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := new(call)
	// 发起请求前加锁
	c.wg.Add(1)
	// 添加到 g.m，表明 key 已经有对应的请求在处理
	g.m[key] = c
	g.mu.Unlock()

	// 调用 fn，发起请求
	c.val, c.err = fn()
	// 请求结束
	c.wg.Done()

	g.mu.Lock()
	// 更新 g.m
	delete(g.m, key)
	g.mu.Unlock()
	// 返回结果
	return c.val, c.err
}
