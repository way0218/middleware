package middleware

import (
	"fmt"
	"testing"
)

func transform(options interface{}, next func(interface{})) {
	u := options.(user)
	fmt.Printf("before type：%T, Name: %s\n", u, u.Name)
	next(options) // 通过验证
}

func validate(options interface{}, next func(interface{})) {
	u := options.(user)
	fmt.Printf("validate type：%T, Title: %s\n", u, u.Title)
	u.Title = "xxxxxxxxx"
	next(u) // 通过验证
}

func send(options interface{}, next func(interface{})) {
	fmt.Printf("send type：%T, value: %s\n", options, options)
	next(nil)
}

type user struct {
	Name  string
	Title string
}

func TestMiddleware(t *testing.T) {
	middleware := New()

	u := user{
		Name:  "Vexth",
		Title: "123",
	}

	middleware.use(transform).use(validate).use(send).execute(u)
}
