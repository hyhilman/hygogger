package outputs

type Output interface {
	Write([]byte) (int, error)
	Close() error
}
