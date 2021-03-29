package declaration

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args    args
		want    Node
		wantErr bool
	}{
		{
			args: args{
				input: `@interface NSMenu : NSObject`,
			},
			want: Node(Interface{
				name: "NSMenu",
				superName: "NSObject",
			}),
		},
		{
			args: args{
				input: `- (NSRect)convertRectToBacking:(NSRect)rect;`,
			},
			want: Node(Method{
				baseName:   "convertRectToBacking",
				returnType: "NSRect",
				args:       []MethodArgs{
					{
						name:    "rect",
						argType: "NSRect",
					},
				},
			}),
		},
		{
			args: args{
				input: `+ (BOOL)menuBarVisible;`,
			},
			want: Node(Method{
				baseName:   "menuBarVisible",
				returnType: "BOOL",
			}),
		},
		{
			args: args{
				input: `+ (void)setMenuBarVisible:(BOOL)visible;`,
			},
			want: Node(Method{
				baseName:   "setMenuBarVisible",
				returnType: "void",
				args:       []MethodArgs{
					{
						name:    "visible",
						argType: "BOOL",
					},
				},
			}),
		},
		{
			args: args{
				input: `- (NSColor *)blendedColorWithFraction:(CGFloat)fraction 
                              ofColor:(NSColor *)color;`,
			},
			want: Node(Method{
				baseName:   "blendedColorWithFraction",
				returnType: "NSColor *",
				args:       []MethodArgs{
					{
						name:    "fraction",
						argType: "CGFloat",
					},{
						name:    "color",
						methodNameAddon: "ofColor",
						argType: "NSColor *",
					},
				},
			}),
		},
		{
			args: args{
				input: `- (void)resetCursorRects;`,
			},
			want: Node(Method{
				baseName:   "resetCursorRects",
				returnType: "void",
			}),
		},
		{
			args: args{
				input: `- (instancetype)initWithContentRect:(NSRect)contentRect 
                          styleMask:(NSWindowStyleMask)style 
                            backing:(NSBackingStoreType)backingStoreType 
                              defer:(BOOL)flag;`,
			},
			want: Node(Method{
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
			}),
		},
		{
			args: args{
				input: `@property CGFloat alphaValue;`,
			},
			want: Node(Property{
				name:   "alphaValue",
				propertyType: "CGFloat",
			}),
		},
		{
			args: args{
				input: `@property(getter=isVisible, readonly) BOOL visible;`,
			},
			want: Node(Property{
				name:   "visible",
				propertyType: "BOOL",
			}),
		},
		{
			args: args{
				input: `@property(class, readonly, strong) NSStatusBar *systemStatusBar;`,
			},
			want: Node(Property{
				name:   "*systemStatusBar",
				propertyType: "NSStatusBar",
			}),
		},
		{
			args: args{
				input: `- (void)removeStatusItem:(NSStatusItem *)item;`,
			},
			want: Node(Method{
				baseName:   "removeStatusItem",
				returnType: "void",
				args:       []MethodArgs{
					{
						name:    "item",
						argType: "NSStatusItem *",
					},
				},
			}),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("parse example %d", i), func(t *testing.T) {
			got, err := Parse(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
