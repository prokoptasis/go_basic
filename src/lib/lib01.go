package lib01

import "fmt"

var pop map[string]string

func init() {
	pop = make(map[string]string)
	pop["A"] = "Apple"
	pop["B"] = "Banana"
	pop["C"] = "Chocolate"
}

func GetItems(item string) string {
	return pop[item]
}

func getKeys() {
	for _, kv := range pop {
		fmt.Println(kv)
	}
}
