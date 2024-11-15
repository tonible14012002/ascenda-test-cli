package commonultils

import (
	"encoding/json"
	"fmt"
)

func PrintJson(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
