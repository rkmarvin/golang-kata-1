package domain

type Libraty struct {
	books     []*Book
	magazines []*Magazine
	authors   []*Author
}

func NewLibrary(a []*Author) *Libraty {
	return &Libraty{
		authors: a,
	}
}

func (l *Libraty) Authors() []*Author {
	return l.authors
}
