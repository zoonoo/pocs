package monads

// Maybe may go wrong by not returning a value.
type Maybe interface {
	Bind()
}
