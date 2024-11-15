package domain

import (
	"encoding/json"
	"strconv"
)

// NOTE: Custom type for fallback parsing float64 from any type
// Already include `isValid` flag, dont need omitempty

type JsonFloat struct {
	Value   float64
	IsValid bool
}

func (f *JsonFloat) UnmarshalJSON(data []byte) error {
	if data == nil || string(data) == "null" || string(data) == "" {
		return nil
	}

	// able to pasrse as float64
	var num float64
	if err := json.Unmarshal(data, &num); err == nil {
		f.Value = num
		f.IsValid = true
		return nil
	}
	// able to pasrse as string
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		if str == "" {
			return nil
		}
		if num, err := strconv.ParseFloat(str, 64); err == nil {
			f.Value = num
			f.IsValid = true
			return nil
		}
	}
	return nil
}
