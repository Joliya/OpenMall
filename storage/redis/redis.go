package redis

type Config struct {
	Addr         string
	Password     string
	DB           int
	MinIdleConn  int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	PoolSize     int
	PoolTimeout  int
	IsTrace      bool
}
