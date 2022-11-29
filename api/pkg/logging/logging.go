package logging

import "os"

var LogFile *os.File

func FilePtr(name string) *os.File {
	LogFile, _ = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	return LogFile
}
