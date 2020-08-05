// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package lenconv

// MToF converts meters to feet.
func MToF(m Meter) Foot { return Foot(m * 3.28084) }

// FToM converts feet to meters.
func FToM(f Foot) Meter { return Meter(f / 3.28084) }

//!-
