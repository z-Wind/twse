package twse

import "testing"

func TestResolveRelative(t *testing.T) {
	type args struct {
		basePath string
		elem     []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test", args{"https://www.alphavantage.co/query", []string{"accounts", "a", "b"}}, "https://www.alphavantage.co/query/accounts/a/b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveRelative(tt.args.basePath, tt.args.elem...); got != tt.want {
				t.Errorf("ResolveRelative() = %v, want %v", got, tt.want)
			}
		})
	}
}
