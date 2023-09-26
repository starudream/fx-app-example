package fake

import (
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

var (
	fake     *gofakeit.Faker
	fakeOnce sync.Once
)

func F() *gofakeit.Faker {
	fakeOnce.Do(func() {
		fake = gofakeit.New(time.Now().UnixNano())
	})
	return fake
}
