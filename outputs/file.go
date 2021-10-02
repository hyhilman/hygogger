package outputs

import (
	"errors"
	pqerror "github.com/hyhilman/pq-error"
	"github.com/uknth/writer"
	"io"
	"os"
	"path/filepath"
	"time"
)

type FileOutput struct {
	Path     string
	Rotation time.Duration
	writer   io.Writer
	closer   func() error
}

func (f *FileOutput) Close() error {
	if f.closer != nil {
		return f.closer()
	}

	return pqerror.NewError(errors.New("unable to close not initialize writer"))
}

func (f *FileOutput) Write(msg []byte) (int, error) {
	if f.writer == nil {
		dir := filepath.Dir(f.Path)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return 0, pqerror.NewError(err)
		}

		if f.Rotation <= 0 {
			if w, err := os.OpenFile(f.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600); err != nil {
				return 0, pqerror.NewError(err)
			} else {
				f.writer = w
				f.closer = w.Close
			}
		} else {
			if r, err := writer.NewWriter(f.Path, int(f.Rotation.Seconds())); err != nil {
				return 0, pqerror.NewError(err)
			} else {
				f.writer = r
				f.closer = r.Close
			}
		}
	}

	return f.writer.Write(msg)
}
