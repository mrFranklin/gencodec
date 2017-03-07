// generated by github.com/fjl/gencodec, do not edit.

package sliceconv

import (
	"encoding/json"
	"errors"
)

func (x *X) MarshalJSON() ([]byte, error) {
	type XJSON struct {
		Slice []replacedInt
		Named namedSlice2
	}
	var enc XJSON
	if x.Slice != nil {
		enc.Slice = make([]replacedInt, len(x.Slice))
		for k, v := range x.Slice {
			enc.Slice[k] = replacedInt(v)
		}
	}
	if x.Named != nil {
		enc.Named = make(namedSlice2, len(x.Named))
		for k, v := range x.Named {
			enc.Named[k] = replacedInt(v)
		}
	}
	return json.Marshal(&enc)
}

func (x *X) UnmarshalJSON(input []byte) error {
	type XJSON struct {
		Slice []replacedInt
		Named namedSlice2
	}
	var dec XJSON
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	var x0 X
	if dec.Slice == nil {
		return errors.New("missing required field 'slice' for X")
	}
	x0.Slice = make([]int, len(dec.Slice))
	for k, v := range dec.Slice {
		x0.Slice[k] = int(v)
	}
	if dec.Named == nil {
		return errors.New("missing required field 'named' for X")
	}
	x0.Named = make(namedSlice, len(dec.Named))
	for k, v := range dec.Named {
		x0.Named[k] = int(v)
	}
	*x = x0
	return nil
}

func (x *X) MarshalYAML() (interface{}, error) {
	type XYAML struct {
		Slice []replacedInt
		Named namedSlice2
	}
	var enc XYAML
	if x.Slice != nil {
		enc.Slice = make([]replacedInt, len(x.Slice))
		for k, v := range x.Slice {
			enc.Slice[k] = replacedInt(v)
		}
	}
	if x.Named != nil {
		enc.Named = make(namedSlice2, len(x.Named))
		for k, v := range x.Named {
			enc.Named[k] = replacedInt(v)
		}
	}
	return &enc, nil
}

func (x *X) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type XYAML struct {
		Slice []replacedInt
		Named namedSlice2
	}
	var dec XYAML
	if err := unmarshal(&dec); err != nil {
		return err
	}
	var x0 X
	if dec.Slice == nil {
		return errors.New("missing required field 'slice' for X")
	}
	x0.Slice = make([]int, len(dec.Slice))
	for k, v := range dec.Slice {
		x0.Slice[k] = int(v)
	}
	if dec.Named == nil {
		return errors.New("missing required field 'named' for X")
	}
	x0.Named = make(namedSlice, len(dec.Named))
	for k, v := range dec.Named {
		x0.Named[k] = int(v)
	}
	*x = x0
	return nil
}
