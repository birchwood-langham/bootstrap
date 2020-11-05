package service

var Usage string
var ShortDescription string
var LongDescription string

func SetCliProperties(usage, short, long string) {
	Usage = usage
	ShortDescription = short
	LongDescription = long
}
