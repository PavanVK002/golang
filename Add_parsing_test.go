// math_test.go
package main

import (
    "testing"
     // Importing the local packag
      
)

func TestAdd_parsing(t *testing.T) {
    cases := []struct {
        a    string
        expected_result int
		expected_bool bool
    }{
        {"What is 2 plus 3?", 5, true},
        {"What is 3 plus 5 multiplied by 2?", 16, true},
        {"What is 52 cubed?", 0, false},
        {"What is -5 minus -5 plus 3 multiplied by 2?", 6, true},
    }

    for _, tc := range cases {
        result,result_bool := Add_parsing(tc.a) // Accessing Add function from test_math package
        if result != tc.expected_result && result_bool!= tc.expected_bool {
            t.Errorf("Add_parsing(%s) = %d,%t; expected %d,%t", tc.a, result,result_bool,tc.expected_result,tc.expected_bool)
        }
    }
}

