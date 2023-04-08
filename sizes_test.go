package sizes

import "testing"

func TestSizeBuilderCalculate(t *testing.T) {

	tests := []struct {
		name    string
		builder *SizeBuilder
		want    uint64
	}{
		// Lets test for the default multiplier and seeing that the default returned sizes are in bytes
		{
			name:    "Bytes",
			builder: Builder().Bytes(),
			want:    1,
		},
		{
			name:    "Kilobytes",
			builder: Builder().Kilobytes(),
			want:    1024,
		},
		{
			name:    "Megabytes",
			builder: Builder().Megabytes(),
			want:    1048576,
		},
		{
			name:    "Gigabytes",
			builder: Builder().Gigabytes(),
			want:    1073741824,
		},
		{
			name:    "Terabytes",
			builder: Builder().Terabytes(),
			want:    1099511627776,
		},
		{
			name:    "Petabytes",
			builder: Builder().Petabytes(),
			want:    1125899906842624,
		},
		{
			name:    "Exabytes",
			builder: Builder().Exabytes(),
			want:    1152921504606846976,
		},

		// Now testing to see that the multiplier works
		{
			name:    "Bytes * 2",
			builder: Builder().Bytes().Multiply(2),
			want:    2,
		},
		{
			name:    "Kilobytes * 2",
			builder: Builder().Kilobytes().Multiply(2),
			want:    2048,
		},
		{
			name:    "Megabytes * 2",
			builder: Builder().Megabytes().Multiply(2),
			want:    2097152,
		},
		{
			name:    "Gigabytes * 2",
			builder: Builder().Gigabytes().Multiply(2),
			want:    2147483648,
		},
		{
			name:    "Terabytes * 2",
			builder: Builder().Terabytes().Multiply(2),
			want:    2199023255552,
		},
		{
			name:    "Petabytes * 2",
			builder: Builder().Petabytes().Multiply(2),
			want:    2251799813685248,
		},
		{
			name:    "Exabytes * 2",
			builder: Builder().Exabytes().Multiply(2),
			want:    2305843009213693952,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.builder.Calculate(); got != tt.want {
				t.Errorf("SizeBuilder.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSizeBuilderCalculateAs(t *testing.T) {
	tests := []struct {
		name    string
		builder *SizeBuilder
		unit    Size
		want    uint64
	}{
		{
			name:    "Bytes",
			builder: Builder().Bytes(),
			unit:    Byte,
			want:    1,
		},
		{
			name:    "Kilobytes",
			builder: Builder().Kilobytes(),
			unit:    Kilobyte,
			want:    1,
		},
		{
			name:    "Megabytes",
			builder: Builder().Megabytes(),
			unit:    Megabyte,
			want:    1,
		},
		{
			name:    "Gigabytes",
			builder: Builder().Gigabytes(),
			unit:    Gigabyte,
			want:    1,
		},
		{
			name:    "Terabytes",
			builder: Builder().Terabytes(),
			unit:    Terabyte,
			want:    1,
		},
		{
			name:    "Petabytes",
			builder: Builder().Petabytes(),
			unit:    Petabyte,
			want:    1,
		},
		{
			name:    "Exabytes",
			builder: Builder().Exabytes(),
			unit:    Exabyte,
			want:    1,
		},
		{
			name:    "Bytes * 2",
			builder: Builder().Bytes().Multiply(2),
			unit:    Byte,
			want:    2,
		},
		{
			name:    "Kilobytes * 2",
			builder: Builder().Kilobytes().Multiply(2),
			unit:    Kilobyte,
			want:    2,
		},
		{
			name:    "Megabytes * 2",
			builder: Builder().Megabytes().Multiply(2),
			unit:    Megabyte,
			want:    2,
		},
		{
			name:    "Gigabytes * 2",
			builder: Builder().Gigabytes().Multiply(2),
			unit:    Gigabyte,
			want:    2,
		},
		{
			name:    "Terabytes * 2",
			builder: Builder().Terabytes().Multiply(2),
			unit:    Terabyte,
			want:    2,
		},
		{
			name:    "Petabytes * 2",
			builder: Builder().Petabytes().Multiply(2),
			unit:    Petabyte,
			want:    2,
		},
		{
			name:    "Exabytes * 2",
			builder: Builder().Exabytes().Multiply(2),
			unit:    Exabyte,
			want:    2,
		},
		{
			name:    "1024 bytes to kilobytes",
			builder: Builder().Bytes().Multiply(1024),
			unit:    Kilobyte,
			want:    1,
		},
		{
			name:    "1024 bytes to megabytes",
			builder: Builder().Bytes().Multiply(1024),
			unit:    Megabyte,
			want:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.builder.CalculateAs(tt.unit); got != tt.want {
				t.Errorf("SizeBuilder.CalculateAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShorthandFunctions(t *testing.T) {
	tests := []struct {
		name             string
		shorthandBuilder func() *SizeBuilder
		want             uint64
	}{
		{
			name:             "Bytes",
			shorthandBuilder: Bytes,
			want:             1,
		},
		{
			name:             "Kilobytes",
			shorthandBuilder: Kilobytes,
			want:             1024,
		},
		{
			name:             "Megabytes",
			shorthandBuilder: Megabytes,
			want:             1048576,
		},
		{
			name:             "Gigabytes",
			shorthandBuilder: Gigabytes,
			want:             1073741824,
		},
		{
			name:             "Terabytes",
			shorthandBuilder: Terabytes,
			want:             1099511627776,
		},
		{
			name:             "Petabytes",
			shorthandBuilder: Petabytes,
			want:             1125899906842624,
		},
		{
			name:             "Exabytes",
			shorthandBuilder: Exabytes,
			want:             1152921504606846976,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.shorthandBuilder().Calculate(); got != tt.want {
				t.Errorf("SizeBuilder.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
