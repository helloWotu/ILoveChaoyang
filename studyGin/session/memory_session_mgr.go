package session

import(
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MemorySeesionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

func NewMemorySessionMgr() SessionMgr {
	return &MemorySeesionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
}

func (m *MemorySeesionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (m *MemorySeesionMgr) CreateSession() (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	//用uuid作为sessionId
	id := uuid.NewV4()
	sessionId := id.String()
	//创建session
	session = NewMemorySession(sessionId)
	//存起来
	m.sessionMap[sessionId] = session
	return
}

func (m *MemorySeesionMgr) GetSession(sessionId string) (session Session, err error) {
	value,ok := m.sessionMap[sessionId]
	if ok {
		session = value
		return value,nil
	}else  {
		err = errors.New("获取session 失败")
		return _,err
	}
}


