package greet

import (
	"time"

	"github.com/flw-cn/go-study/oo/base"
)

type Greet struct {
	*base.Base
	config   Config
	interval time.Duration
}

type Config struct {
	base.BaseConfig `yaml:",inline"`
	Interval        int `json:"Interval" yaml:"Interval" flag:"|1|Interval"`
}

func NewGreet(config Config) *Greet {
	return &Greet{
		Base:   base.NewBase(config.BaseConfig),
		config: config,
	}
}

func (f *Greet) Init() error {
	f.Base.Init()
	f.interval = time.Duration(f.config.Interval) * time.Second
	f.Logger.Printf("初始化 greet.....完成。")
	return nil
}

func (f *Greet) Start() error {
	err := f.Base.Start()
	if err != nil {
		return err
	}

	f.Logger.Printf("启动 greet.....完成。")

	go f.Run()

	return nil
}

func (f *Greet) Run() {
	for {
		select {
		case <-time.After(f.interval):
			f.Logger.Printf("[普通日志] Hello, world!")
			f.Debug("[调试日志] Hello, world!")
		}
	}
}
