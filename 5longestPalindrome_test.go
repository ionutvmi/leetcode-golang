package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expandPalindrome(t *testing.T) {
	type args struct {
		chars []string
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expandPalindrome(tt.args.chars, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("expandPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
