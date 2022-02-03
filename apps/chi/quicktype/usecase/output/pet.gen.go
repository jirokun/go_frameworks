// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    pet, err := UnmarshalPet(bytes)
//    bytes, err = pet.Marshal()

package output

import "encoding/json"

func UnmarshalPet(data []byte) (Pet, error) {
	var r Pet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Pet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Pet struct {
	ID   int64   `json:"id"`  
	Name string  `json:"name"`
	Tag  *string `json:"tag"` 
}
