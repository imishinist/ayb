package witticism

import (
	_ "embed"
	"encoding/json"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Witticism struct {
	Text string `json:"text"`
}

type Witticisms []Witticism

var (
	//go:embed witticism.json
	witticismJson []byte
	witticisms    *Witticisms

	once sync.Once
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Get() Witticisms {
	once.Do(func() {
		witticisms = new(Witticisms)
		if err := json.Unmarshal(witticismJson, witticisms); err != nil {
			log.Println(err)
			panic("failed to load witticism.json")
		}
	})

	return *witticisms
}

func (w *Witticisms) Random() Witticism {
	n := rand.Intn(len(*w))
	return (*w)[n]
}
