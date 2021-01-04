package team5

import (
	"fmt"

	"github.com/SOMAS2020/SOMAS2020/internal/common/shared"
	"github.com/SOMAS2020/SOMAS2020/pkg/miscutils"
)

type wealthTier int

// Array to keep tracks of CP requests and allocations history
type CPRequestHistory []shared.Resources
type CPAllocationHistory []shared.Resources

type clientConfig struct {
	// Initial non planned foraging
	InitialForageTurns uint

	// Skip forage for x amount of returns if theres no return > 1* multiplier
	SkipForage uint

	// If resources go above this limit we are balling with money
	JBThreshold shared.Resources

	// Middle class:  Middle < Jeff bezos
	MiddleThreshold shared.Resources

	// Poor: Imperial student < Middle
	ImperialThreshold shared.Resources
}

const (
	dying           wealthTier = iota // Sets values = 0
	imperialStudent                   // iota sets the folloing values =1
	middleClass                       // = 2
	jeffBezos                         // = 3
)

func (wt wealthTier) String() string {
	strings := [...]string{"Dying", "Imperial_Student", "Middle_Class", "Jeff_Bezos"}
	if wt >= 0 && int(wt) < len(strings) {
		return strings[wt]
	}
	return fmt.Sprintf("Unkown internal state '%v'", int(wt))
}

// GoString implements GoStringer
func (wt wealthTier) GoString() string {
	return wt.String()
}

// MarshalText implements TextMarshaler
func (wt wealthTier) MarshalText() ([]byte, error) {
	return miscutils.MarshalTextForString(wt.String())
}

// MarshalJSON implements RawMessage
func (wt wealthTier) MarshalJSON() ([]byte, error) {
	return miscutils.MarshalJSONForString(wt.String())
}
