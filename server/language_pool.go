package server

import (
	"errors"
	"sync"
)

type LanguagePool struct {
	mutex     sync.RWMutex
	languages map[string]*Language
}

func NewLanguagePool() *LanguagePool {
	p := new(LanguagePool)
	p.languages = make(map[string]*Language)
	return p
}

func (lp *LanguagePool) Add(l *Language) error {
	lp.mutex.Lock()
	defer lp.mutex.Unlock()

	if _, ok := lp.languages[l.Name]; ok {
		return errors.New("Language with this name already exists")
	}

	lp.languages[l.Name] = l
	return nil
}

func (lp *LanguagePool) Remove(l *Language) {
	lp.mutex.Lock()
	defer lp.mutex.Unlock()

	delete(lp.languages, l.Name)
}

func (lp *LanguagePool) Get(name string) (*Language, bool) {
	language, ok := lp.languages[name]
	return language, ok
}