package sizes

// Size is a type that represents a size in bytes.
type Size uint64

const (
	Byte Size = 1 << (10 * iota)
	Kilobyte
	Megabyte
	Gigabyte
	Terabyte
	Petabyte
	Exabyte
)

// Uint64 returns the Size as an uint64.
func (s Size) Uint64() uint64 {
	return uint64(s)
}

// SizeBuilder is a type that can be used to build a Size.
// It uses the builder pattern to allow for a fluent interface.
// The default size representation is bytes.
type SizeBuilder struct {
	size       Size
	multiplier uint64
}

//
// Shorthand functions to skip the `Builder()` call.
//

// Bytes returns the size in bytes.
func Bytes() *SizeBuilder {
	return Builder().Bytes()
}

// Kilobytes returns the size in kilobytes.
func Kilobytes() *SizeBuilder {
	return Builder().Kilobytes()
}

// Megabytes returns the size in megabytes.
func Megabytes() *SizeBuilder {
	return Builder().Megabytes()
}

// Gigabytes returns the size in gigabytes.
func Gigabytes() *SizeBuilder {
	return Builder().Gigabytes()
}

// Terabytes returns the size in terabytes.
func Terabytes() *SizeBuilder {
	return Builder().Terabytes()
}

// Petabytes returns the size in petabytes.
func Petabytes() *SizeBuilder {
	return Builder().Petabytes()
}

// Exabytes returns the size in exabytes.
func Exabytes() *SizeBuilder {
	return Builder().Exabytes()
}

// Builder returns a new SizeBuilder.
// The default multiplier is 1.
func Builder() *SizeBuilder {
	return &SizeBuilder{
		size: Byte,
		// Multiplier is set to 1 by default.
		multiplier: 1,
	}
}

// Bytes sets the size to be in bytes.
func (s *SizeBuilder) Bytes() *SizeBuilder {
	s.size = Byte
	return s
}

// Kilobytes sets the size to be in kilobytes.
func (s *SizeBuilder) Kilobytes() *SizeBuilder {
	s.size = Kilobyte
	return s
}

// Megabytes sets the size to be in megabytes.
func (s *SizeBuilder) Megabytes() *SizeBuilder {
	s.size = Megabyte
	return s
}

// Gigabytes sets the size to be in gigabytes.
func (s *SizeBuilder) Gigabytes() *SizeBuilder {
	s.size = Gigabyte
	return s
}

// Terabytes sets the size to be in terabytes.
func (s *SizeBuilder) Terabytes() *SizeBuilder {
	s.size = Terabyte
	return s
}

// Petabytes sets the size to be in petabytes.
func (s *SizeBuilder) Petabytes() *SizeBuilder {
	s.size = Petabyte
	return s
}

// Exabytes sets the size to be in exabytes.
func (s *SizeBuilder) Exabytes() *SizeBuilder {
	s.size = Exabyte
	return s
}

// Multiply sets the multiplier for the size.
func (s *SizeBuilder) Multiply(multiplier uint64) *SizeBuilder {
	s.multiplier = multiplier
	return s
}

// Calculate returns the calculated Size as an uint64.
func (s *SizeBuilder) Calculate() uint64 {
	return s.size.Uint64() * s.multiplier
}

// CalculateAs returns the calculated Size in the specified unit.
// Mind that the result is rounded down, as it is an uint64 and not a float64.
func (s *SizeBuilder) CalculateAs(unit Size) uint64 {
	return s.Calculate() / unit.Uint64()
}
