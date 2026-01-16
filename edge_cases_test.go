package sizes

import (
	"math"
	"testing"
)

// =============================================================================
// EDGE CASES AND BUG ANALYSIS
// =============================================================================
// This file contains tests for edge cases and documents potential issues
// that library users should be aware of.

// =============================================================================
// BUG #1: Silent Overflow in Large Unit Calculations
// =============================================================================
// When calculating sizes that exceed uint64 max (18,446,744,073,709,551,615),
// the result silently wraps around due to unsigned integer overflow.
//
// Max exabytes that fit in uint64: 15 (15 * 2^60 = 17,293,822,569,102,704,640)
// 16 exabytes = 18,446,744,073,709,551,616 which overflows to 0

func TestExabytesOverflow(t *testing.T) {
	tests := []struct {
		name     string
		amount   uint64
		expected uint64
		overflow bool
	}{
		{
			name:     "15 exabytes - maximum safe value",
			amount:   15,
			expected: 15 * 1_152_921_504_606_846_976,
			overflow: false,
		},
		{
			name:     "16 exabytes - OVERFLOWS silently to 0",
			amount:   16,
			expected: 0, // Bug: silently wraps to 0
			overflow: true,
		},
		{
			name:     "17 exabytes - OVERFLOWS silently",
			amount:   17,
			expected: 1_152_921_504_606_846_976, // Wraps to 1 exabyte
			overflow: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Exabytes[uint64](tt.amount)
			if got != tt.expected {
				t.Errorf("Exabytes(%d) = %v, want %v", tt.amount, got, tt.expected)
			}
			if tt.overflow {
				t.Logf("WARNING: Overflow detected - %d exabytes wraps to %d bytes", tt.amount, got)
			}
		})
	}
}

// Similar overflow can occur with Petabytes
func TestPetabytesOverflow(t *testing.T) {
	// Max petabytes: 18,446,744,073,709,551,615 / 1,125,899,906,842,624 = 16,383
	t.Run("16383 petabytes - maximum safe value", func(t *testing.T) {
		got := Petabytes[uint64](16383)
		expected := uint64(16383) * uint64(Petabyte)
		if got != expected {
			t.Errorf("Petabytes(16383) = %v, want %v", got, expected)
		}
	})

	t.Run("16384 petabytes - OVERFLOWS to 0", func(t *testing.T) {
		got := Petabytes[uint64](16384)
		// 16384 * 2^50 = 2^64 = 0 (overflow wraps to 0)
		expected := uint64(0)
		if got != expected {
			t.Errorf("Petabytes(16384) = %v, want %v (overflow to 0)", got, expected)
		}
		t.Logf("WARNING: Overflow detected - 16384 petabytes wraps to %d bytes", got)
	})

	t.Run("16385 petabytes - OVERFLOWS to 1 PB", func(t *testing.T) {
		got := Petabytes[uint64](16385)
		// Wraps around: (16385 * 2^50) mod 2^64 = 2^50 = 1 PB
		expected := uint64(Petabyte)
		if got != expected {
			t.Errorf("Petabytes(16385) = %v, want %v", got, expected)
		}
		t.Logf("WARNING: Overflow detected - 16385 petabytes wraps to %d bytes (1 PB)", got)
	})
}

// =============================================================================
// BUG #2: Silent Truncation When Using Smaller Generic Types
// =============================================================================
// The generic type parameter allows users to specify return types that may be
// too small to hold the result, causing silent truncation/wrap-around.

func TestGenericTypeTruncation(t *testing.T) {
	t.Run("uint8 truncation - Kilobytes(1) should be 1024, but uint8 max is 255", func(t *testing.T) {
		got := Kilobytes[uint8](1)
		// 1024 % 256 = 0 (truncated)
		expected := uint8(0)
		if got != expected {
			t.Errorf("Kilobytes[uint8](1) = %v, expected truncation to %v", got, expected)
		}
		t.Logf("BUG: Kilobytes[uint8](1) returns %d instead of expected 1024 (truncated)", got)
	})

	t.Run("int8 truncation - Bytes(200) exceeds int8 max of 127", func(t *testing.T) {
		got := Bytes[int8](200)
		// 200 as int8 wraps to -56
		expected := int8(-56)
		if got != expected {
			t.Errorf("Bytes[int8](200) = %v, expected %v", got, expected)
		}
		t.Logf("BUG: Bytes[int8](200) returns %d instead of 200 (signed overflow)", got)
	})

	t.Run("uint16 truncation - Megabytes(1) should be 1048576, but uint16 max is 65535", func(t *testing.T) {
		got := Megabytes[uint16](1)
		// 1048576 % 65536 = 0
		expected := uint16(0)
		if got != expected {
			t.Errorf("Megabytes[uint16](1) = %v, expected truncation to %v", got, expected)
		}
		t.Logf("BUG: Megabytes[uint16](1) returns %d instead of 1048576 (truncated)", got)
	})

	t.Run("int32 overflow at boundary - Gigabytes(2) equals int32 max + 1", func(t *testing.T) {
		got := Gigabytes[int32](2)
		// 2 * 1073741824 = 2147483648 = int32 min (wraps to negative)
		expected := int32(math.MinInt32)
		if got != expected {
			t.Errorf("Gigabytes[int32](2) = %v, expected %v", got, expected)
		}
		t.Logf("BUG: Gigabytes[int32](2) returns %d (negative!) instead of 2147483648", got)
	})
}

// =============================================================================
// EDGE CASE #3: Precision Loss in Unit Conversions
// =============================================================================
// Converting smaller units to larger units uses integer division, which
// loses precision. This is documented but not tested.

func TestPrecisionLossInConversions(t *testing.T) {
	t.Run("1025 bytes as kilobytes loses 1 byte", func(t *testing.T) {
		// 1025 bytes = 1 KB + 1 byte
		// Converting to KB: 1025 / 1024 = 1 (loses the extra byte)
		got := BytesAs[uint64](1025, Kilobyte)
		if got != 1 {
			t.Errorf("BytesAs(1025, Kilobyte) = %v, want 1", got)
		}

		// Converting back: 1 KB = 1024 bytes (lost 1 byte)
		backToBytes := KilobytesAs[uint64](got, Byte)
		if backToBytes != 1024 {
			t.Errorf("KilobytesAs(1, Byte) = %v, want 1024", backToBytes)
		}

		t.Logf("Precision loss: 1025 bytes -> 1 KB -> 1024 bytes (lost 1 byte)")
	})

	t.Run("1023 bytes as kilobytes rounds to 0", func(t *testing.T) {
		got := BytesAs[uint64](1023, Kilobyte)
		if got != 0 {
			t.Errorf("BytesAs(1023, Kilobyte) = %v, want 0", got)
		}
		t.Logf("Edge case: 1023 bytes rounds down to 0 KB")
	})

	t.Run("values just below threshold", func(t *testing.T) {
		// 1 MB - 1 byte = 1048575 bytes
		got := BytesAs[uint64](1048575, Megabyte)
		if got != 0 {
			t.Errorf("BytesAs(1048575, Megabyte) = %v, want 0", got)
		}
		t.Logf("Edge case: 1 MB - 1 byte (1048575 bytes) rounds down to 0 MB")
	})

	t.Run("chained conversion precision loss", func(t *testing.T) {
		// Start with 1.5 MB worth of bytes
		startBytes := uint64(1572864) // 1.5 MB = 1.5 * 1024 * 1024

		// Convert to MB (loses the .5)
		mb := BytesAs[uint64](startBytes, Megabyte)
		if mb != 1 {
			t.Errorf("Expected 1 MB, got %d", mb)
		}

		// Convert to GB (becomes 0)
		gb := MegabytesAs[uint64](mb, Gigabyte)
		if gb != 0 {
			t.Errorf("Expected 0 GB, got %d", gb)
		}

		t.Logf("Chained loss: 1572864 bytes -> 1 MB -> 0 GB")
	})
}

// =============================================================================
// EDGE CASE #4: Boundary Values
// =============================================================================
// Test exact boundary values and values at uint64 limits

func TestBoundaryValues(t *testing.T) {
	t.Run("exact unit boundaries", func(t *testing.T) {
		// Exactly 1 KB
		if BytesAs[uint64](1024, Kilobyte) != 1 {
			t.Error("1024 bytes should be exactly 1 KB")
		}
		// Exactly 1 MB
		if BytesAs[uint64](1048576, Megabyte) != 1 {
			t.Error("1048576 bytes should be exactly 1 MB")
		}
		// Exactly 1 GB
		if BytesAs[uint64](1073741824, Gigabyte) != 1 {
			t.Error("1073741824 bytes should be exactly 1 GB")
		}
	})

	t.Run("max uint64 as bytes", func(t *testing.T) {
		maxUint64 := uint64(math.MaxUint64)
		// Should return the same value
		got := Bytes[uint64](maxUint64)
		if got != maxUint64 {
			t.Errorf("Bytes(MaxUint64) = %v, want %v", got, maxUint64)
		}
	})

	t.Run("max uint64 bytes converted to exabytes", func(t *testing.T) {
		maxUint64 := uint64(math.MaxUint64)
		got := BytesAs[uint64](maxUint64, Exabyte)
		// 18446744073709551615 / 1152921504606846976 = 15
		expected := uint64(15)
		if got != expected {
			t.Errorf("BytesAs(MaxUint64, Exabyte) = %v, want %v", got, expected)
		}
	})
}

// =============================================================================
// EDGE CASE #5: All Supported Generic Types
// =============================================================================
// Verify behavior across all generic types in the AnyInteger constraint

func TestAllGenericTypes(t *testing.T) {
	// Test small safe value that fits in all types
	t.Run("small value fits all types", func(t *testing.T) {
		// 100 bytes should fit in all integer types
		if Bytes[int](100) != 100 {
			t.Error("int failed")
		}
		if Bytes[int8](100) != 100 {
			t.Error("int8 failed")
		}
		if Bytes[int16](100) != 100 {
			t.Error("int16 failed")
		}
		if Bytes[int32](100) != 100 {
			t.Error("int32 failed")
		}
		if Bytes[int64](100) != 100 {
			t.Error("int64 failed")
		}
		if Bytes[uint](100) != 100 {
			t.Error("uint failed")
		}
		if Bytes[uint8](100) != 100 {
			t.Error("uint8 failed")
		}
		if Bytes[uint16](100) != 100 {
			t.Error("uint16 failed")
		}
		if Bytes[uint32](100) != 100 {
			t.Error("uint32 failed")
		}
		if Bytes[uint64](100) != 100 {
			t.Error("uint64 failed")
		}
	})
}

// =============================================================================
// EDGE CASE #6: Zero Values
// =============================================================================
// Comprehensive zero value testing across all functions

func TestZeroValues(t *testing.T) {
	// All basic functions with 0
	if Bytes[uint64](0) != 0 {
		t.Error("Bytes(0) should be 0")
	}
	if Kilobytes[uint64](0) != 0 {
		t.Error("Kilobytes(0) should be 0")
	}
	if Megabytes[uint64](0) != 0 {
		t.Error("Megabytes(0) should be 0")
	}
	if Gigabytes[uint64](0) != 0 {
		t.Error("Gigabytes(0) should be 0")
	}
	if Terabytes[uint64](0) != 0 {
		t.Error("Terabytes(0) should be 0")
	}
	if Petabytes[uint64](0) != 0 {
		t.Error("Petabytes(0) should be 0")
	}
	if Exabytes[uint64](0) != 0 {
		t.Error("Exabytes(0) should be 0")
	}

	// All conversion functions with 0
	if BytesAs[uint64](0, Kilobyte) != 0 {
		t.Error("BytesAs(0, KB) should be 0")
	}
	if KilobytesAs[uint64](0, Megabyte) != 0 {
		t.Error("KilobytesAs(0, MB) should be 0")
	}
	if MegabytesAs[uint64](0, Gigabyte) != 0 {
		t.Error("MegabytesAs(0, GB) should be 0")
	}
	if GigabytesAs[uint64](0, Terabyte) != 0 {
		t.Error("GigabytesAs(0, TB) should be 0")
	}
	if TerabytesAs[uint64](0, Petabyte) != 0 {
		t.Error("TerabytesAs(0, PB) should be 0")
	}
	if PetabytesAs[uint64](0, Exabyte) != 0 {
		t.Error("PetabytesAs(0, EB) should be 0")
	}
	if ExabytesAs[uint64](0, Petabyte) != 0 {
		t.Error("ExabytesAs(0, PB) should be 0")
	}
}

// =============================================================================
// DOCUMENTATION: Size Constant Verification
// =============================================================================
// Verify that all size constants are correctly defined

func TestSizeConstants(t *testing.T) {
	tests := []struct {
		name     string
		size     Size
		expected uint64
	}{
		{"Byte", Byte, 1},
		{"Kilobyte", Kilobyte, 1024},
		{"Megabyte", Megabyte, 1048576},
		{"Gigabyte", Gigabyte, 1073741824},
		{"Terabyte", Terabyte, 1099511627776},
		{"Petabyte", Petabyte, 1125899906842624},
		{"Exabyte", Exabyte, 1152921504606846976},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.size.Uint64() != tt.expected {
				t.Errorf("%s = %v, want %v", tt.name, tt.size.Uint64(), tt.expected)
			}
		})
	}

	// Verify the relationships between units
	t.Run("unit relationships", func(t *testing.T) {
		if Kilobyte != Byte*1024 {
			t.Error("Kilobyte should be 1024 Bytes")
		}
		if Megabyte != Kilobyte*1024 {
			t.Error("Megabyte should be 1024 Kilobytes")
		}
		if Gigabyte != Megabyte*1024 {
			t.Error("Gigabyte should be 1024 Megabytes")
		}
		if Terabyte != Gigabyte*1024 {
			t.Error("Terabyte should be 1024 Gigabytes")
		}
		if Petabyte != Terabyte*1024 {
			t.Error("Petabyte should be 1024 Terabytes")
		}
		if Exabyte != Petabyte*1024 {
			t.Error("Exabyte should be 1024 Petabytes")
		}
	})
}
