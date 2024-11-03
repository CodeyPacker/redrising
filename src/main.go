package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Character struct {
	Name       string   `json:"name"`
	Pronoun       string   `json:"pronoun"`
	Color      string   `json:"color"`
	House      string   `json:"house"`
	Background string   `json:"background"`
	KeyEvents  []string `json:"keyEvents"`
}

func lowerCaseString(str string) string {
    return strings.ToLower(string(str[0])) + str[1:]
}

func readFile() ([]Character, error) {
	file, err := os.Open("characters.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var characters []Character

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&characters)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return characters, nil
}

func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("👇 Enter a character's name from Red Rising")
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	return input, err
}

func lookupCharacter(name string, data []Character) ([]Character, error) {
	var foundCharacter []Character

	for _, c := range data {
		if strings.Contains(lowerCaseString(c.Name), lowerCaseString(name)) {
			foundCharacter = append(foundCharacter, c)
      break
		}
	}

	if foundCharacter == nil {
	   return nil, fmt.Errorf("No matching characters found for %s", name)
	 }

	return foundCharacter, nil
}
 
func describeCharacter (character []Character) error {
  for _, c := range character {
    c.Background = lowerCaseString(c.Background)
    fmt.Printf("\n\n%s is a %s of %s. \n %s is %s \n Key events:", c.Name, c.Color, c.House, c.Pronoun, c.Background) 
    
    for _, k := range c.KeyEvents {
      fmt.Printf("\n - %s", k) 
    }
  }

  return nil
}

func main() {
	characterData, err := readFile()
	if err != nil {
		fmt.Println("No character data found", err)
	}

	input, err := getInput()
	if err != nil {
		fmt.Println("Couldn't read the input", err)
	}

	selectedCharacter, err := lookupCharacter(input, characterData)
	if err != nil {
		fmt.Println(err)
	}

  describeCharacter(selectedCharacter)
}
