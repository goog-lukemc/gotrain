package main

import (
	"errors"
	"testing"
)

func TestGetPathBase(t *testing.T) {
	tests := []struct {
		url    string
		result string
		err    error
	}{
		{"http://www.homedepot.com/findit", "findit", nil},
		{"http://google.com/search", "search", nil},
		{"http://bo.bo.com/one/two/shoe/", "shoe", nil},
		{"http://youtube.com/path/to/file?is=12345", "file", nil},
		{"", "", errors.New(`parse : empty url`)},
	}

	for _, tt := range tests {

		t.Run(tt.url, func(t *testing.T) {

			// Run our function
			gotBase, gotErr := GetPathBase(tt.url)

			// If we got an error we weren't expecting fail the test
			if gotErr != nil {
				// Check to see if we got an error we expected
				if gotErr.Error() != tt.err.Error() {

					// using %q when dealing with strings will make your test debugging a bit nicer 😀
					t.Errorf("GetPathBase() err = %q; want %q", gotErr.Error(), tt.err.Error())
				}
				return
			}

			// if we got a result we expected
			if *gotBase != tt.result {
				t.Errorf("Nope - Check out the path package!, got: %v, want: %v.", tt.url, tt.result)
			}
		})

	}
}
