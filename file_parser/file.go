package file_parser

type File struct {
	pathName string
}

func (f *File) PathName() string {
	return f.pathName
}
