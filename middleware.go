package middleware

// Middleware .
type Middleware struct {
	Cache       []func(options interface{}, next func(interface{}))
	Middlewares []func(options interface{}, next func(interface{}))
}

// New 初始化Middleware
func New() Middleware {
	return Middleware{}
}

// 注册中间件
func (m Middleware) use(f ...func(options interface{}, next func(interface{}))) Middleware {
	m.Cache = append(m.Cache, f...)
	return m
}

// 每个中间件只有两个形参 第一是传进来的参数 第二个是调用下一个中间件的函数
// 中间件的执行顺序是根据你注册中间件的顺序来去调用的
func (m Middleware) next(params interface{}) {
	if len(m.Middlewares) == 0 {
		return
	}
	ware := m.Middlewares[0]
	m.Middlewares = m.Middlewares[1:]
	ware(params, m.next)
}

func (m Middleware) execute(params interface{}) {
	m.Middlewares = append(m.Middlewares, m.Cache...) // 复制一份
	m.next(params)
}
