package main

import "testing"

func TestGenerateRandomElements_ZeroSize(t *testing.T) {
	got := generateRandomElements(0)
	if got == nil {
		t.Fatalf("ожидался пустой слайс, получен nil")
	}
	if len(got) != 0 {
		t.Fatalf("ожидалась длина 0, получено %d", len(got))
	}
}

func TestGenerateRandomElements_NegativeSize(t *testing.T) {
	got := generateRandomElements(-5)
	if got == nil {
		t.Fatalf("ожидался пустой слайс, получен nil")
	}
	if len(got) != 0 {
		t.Fatalf("ожидалась длина 0 для отрицательного размера, получено %d", len(got))
	}
}

func TestGenerateRandomElements_PositiveSize(t *testing.T) {
	size := 10
	got := generateRandomElements(size)
	if got == nil {
		t.Fatalf("ожидался непустой (non-nil) слайс, получен nil")
	}
	if len(got) != size {
		t.Fatalf("ожидалась длина %d, получено %d", size, len(got))
	}
	for i, v := range got {
		if v < 0 {
			t.Fatalf("ожидались неотрицательные значения от rand.Int, получено %d на индексе %d", v, i)
		}
	}
}

// Тесты для maximum()
func TestMaximum_EmptySlice(t *testing.T) {
	var data []int
	got := maximum(data)
	if got != 0 {
		t.Fatalf("для пустого слайса ожидалось 0, получено %d", got)
	}
}

func TestMaximum_SingleElement(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want int
	}{
		{"ноль", []int{0}, 0},
		{"положительное", []int{42}, 42},
		{"отрицательное", []int{-17}, -17},
	}
	for _, tc := range cases {
		if got := maximum(tc.in); got != tc.want {
			t.Fatalf("%s: ожидалось %d, получено %d", tc.name, tc.want, got)
		}
	}
}

func TestMaximum_MultipleElements(t *testing.T) {
	data := []int{3, 7, 2, 9, 5}
	want := 9
	if got := maximum(data); got != want {
		t.Fatalf("ожидалось %d, получено %d", want, got)
	}
}

func TestMaximum_AllEqual(t *testing.T) {
	data := []int{5, 5, 5, 5}
	want := 5
	if got := maximum(data); got != want {
		t.Fatalf("ожидалось %d, получено %d", want, got)
	}
}

func TestMaximum_MixedNegatives(t *testing.T) {
	data := []int{-10, -3, -50, -4}
	want := -3
	if got := maximum(data); got != want {
		t.Fatalf("ожидалось %d, получено %d", want, got)
	}
}
