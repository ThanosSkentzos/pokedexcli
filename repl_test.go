package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input 		string
		expected	[]string 
	}{
		{
			input: 		"hello world ",
			expected: 	[]string{"hello","world"},
		},
		{
			input: 		"Hello world ",
			expected: 	[]string{"hello","world"},
		},
		{
			input: 		"Hello,world ",
			expected: 	[]string{"hello","world"},
		},
		{
			input: 		"Hello,World ",
			expected: 	[]string{"hello","world"},
		},
	}	

	for i,c := range cases{
		actual := cleanInput(c.input)
		if len(actual)!=len(c.expected){
			t.Errorf("Test '%v' Failed: expected length '%d', got '%v' ",i,len(c.expected),len(actual))
		}
		for pos := range actual{
			actualWord := actual[pos]
			expectedWord := c.expected[pos]
			if actualWord!= expectedWord{
				t.Errorf("Test '%v' Failed: expected '%v', got '%v' ",i,expectedWord,actualWord)	
			}
		}
	}
}