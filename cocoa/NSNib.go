package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSNib struct {
	objc.Object
}

func NSNib_Init(name string, bundle NSBundle) NSNib {
	return NSNib{objc.Get("NSNib").Alloc().Send("initWithNibNamed:bundle:",
		core.String(name), bundle)}
}

func (nib NSNib) InstantiateWithOwner(owner objc.Object) {
	nib.Send("instantiateNibWithOwner:topLevelObjects:", owner, nil)
}
