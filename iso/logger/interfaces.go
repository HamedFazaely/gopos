package logger

//Logger reperesnts the logger component
//implementations may output to stdout, stderr or any thing that fits their needs
type Logger interface {
	Log(level uint8, formatstr string, subs ...interface{})
}
