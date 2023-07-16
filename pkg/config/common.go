package config

type common struct {
	ServerPort int64
}

var Common = &common{}

func (c *common) Init() {
}
