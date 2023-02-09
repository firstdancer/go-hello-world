package methods

import "testing"

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		args    string
		want    string
		wantErr bool
	}{
		{
			args:    "",
			want:    "",
			wantErr: true,
		},
		{
			args: "root",
			want: "hello, root!",
		},
		{
			args: " root ",
			want: "hello, root!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			got, err := HelloWorld(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("HelloWorld() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HelloWorld() got = %v, want %v", got, tt.want)
			}
		})
	}
}
