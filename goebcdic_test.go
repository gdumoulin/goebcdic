// Copyright 2012 Guillaume DUMOULIN. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goebcdic

import (
	"testing"
)

// Wikipedia example
var teststring = "Man is distinguished, not only by his reason, but by this singular passion from " +
	"other animals, which is a lust of the mind, that by a perseverance of delight in " +
	"the continued and indefatigable generation of knowledge, exceeds the short " +
	"vehemence of any carnal pleasure."

func TestConsistencyConvertion(t *testing.T) {
	result := EBCDICtoASCIIofString(ASCIItoEBCDICofString(teststring))
	if teststring != result {
		t.Errorf("Consistency problem of convertion.", teststring, result)
	}
}
