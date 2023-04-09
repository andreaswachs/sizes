package sizes

import "testing"

func TestBytes(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		want   uint64
	}{
		{
			name:   "1 byte",
			amount: 1,
			want:   1,
		},
		{
			name:   "2 bytes",
			amount: 2,
			want:   2,
		},
		{
			name:   "0 bytes",
			amount: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes(tt.amount); got != tt.want {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesAs(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		unit   Size
		want   uint64
	}{
		{
			name:   "1 byte",
			amount: 1,
			unit:   Byte,
			want:   1,
		},
		{
			name:   "0 bytes",
			amount: 0,
			unit:   Byte,
			want:   0,
		},
		{
			name:   "1024 bytes as kilobytes",
			amount: 1024,
			unit:   Kilobyte,
			want:   1,
		},
		{
			name:   "bytes to a Megabyte",
			amount: 1_048_576,
			unit:   Megabyte,
			want:   1,
		},
		{
			name:   "Bytes to a Gigabyte",
			amount: 1_073_741_824,
			unit:   Gigabyte,
			want:   1,
		},
		{
			name:   "Bytes to a Terabyte",
			amount: 1_099_511_627_776,
			unit:   Terabyte,
			want:   1,
		},
		{
			name:   "Bytes to a Petabyte",
			amount: 1_125_899_906_842_624,
			unit:   Petabyte,
			want:   1,
		},
		{
			name:   "Bytes to a Exabyte",
			amount: 1_152_921_504_606_846_976,
			unit:   Exabyte,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesAs(tt.amount, tt.unit); got != tt.want {
				t.Errorf("BytesAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKilobytes(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		want   uint64
	}{
		{
			name:   "1 kilobyte",
			amount: 1,
			want:   1024,
		},
		{
			name:   "2 kilobytes",
			amount: 2,
			want:   2048,
		},
		{
			name:   "0 kilobytes",
			amount: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kilobytes(tt.amount); got != tt.want {
				t.Errorf("Kilobytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKilobytesAs(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		unit   Size
		want   uint64
	}{
		{
			name:   "1 kilobyte",
			amount: 1,
			unit:   Kilobyte,
			want:   1,
		},
		{
			name:   "0 kilobytes",
			amount: 0,
			unit:   Kilobyte,
			want:   0,
		},
		{
			name:   "1 kilobyte as bytes",
			amount: 1,
			unit:   Byte,
			want:   1024,
		},
		{
			name:   "1 kilobyte as megabytes",
			amount: 1,
			unit:   Megabyte,
			want:   0,
		},
		{
			name:   "1 kilobyte as gigabytes",
			amount: 1,
			unit:   Gigabyte,
			want:   0,
		},
		{
			name:   "1 kilobyte as terabytes",
			amount: 1,
			unit:   Terabyte,
			want:   0,
		},
		{
			name:   "1 kilobyte as petabytes",
			amount: 1,
			unit:   Petabyte,
			want:   0,
		},
		{
			name:   "1 kilobyte as exabytes",
			amount: 1,
			unit:   Exabyte,
			want:   0,
		},
		{
			name:   "kilotbytes to a Megabyte",
			amount: 1_024,
			unit:   Megabyte,
			want:   1,
		},
		{
			name:   "kilotbytes to a Gigabyte",
			amount: 1_048_576,
			unit:   Gigabyte,
			want:   1,
		},
		{
			name:   "kilotbytes to a Terabyte",
			amount: 1_073_741_824,
			unit:   Terabyte,
			want:   1,
		},
		{
			name:   "kilotbytes to a Petabyte",
			amount: 1_099_511_627_776,
			unit:   Petabyte,
			want:   1,
		},
		{
			name:   "kilotbytes to a Exabyte",
			amount: 1_125_899_906_842_624,
			unit:   Exabyte,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KilobytesAs(tt.amount, tt.unit); got != tt.want {
				t.Errorf("KilobytesAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMegabytes(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		want   uint64
	}{
		{
			name:   "1 megabyte",
			amount: 1,
			want:   1_048_576,
		},
		{
			name:   "2 megabytes",
			amount: 2,
			want:   2_097_152,
		},
		{
			name:   "0 megabytes",
			amount: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Megabytes(tt.amount); got != tt.want {
				t.Errorf("Megabytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMegabytesAs(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		unit   Size
		want   uint64
	}{
		{
			name:   "1 megabyte",
			amount: 1,
			unit:   Megabyte,
			want:   1,
		},
		{
			name:   "0 megabytes",
			amount: 0,
			unit:   Megabyte,
			want:   0,
		},
		{
			name:   "1 megabyte as bytes",
			amount: 1,
			unit:   Byte,
			want:   1_048_576,
		},
		{
			name:   "1 megabyte as kilobytes",
			amount: 1,
			unit:   Kilobyte,
			want:   1_024,
		},
		{
			name:   "1 megabyte as gigabytes",
			amount: 1,
			unit:   Gigabyte,
			want:   0,
		},
		{
			name:   "1 megabyte as terabytes",
			amount: 1,
			unit:   Terabyte,
			want:   0,
		},
		{
			name:   "1 megabyte as petabytes",
			amount: 1,
			unit:   Petabyte,
			want:   0,
		},
		{
			name:   "1 megabyte as exabytes",
			amount: 1,
			unit:   Exabyte,
			want:   0,
		},
		{
			name:   "megabytes to a Gigabyte",
			amount: 1_048,
			unit:   Gigabyte,
			want:   1,
		},
		{
			name:   "megabytes to a Terabyte",
			amount: 1_048_576,
			unit:   Terabyte,
			want:   1,
		},
		{
			name:   "megabytes to a Petabyte",
			amount: 1_073_741_824,
			unit:   Petabyte,
			want:   1,
		},
		{
			name:   "megabytes to a Exabyte",
			amount: 1_099_511_627_776,
			unit:   Exabyte,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MegabytesAs(tt.amount, tt.unit); got != tt.want {
				t.Errorf("MegabytesAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGigabytes(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		want   uint64
	}{
		{
			name:   "1 gigabyte",
			amount: 1,
			want:   1_073_741_824,
		},
		{
			name:   "2 gigabytes",
			amount: 2,
			want:   2_147_483_648,
		},
		{
			name:   "0 gigabytes",
			amount: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gigabytes(tt.amount); got != tt.want {
				t.Errorf("Gigabytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGigabytesAs(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		unit   Size
		want   uint64
	}{
		{
			name:   "1 gigabyte",
			amount: 1,
			unit:   Gigabyte,
			want:   1,
		},
		{
			name:   "0 gigabytes",
			amount: 0,
			unit:   Gigabyte,
			want:   0,
		},
		{
			name:   "1 gigabyte as bytes",
			amount: 1,
			unit:   Byte,
			want:   1_073_741_824,
		},
		{
			name:   "1 gigabyte as kilobytes",
			amount: 1,
			unit:   Kilobyte,
			want:   1_048_576,
		},
		{
			name:   "1 gigabyte as megabytes",
			amount: 1,
			unit:   Megabyte,
			want:   1_024,
		},
		{
			name:   "1 gigabyte as terabytes",
			amount: 1,
			unit:   Terabyte,
			want:   0,
		},
		{
			name:   "1 gigabyte as petabytes",
			amount: 1,
			unit:   Petabyte,
			want:   0,
		},
		{
			name:   "1 gigabyte as exabytes",
			amount: 1,
			unit:   Exabyte,
			want:   0,
		},
		{
			name:   "gigabytes to a Terabyte",
			amount: 1_024,
			unit:   Terabyte,
			want:   1,
		},
		{
			name:   "gigabytes to a Petabyte",
			amount: 1_048_576,
			unit:   Petabyte,
			want:   1,
		},
		{
			name:   "gigabytes to a Exabyte",
			amount: 1_073_741_824,
			unit:   Exabyte,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GigabytesAs(tt.amount, tt.unit); got != tt.want {
				t.Errorf("GigabytesAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTerabytes(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		want   uint64
	}{
		{
			name:   "1 terabyte",
			amount: 1,
			want:   1_099_511_627_776,
		},
		{
			name:   "2 terabytes",
			amount: 2,
			want:   2_199_023_255_552,
		},
		{
			name:   "0 terabytes",
			amount: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Terabytes(tt.amount); got != tt.want {
				t.Errorf("Terabytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTerabytesAs(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		unit   Size
		want   uint64
	}{
		{
			name:   "1 terabyte",
			amount: 1,
			unit:   Terabyte,
			want:   1,
		},
		{
			name:   "0 terabytes",
			amount: 0,
			unit:   Terabyte,
			want:   0,
		},
		{
			name:   "1 terabyte as bytes",
			amount: 1,
			unit:   Byte,
			want:   1_099_511_627_776,
		},
		{
			name:   "1 terabyte as kilobytes",
			amount: 1,
			unit:   Kilobyte,
			want:   1_073_741_824,
		},
		{
			name:   "1 terabyte as megabytes",
			amount: 1,
			unit:   Megabyte,
			want:   1_048_576,
		},
		{
			name:   "1 terabyte as gigabytes",
			amount: 1,
			unit:   Gigabyte,
			want:   1_024,
		},
		{
			name:   "1 terabyte as petabytes",
			amount: 1,
			unit:   Petabyte,
			want:   0,
		},
		{
			name:   "1 terabyte as exabytes",
			amount: 1,
			unit:   Exabyte,
			want:   0,
		},
		{
			name:   "terabytes to a Petabyte",
			amount: 1_024,
			unit:   Petabyte,
			want:   1,
		},
		{
			name:   "terabytes to a Exabyte",
			amount: 1_048_576,
			unit:   Exabyte,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TerabytesAs(tt.amount, tt.unit); got != tt.want {
				t.Errorf("TerabytesAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPetabytes(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		want   uint64
	}{
		{
			name:   "1 petabyte",
			amount: 1,
			want:   1_125_899_906_842_624,
		},
		{
			name:   "2 petabytes",
			amount: 2,
			want:   2_251_799_813_685_248,
		},
		{
			name:   "0 petabytes",
			amount: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Petabytes(tt.amount); got != tt.want {
				t.Errorf("Petabytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPetabytesAs(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		unit   Size
		want   uint64
	}{
		{
			name:   "1 petabyte",
			amount: 1,
			unit:   Petabyte,
			want:   1,
		},
		{
			name:   "0 petabytes",
			amount: 0,
			unit:   Petabyte,
			want:   0,
		},
		{
			name:   "1 petabyte as bytes",
			amount: 1,
			unit:   Byte,
			want:   1_125_899_906_842_624,
		},
		{
			name:   "1 petabyte as kilobytes",
			amount: 1,
			unit:   Kilobyte,
			want:   1_099_511_627_776,
		},
		{
			name:   "1 petabyte as megabytes",
			amount: 1,
			unit:   Megabyte,
			want:   1_073_741_824,
		},
		{
			name:   "1 petabyte as gigabytes",
			amount: 1,
			unit:   Gigabyte,
			want:   1_048_576,
		},
		{
			name:   "1 petabyte as terabytes",
			amount: 1,
			unit:   Terabyte,
			want:   1_024,
		},
		{
			name:   "1 petabyte as exabytes",
			amount: 1,
			unit:   Exabyte,
			want:   0,
		},
		{
			name:   "petabytes to a Exabyte",
			amount: 1_024,
			unit:   Exabyte,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PetabytesAs(tt.amount, tt.unit); got != tt.want {
				t.Errorf("PetabytesAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExabytes(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		want   uint64
	}{
		{
			name:   "1 exabyte",
			amount: 1,
			want:   1_152_921_504_606_846_976,
		},
		{
			name:   "2 exabytes",
			amount: 2,
			want:   2_305_843_009_213_693_952,
		},
		{
			name:   "0 exabytes",
			amount: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exabytes(tt.amount); got != tt.want {
				t.Errorf("Exabytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExabytesAs(t *testing.T) {
	tests := []struct {
		name   string
		amount uint64
		unit   Size
		want   uint64
	}{
		{
			name:   "1 exabyte",
			amount: 1,
			unit:   Exabyte,
			want:   1,
		},
		{
			name:   "0 exabytes",
			amount: 0,
			unit:   Exabyte,
			want:   0,
		},
		{
			name:   "1 exabyte as bytes",
			amount: 1,
			unit:   Byte,
			want:   1_152_921_504_606_846_976,
		},
		{
			name:   "1 exabyte as kilobytes",
			amount: 1,
			unit:   Kilobyte,
			want:   1_125_899_906_842_624,
		},
		{
			name:   "1 exabyte as megabytes",
			amount: 1,
			unit:   Megabyte,
			want:   1_099_511_627_776,
		},
		{
			name:   "1 exabyte as gigabytes",
			amount: 1,
			unit:   Gigabyte,
			want:   1_073_741_824,
		},
		{
			name:   "1 exabyte as terabytes",
			amount: 1,
			unit:   Terabyte,
			want:   1_048_576,
		},
		{
			name:   "1 exabyte as petabytes",
			amount: 1,
			unit:   Petabyte,
			want:   1_024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExabytesAs(tt.amount, tt.unit); got != tt.want {
				t.Errorf("ExabytesAs() = %v, want %v", got, tt.want)
			}
		})
	}
}
