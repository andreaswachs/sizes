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

type AnyInteger interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Bytes returns the size in bytes.
func Bytes[T AnyInteger](amount uint64) T {
	return T(newBuilder().
		setSize(Byte).
		setMultiplier(amount).
		calculate())
}

// BytesAs returns the size in bytes as the specified type.
func BytesAs[T AnyInteger](amount uint64, unit Size) T {
	return T(newBuilder().
		setSize(Byte).
		setMultiplier(amount).
		calculateAs(unit))
}

// Kilobytes returns the size in kilobytes.
func Kilobytes[T AnyInteger](amount uint64) T {
	return T(newBuilder().
		setSize(Kilobyte).
		setMultiplier(amount).
		calculate())
}

// KilobytesAs returns the size in kilobytes as the specified type.
func KilobytesAs[T AnyInteger](amount uint64, unit Size) T {
	return T(newBuilder().
		setSize(Kilobyte).
		setMultiplier(amount).
		calculateAs(unit))
}

// Megabytes returns the size in megabytes.
func Megabytes[T AnyInteger](amount uint64) T {
	return T(newBuilder().
		setSize(Megabyte).
		setMultiplier(amount).
		calculate())
}

// MegabytesAs returns the size in megabytes as the specified type.
func MegabytesAs[T AnyInteger](amount uint64, unit Size) T {
	return T(newBuilder().
		setSize(Megabyte).
		setMultiplier(amount).
		calculateAs(unit))
}

// Gigabytes returns the size in gigabytes.
func Gigabytes[T AnyInteger](amount uint64) T {
	return T(newBuilder().
		setSize(Gigabyte).
		setMultiplier(amount).
		calculate())
}

// GigabytesAs returns the size in gigabytes as the specified type.
func GigabytesAs[T AnyInteger](amount uint64, unit Size) T {
	return T(newBuilder().
		setSize(Gigabyte).
		setMultiplier(amount).
		calculateAs(unit))
}

// Terabytes returns the size in terabytes.
func Terabytes[T AnyInteger](amount uint64) T {
	return T(newBuilder().
		setSize(Terabyte).
		setMultiplier(amount).
		calculate())
}

// TerabytesAs returns the size in terabytes as the specified type.
func TerabytesAs[T AnyInteger](amount uint64, unit Size) T {
	return T(newBuilder().
		setSize(Terabyte).
		setMultiplier(amount).
		calculateAs(unit))
}

// Petabytes returns the size in petabytes.
func Petabytes[T AnyInteger](amount uint64) T {
	return T(newBuilder().
		setSize(Petabyte).
		setMultiplier(amount).
		calculate())
}

// PetabytesAs returns the size in petabytes as the specified type.
func PetabytesAs[T AnyInteger](amount uint64, unit Size) T {
	return T(newBuilder().
		setSize(Petabyte).
		setMultiplier(amount).
		calculateAs(unit))
}

// Exabytes returns the size in exabytes.
func Exabytes[T AnyInteger](amount uint64) T {
	return T(newBuilder().
		setSize(Exabyte).
		setMultiplier(amount).
		calculate())
}

// ExabytesAs returns the size in exabytes as the specified type.
func ExabytesAs[T AnyInteger](amount uint64, unit Size) T {
	return T(newBuilder().
		setSize(Exabyte).
		setMultiplier(amount).
		calculateAs(unit))
}

// newBuilder returns a new SizeBuilder.
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
