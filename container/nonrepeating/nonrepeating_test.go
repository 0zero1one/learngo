package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct{
		s string
		ans int
	}{
		{"abcabcbb", 3},
		{"123456123456712", 7},
		{"我是阿飞,你是谁啊", 7},
		{"我是afei,你是谁啊", 9},
	}

	for _, tt := range tests {
		actual := mbnorepeating(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "我是afei,你是谁啊"
	for i := 0; i < 20; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 10
	b.ResetTimer()

	for i := 0; i< b.N; i++ {
		actual := mbnorepeating(s)
		if actual != ans {
			b.Errorf("got %d input %s; expected %d", actual, s, ans)
		}
	}
}