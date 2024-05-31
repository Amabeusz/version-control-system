package common

type Set map[string]bool

func (set Set) New() Set {
	return make(map[string]bool, 0)
}

func (set Set) Add(key string) {
	set[key] = true
}

func (set Set) Delete(key string) {
	delete(set, key)
}

func (set Set) Has(v string) bool {
	return set[v]
}
