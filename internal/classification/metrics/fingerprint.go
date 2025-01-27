package metrics

import (
	"fmt"
	"math"

	"github.com/baking-bad/bcdhub/internal/bcd/consts"
	"github.com/baking-bad/bcdhub/internal/models/contract"
)

// Fingerprint -
type Fingerprint struct {
	Section string
}

// NewFingerprint -
func NewFingerprint(section string) *Fingerprint {
	return &Fingerprint{
		Section: section,
	}
}

// Compute -
func (m *Fingerprint) Compute(a, b contract.Script) Feature {
	f := Feature{
		Name: fmt.Sprintf("fingerprint_%s", m.Section),
	}

	var x, y []byte
	switch m.Section {
	case consts.PARAMETER:
		x = a.FingerprintParameter
		y = b.FingerprintParameter
	case consts.STORAGE:
		x = a.FingerprintStorage
		y = b.FingerprintStorage
	case consts.CODE:
		x = a.FingerprintCode
		y = b.FingerprintCode
	default:
		return f
	}

	dist := float64(distance(x, y))
	maxLen := math.Max(float64(len(x)), float64(len(y)))
	val := 1. - math.Pow(dist/maxLen, 1.25)
	f.Value = round(val)
	return f
}

func distance(a, b []byte) int {
	if len(a) == 0 {
		return len(b)
	}

	if len(b) == 0 {
		return len(a)
	}

	if len(a) == len(b) {
		eq := true
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				eq = false
				break
			}
		}
		if eq {
			return 0
		}
	}

	// swap to save some memory O(min(a,b)) instead of O(a)
	if len(a) > len(b) {
		a, b = b, a
	}

	lenA := len(a)
	lenB := len(b)

	x := make([]uint16, lenA+1)
	for i := 1; i < len(x); i++ {
		x[i] = uint16(i)
	}

	// make a dummy bounds check to prevent the 2 bounds check down below.
	// The one inside the loop is particularly costly.
	_ = x[lenA]

	// fill in the rest
	for i := 1; i <= lenB; i++ {
		prev := uint16(i)
		var current uint16

		// log.Printf("b[%d] = %s", i-1, string(b[i-1]))
		for j := 1; j <= lenA; j++ {
			// log.Printf("a[%d] = %s", j-1, string(a[j-1]))
			if b[i-1] == a[j-1] {
				current = x[j-1] // match
			} else {
				current = min(x[j-1]+1, prev+1, x[j]+1)
			}
			x[j-1] = prev
			prev = current
		}
		x[lenA] = prev
	}
	return int(x[lenA])
}

func min(a, b, c uint16) uint16 {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}

// FingerprintLength -
type FingerprintLength struct {
	Section string
}

// NewFingerprintLength -
func NewFingerprintLength(section string) *FingerprintLength {
	return &FingerprintLength{
		Section: section,
	}
}

// Compute -
func (m *FingerprintLength) Compute(a, b contract.Script) Feature {
	f := Feature{
		Name: fmt.Sprintf("fingerprint_length_%s", m.Section),
	}

	var x, y []byte
	switch m.Section {
	case consts.PARAMETER:
		x = a.FingerprintParameter
		y = b.FingerprintParameter
	case consts.STORAGE:
		x = a.FingerprintStorage
		y = b.FingerprintStorage
	case consts.CODE:
		x = a.FingerprintCode
		y = b.FingerprintCode
	default:
		return f
	}

	lx := float64(len(x))
	ly := float64(len(y))
	sum := math.Min(lx, ly) / math.Max(lx, ly)
	f.Value = round(sum)
	return f
}
