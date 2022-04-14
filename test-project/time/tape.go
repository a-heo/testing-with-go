package poker

import (
	"os"
)

//change file type to os.file to truncate
type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	//truncate empties file
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}
