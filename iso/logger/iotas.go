package logger

const (
	//Trace is just about anything
	Trace uint8 = 1 << iota
	//Info indicates important events
	Info
	//Warning is about potentially dangerous events
	Warning
	//Error is about error events
	Error
)
