package declaration

import "testing"

func Test_Method_Name(t *testing.T) {
	tests := []struct {
		name   string
		method Method
		want   string
	}{
		{
			name:   "with many args",
			method: Method{
				baseName:   "initWithContentRect",
				returnType: "instancetype",
				args:       []MethodArgs{
					{
						name:    "contentRect",
						argType: "NSRect",
					},{
						name:    "style",
						methodNameAddon: "styleMask",
						argType: "NSWindowStyleMask",
					},{
						name:    "backingStoreType",
						methodNameAddon: "backing",
						argType: "NSBackingStoreType",
					},{
						name:    "flag",
						methodNameAddon: "defer",
						argType: "BOOL",
					},
				},
			},
			want:   "initWithContentRect:styleMask:backing:defer:",
		},
		{
			name: "with one arg",
			method: Method{
				baseName:   "convertRectToBacking",
				returnType: "NSRect",
				args:       []MethodArgs{
					{
						name:    "rect",
						argType: "NSRect",
					},
				},
			},
			want: "convertRectToBacking:",
		},
		{
			name:   "with no args",
			method: Method{
				baseName:   "resetCursorRects",
				returnType: "void",
			},
			want:   "resetCursorRects",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.method.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
