package monkeypatch

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
)

func init() {
	gomonkey.ApplyFunc(GetRandomCode, func() int { return 666 })
}

func TestRandomCode(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "test1",
			want: 666,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomCode(); got != tt.want {
				t.Errorf("GetRandomCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
