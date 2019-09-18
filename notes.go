package notes

import "bytes"

type Archive struct {
	buffer bytes.Buffer
}

func NewArchive() Archive {
	var b bytes.Buffer

	return Archive{b}
}

func (a *Archive) Type(sentence string) {
	a.buffer.WriteString(sentence)
}
