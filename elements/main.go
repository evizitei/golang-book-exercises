package main

import "fmt"

func buildElement(name string, state string) map[string]string {
	return map[string]string{
		"name":  name,
		"state": state,
	}
}

func main() {
	elements := map[string]map[string]string{
		"H":  buildElement("Hydrogen", "gas"),
		"He": buildElement("Helium", "gas"),
		"Li": buildElement("Lithium", "solid"),
		"Be": buildElement("Beryllium", "solid"),
		"B":  buildElement("Boron", "solid"),
		"C":  buildElement("Carbon", "solid"),
		"N":  buildElement("Nitrogen", "gas"),
		"O":  buildElement("Oxygen", "gas"),
		"F":  buildElement("Fluorine", "gas"),
		"Ne": buildElement("Neon", "gas"),
	}

	fmt.Println(elements["Li"])
	name, ok := elements["Un"]
	fmt.Println(name, ok)
}
