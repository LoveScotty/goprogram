package factory

import (
	"fmt"
	"sync"

	"github.com/LoveScotty/goprogram/internal/store"
)

var (
	providerMu  sync.RWMutex
	providerMap = make(map[string]store.Factory) // 生产者map
)

// Register 注册一个生产者
func Register(name string, provider store.Factory) {
	if provider == nil {
		panic("provider is null")
	}
	providerMu.Lock()
	defer providerMu.Unlock()
	_, ok := providerMap[name]
	if ok {
		panic("provider already register")
	}

	providerMap[name] = provider
}

// New 获取生产者实例
func New(name string) (store.Factory, error) {
	providerMu.RLock()
	p, ok := providerMap[name]
	providerMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("unknown provider name: %s", name)
	}

	return p, nil
}
