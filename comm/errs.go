package comm

type ClassyErr interface {
	Class() string
	error
}

type classyErr struct {
	class string
	error
}

func (ce *classyError) Class() string {
	return class
}

type ErrClass int

func Classify(class ErrClass, err error) ClassyErr {

}
