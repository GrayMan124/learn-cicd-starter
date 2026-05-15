package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAPI(t *testing.T) {
	type test struct {
		input  http.Header
		output string
	}
	tests := []test{
		{input: http.Header{}, output: "tego"},
		{input: http.Header{}, output: "tamtego"},
	}
	tests[0].input.Add("Authorization", "ApiKey tego")
	tests[1].input.Add("Authorization", "ApiKey tamtego")
	for i, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if err != nil {
			t.Fatalf("test %d: Unexpected error %v", i+1, err)
		}
		if !reflect.DeepEqual(tc.output, got) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.output, got)
		}

	}

}
