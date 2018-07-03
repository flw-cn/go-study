package base

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Base struct {
	Service
	Logger      Logger
	config      BaseConfig
	debug       bool
	initialized bool
}

type BaseConfig struct {
	Debug   bool   `json:"Debug" yaml:"Debug"`
	LogFile string `json:"LogFile" yaml:"LogFile"`
}

func NewBase(config BaseConfig) *Base {
	return &Base{
		config: config,
		Logger: log.New(ioutil.Discard, "", log.LstdFlags),
	}
}

func (b *Base) Init() error {
	if b.config.LogFile == "" {
	} else if b.config.LogFile == "-" {
		b.Logger.SetOutput(os.Stderr)
	} else {
		file, err := os.Create(b.config.LogFile)
		if err == nil {
			b.Logger.SetOutput(file)
		} else {
			b.Logger.SetOutput(os.Stderr)
			b.Logger.Printf("打开日志文件时遇到错误，日志已禁用。Error: %v", err)
			b.Logger.SetOutput(ioutil.Discard)
		}
	}

	b.initialized = true
	b.Logger.Printf("初始化 base.....完成。")

	return nil
}

func (b *Base) SetLogger(logger Logger) {
	if logger == nil {
		b.Logger = log.New(os.Stderr, "", log.LstdFlags)
	}
	b.Logger = logger
}

func (b *Base) SetLogOutput(w io.Writer) {
	b.Logger.SetOutput(w)
}

func (b *Base) SetDebug(debug bool) {
	b.debug = debug
}

func (b *Base) Debug(args ...interface{}) {
	if !b.debug {
		return
	}

	b.Logger.Print(args...)
}

func (b *Base) Debugf(format string, args ...interface{}) {
	if !b.debug {
		return
	}

	b.Logger.Printf(format, args...)
}

func (b *Base) Start() error {
	if !b.initialized {
		return errors.New("没有初始化，不能启动。")
	}

	b.Logger.Printf("启动 base......完成。")

	return nil
}

func (b *Base) Stop() error {
	return nil
}
