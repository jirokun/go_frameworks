// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    newPet, err := UnmarshalNewPet(bytes)
//    bytes, err = newPet.Marshal()

package input

import "encoding/json"

func UnmarshalNewPet(data []byte) (NewPet, error) {
	var r NewPet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NewPet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NewPet struct {
	Name string  `json:"name"`
	Tag  *string `json:"tag"` 
}
