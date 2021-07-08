package adapters

import "testing"

func TestAddSlashToPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "positive: slash should be added to the end of the path if it is not exists",
			args: args{
				path: "./data/source",
			},
			want: "./data/source/",
		},
		{
			name: "positive: slash should not be added to the end of the path if it is exists",
			args: args{
				path: "./data/source/",
			},
			want: "./data/source/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddSlashToPath(tt.args.path); got != tt.want {
				t.Errorf("AddSlashToPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileName(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "positive: file name should be selected from path",
			args: args{
				path: "./data/source/input.xlsx::Products",
			},
			want: "input",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileName(tt.args.path); got != tt.want {
				t.Errorf("GetFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileExt(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "positive: file extension should be selected from path",
			args: args{
				path: "./data/source/input.xlsx",
			},
			want: "xlsx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileExt(tt.args.path); got != tt.want {
				t.Errorf("GetFileExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isXLSXFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "positive: file with extension .xlsx is considered as excel file",
			args: args{
				path: "./data/input.Xlsx",
			},
			want: true,
		},
		{
			name: "positive: file with extension .xls is considered as excel file",
			args: args{
				path: "./data/input.Xls",
			},
			want: true,
		},
		{
			name: "positive: file with extension differs from .xlsx or .xls is considered as NOT excel file",
			args: args{
				path: "./data/input.csv",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isXLSXFile(tt.args.path); got != tt.want {
				t.Errorf("isXLSXFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
