// Currently Not To Be Used in This Repo, Change Below pkg When Required and Structure Properly
package fake

import (
	"encoding/json"
	"fmt"
)

var jsonData = []byte(`
    {
        "name": "Raj",
        "age": 25,
        "is_student": true,
        "grades": [78, 85, 92],
        "address": {
            "city": "Mumbai",
            "zipcode": "400001"
        },
        "additional_info": {
            "hobbies": ["cricket", "cooking"],
            "languages": {
                "English": "fluent",
                "Hindi": "native"
            }
        }
    }
`)

func main() {
	var data interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(data)

	if obj, ok := data.(map[string]interface{}); ok {
		name := obj["name"].(string)
		age := obj["age"].(float64)
		isStudent := obj["is_student"].(bool)

		fmt.Println("Name:", name)
		fmt.Println("Age:", age)
		fmt.Println("Is Student:", isStudent)

		address := obj["address"].(map[string]interface{})
		city := address["city"].(string)
		zipcode := address["zipcode"].(string)

		fmt.Println("City:", city)
		fmt.Println("Zipcode:", zipcode)

		grades := obj["grades"].([]interface{})
		fmt.Println("Grades:", grades)

		hobbies := obj["additional_info"].(map[string]interface{})["hobbies"].([]interface{})
		fmt.Println("Hobbies:", hobbies)

		languages := obj["additional_info"].(map[string]interface{})["languages"].(map[string]interface{})
		fmt.Println("Languages:", languages)
	}
}
