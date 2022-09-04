package paasaathai

// The consonant classes
type ConsonantClass int

const (
	UndefinedClass ConsonantClass = 0
	HighClass                     = 1
	MidClass                      = 2
	LowClass                      = 3
)

func (s ConsonantClass) String() string {
	switch s {
	case HighClass:
		return "HighClass"
	case MidClass:
		return "MidClass"
	case LowClass:
		return "LowClass"
	default:
		return "UndefinedClass"
	}
}

// The tones
type Tone int

const (
	UndefinedTone Tone = 0
	LowTone            = 1
	FallingTone        = 2
	HighTone           = 3
	RisingTone         = 4
	MidTone            = 5
)

func (s Tone) String() string {
	switch s {
	case MidTone:
		return "M"
	case LowTone:
		return "L"
	case FallingTone:
		return "F"
	case HighTone:
		return "H"
	case RisingTone:
		return "R"
	default:
		return "?"
	}
}
