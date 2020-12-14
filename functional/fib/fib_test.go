package main

import "testing"

func TestRecFib(t *testing.T) {
	tests := []struct{
		number int
		ans int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{9, 34},
		{20, 6765},
	}

	for _, tt := range tests {
		actul := RecFib(tt.number)
		if actul != tt.ans {
			t.Errorf("got %d for input %d; expected %d", actul, tt.number, tt.ans)
		}
	}
}

func TestCountFib(t *testing.T) {
	tests := []struct{
		a int
		b int
		count int
		ans int
	}{
		{0, 1, 1, 1},
		{0, 1, 2, 2},
		{0, 1, 3, 3},
		{0, 1, 4, 5},
		{0, 1, 5, 8},
		{0, 1, 6, 13},
		{0, 1, 7, 21},
		{0, 1, 8, 34},
	}

	for _, tt := range tests {
		actul := CountFib(tt.a, tt.b, tt.count)
		if actul != tt.ans {
			t.Errorf("got %d for input %d; expected %d", actul, tt.count, tt.ans)
		}
	}
}

func BenchmarkRecFib(b *testing.B) {
	number := 50
	ans := 12586269025
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actul := RecFib(number)
		if actul != ans {
			b.Errorf("got %d for input %d; expected %d", actul, number, ans)
		}
	}
}

func BenchmarkCountFib(b *testing.B) {
	a := 0
	n := 1
	count := 49
	ans := 12586269025
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actul := CountFib(a, n, count)
		if actul != ans {
			b.Errorf("got %d for input %d; expected %d", actul, count, ans)
		}
	}
}