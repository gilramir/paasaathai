package paasaathai

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
