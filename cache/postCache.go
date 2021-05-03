package cache

type PostCache interface {
	Set(key string, entity interface{})
	Get(key string) interface{}
}
