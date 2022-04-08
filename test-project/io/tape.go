package main

import "io"

type tape struct {
	file io.ReadWriteSeeker
}

//currently writes over it
func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Seek(0, 0)
	return t.file.Write(p)
}
