package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests :=
		map[string]struct {
			input http.Header
			want  string
		}{
			"first_case": {
				input: http.Header{
					"Authorization": []string{"ApiKey MarouaneAPIKey47"},
				},
				want: "MarouaneAPIKey4",
			},
		}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if err != nil {
				t.Fatalf("failed: %v", err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
