package base

import (
	"io"
)

type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	SetOutput(io.Writer)
}

type Logable interface {
	SetLogger(Logger)
	SetLogOutput(io.Writer)
}

type Debugable interface {
	Logable
	SetDebug(bool)
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
}

type Runable interface {
	Init() error
	Start() error
	Stop() error
}

type Service interface {
	Runable
	Debugable
}
