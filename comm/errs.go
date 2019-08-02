package comm

type ClassyErr interface {
	Class() string
	error
}

type ErrClass int

// I'm not going to use all of these
const (
	NONFATAL ErrClass = 1 + iota
	FATAL
	REBOOT
	SHUTDOWN
	CONTINUE
	WARNUSER
)

var errStrings = []string{
	"NONFATAL", "FATAL", "REBOOT", "SHUTDOWN", "CONTINUE", "WARNUSER",
}

func (ec ErrClass) String() string {
	if ec < 1 || ec > ErrClass(len(errStrings)) {
		return "UNCLASSIFIED"
	}
	return errStrings[ec]
}

type classyErr struct {
	class string
	error
}

func (ce *classyErr) Class() string {
	return ce.class
}

func (ce *classyErr) Cause() error {
	return ce.error
}

func (ce *classyErr) Error() string {
	return ce.class + ce.error.Error()
}

func Classify(c ErrClass, err error) ClassyErr {
	return &classyErr{
		class: c.String(),
		error: err,
	}
}
