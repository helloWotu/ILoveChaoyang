package session

import (
	"errors"
	"sync"
)

type MemorySession struct {
	sessionId string
	//存kv
	data   map[string]interface{}
	rwlock sync.RWMutex
}

//构造函数
func NewMemorySession(id string) Session {
	return &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	//加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	//设置值
	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (interface{}, error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err := errors.New("key not exists in session")
		return _, err
	} else {
		return value, nil
	}
}

func (m *MemorySession) Del(key string) error {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return nil
}

func (MemorySession) Save() error {
	return nil
}
