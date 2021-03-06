// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/sml"
)

func TestCT_FileSharingConstructor(t *testing.T) {
	v := sml.NewCT_FileSharing()
	if v == nil {
		t.Errorf("sml.NewCT_FileSharing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FileSharing should validate: %s", err)
	}
}

func TestCT_FileSharingMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FileSharing()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FileSharing()
	xml.Unmarshal(buf, v2)
}
