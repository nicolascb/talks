package database

type memDatabase struct {
	cache map[string]string
}

func NewMemoryDatabase() Database {
	return &memDatabase{
		cache: make(map[string]string),
	}
}

func (kv *memDatabase) Exists(key string) bool {
	_, ok := kv.cache[key]
	return ok
}

func (kv *memDatabase) Get(key string) string {
	if v, ok := kv.cache[key]; ok {
		return v
	}

	return ""
}

func (kv *memDatabase) Insert(key, value string) error {
	kv.cache[key] = value
	return nil
}
