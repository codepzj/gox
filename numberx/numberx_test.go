package numberx

import "testing"

func TestMax(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		if Max(3, 5) != 5 {
			t.Errorf("Max(3,5) should be 5")
		}
	})
	t.Run("float64", func(t *testing.T) {
		if Max(5.5, 2.2) != 5.5 {
			t.Errorf("Max(5.5,2.2) should be 5.5")
		}
	})
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		a, b     any
		expected any
	}{
		{"int", 3, 5, 3},
		{"float64", 5.5, 2.2, 2.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.a.(type) {
			case int:
				if Min(v, tt.b.(int)) != tt.expected.(int) {
					t.Errorf("Min(%v,%v) should be %v", tt.a, tt.b, tt.expected)
				}
			case float64:
				if Min(v, tt.b.(float64)) != tt.expected.(float64) {
					t.Errorf("Min(%v,%v) should be %v", tt.a, tt.b, tt.expected)
				}
			}
		})
	}
}

func TestAbs(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		if Abs(-3) != 3 {
			t.Errorf("Abs(-3) should be 3")
		}
	})
	t.Run("float64", func(t *testing.T) {
		if Abs(-3.5) != 3.5 {
			t.Errorf("Abs(-3.5) should be 3.5")
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		if Sum(1, 2, 3) != 6 {
			t.Errorf("Sum(1,2,3) should be 6")
		}
	})
	t.Run("float64", func(t *testing.T) {
		if Sum(1.1, 2.2, 3.3) != 6.6 {
			t.Errorf("Sum(1.1,2.2,3.3) should be 6.6")
		}
	})
}

func TestAverage(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		if Average(1, 2, 3, 4) != 2.5 {
			t.Errorf("Average(1,2,3,4) should be 2.5")
		}
	})
}

func TestIsBetween(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		if !IsBetween(5, 1, 10) {
			t.Errorf("IsBetween(5,1,10) should be true")
		}
	})
	t.Run("float64", func(t *testing.T) {
		if !IsBetween(2.5, 1.0, 3.0) {
			t.Errorf("IsBetween(2.5,1.0,3.0) should be true")
		}
	})
}
