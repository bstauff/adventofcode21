package helpers

import (
	"bufio"
	"os"
)

//todo rewrite this with generics when 1.18 is out
func ReadInput(filePath string, transform func(s string) (interface{}, error)) ([]interface{}, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	transformedData := make([]interface{}, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		transformResult, err := transform(scanner.Text())
		if err != nil {
			return nil, err
		}
		transformedData = append(transformedData, transformResult)
	}

	return transformedData, nil
}
