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

// sizeBuilder is a type that can be used to build a Size.
type sizeBuilder struct {
	size       Size
	multiplier uint64
}

//
// Shorthand functions to skip the `Builder()` call.
//

// Bytes returns the size in bytes.
func Bytes(amount uint64) uint64 {
	return newBuilder().
		setSize(Byte).
		setMultiplier(amount).
		calculate()
}

// BytesAs returns the size in bytes as the specified unit.
func BytesAs(amount uint64, unit Size) uint64 {
	return newBuilder().
		setSize(Byte).
		setMultiplier(amount).
		calculateAs(unit)
}

// Kilobytes returns the size in kilobytes.
func Kilobytes(amount uint64) uint64 {
	return newBuilder().
		setSize(Kilobyte).
		setMultiplier(amount).
		calculate()
}

// KilobytesAs returns the size in kilobytes as the specified unit.
func KilobytesAs(amount uint64, unit Size) uint64 {
	return newBuilder().
		setSize(Kilobyte).
		setMultiplier(amount).
		calculateAs(unit)
}

// Megabytes returns the size in megabytes.
func Megabytes(amount uint64) uint64 {
	return newBuilder().
		setSize(Megabyte).
		setMultiplier(amount).
		calculate()
}

// MegabytesAs returns the size in megabytes as the specified unit.
func MegabytesAs(amount uint64, unit Size) uint64 {
	return newBuilder().
		setSize(Megabyte).
		setMultiplier(amount).
		calculateAs(unit)
}

// Gigabytes returns the size in gigabytes.
func Gigabytes(amount uint64) uint64 {
	return newBuilder().
		setSize(Gigabyte).
		setMultiplier(amount).
		calculate()
}

// GigabytesAs returns the size in gigabytes as the specified unit.
func GigabytesAs(amount uint64, unit Size) uint64 {
	return newBuilder().
		setSize(Gigabyte).
		setMultiplier(amount).
		calculateAs(unit)
}

// Terabytes returns the size in terabytes.
func Terabytes(amount uint64) uint64 {
	return newBuilder().
		setSize(Terabyte).
		setMultiplier(amount).
		calculate()
}

// TerabytesAs returns the size in terabytes as the specified unit.
func TerabytesAs(amount uint64, unit Size) uint64 {
	return newBuilder().
		setSize(Terabyte).
		setMultiplier(amount).
		calculateAs(unit)
}

// Petabytes returns the size in petabytes.
func Petabytes(amount uint64) uint64 {
	return newBuilder().
		setSize(Petabyte).
		setMultiplier(amount).
		calculate()
}

// PetabytesAs returns the size in petabytes as the specified unit.
func PetabytesAs(amount uint64, unit Size) uint64 {
	return newBuilder().
		setSize(Petabyte).
		setMultiplier(amount).
		calculateAs(unit)
}

// Exabytes returns the size in exabytes.
func Exabytes(amount uint64) uint64 {
	return newBuilder().
		setSize(Exabyte).
		setMultiplier(amount).
		calculate()
}

// ExabytesAs returns the size in exabytes as the specified unit.
func ExabytesAs(amount uint64, unit Size) uint64 {
	return newBuilder().
		setSize(Exabyte).
		setMultiplier(amount).
		calculateAs(unit)
}

// newBuilder returns a new SizeBuilder.
// The default multiplier is 1.
func newBuilder() *sizeBuilder {
	return &sizeBuilder{}
}

func (s *sizeBuilder) setSize(size Size) *sizeBuilder {
	s.size = size
	return s
}

// setMultiplier sets the multiplier for the size.
func (s *sizeBuilder) setMultiplier(multiplier uint64) *sizeBuilder {
	s.multiplier = multiplier
	return s
}

// calculate returns the calculated Size as an uint64.
func (s *sizeBuilder) calculate() uint64 {
	return s.size.Uint64() * s.multiplier
}

// calculateAs returns the calculated Size in the specified unit. Mind that the result is rounded down, as it is an uint64 and not a float64.
func (s *sizeBuilder) calculateAs(unit Size) uint64 {
	return s.calculate() / unit.Uint64()
}
