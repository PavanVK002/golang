// math_test.go
package unit_test

import (
    "testing"
     // Importing the local package
     "stock_market"
      
)

func TestAdd(t *testing.T) {
    cases := []struct {
        a, b     int
        expected int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
        {10, -5, 5},
    }

    for _, tc := range cases {
        result :=stock_market. Add_(tc.a, tc.b) // Accessing Add function from test_math package
        if result != tc.expected {
            t.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
        }
    }
}

func TestSubtract(t *testing.T) {
    cases := []struct {
        a, b     int
        expected int
    }{
        {3, 2, 1},
        {0, 0, 0},
        {1, 1, 0},
        {-5, -10, 5},
    }

    for _, tc := range cases {
        result := stock_market.Subtract_(tc.a, tc.b) // Accessing Subtract function from test_math package
        if result != tc.expected {
            t.Errorf("Subtract(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
        }
    }
}
