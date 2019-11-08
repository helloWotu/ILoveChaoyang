package session

type Session interface {
	Set(key string, value interface{}) (err error)
	Get(key string) (interface{}, error)
	Del(key string) (err error)
	Save() (err error)
}
