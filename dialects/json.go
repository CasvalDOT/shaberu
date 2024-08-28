package dialects

import (
	"encoding/json"
	"fmt"
	"os"
)

type jsondialect struct {
	dialect
}

func (d *jsondialect) Write(path *string, content map[string]string) error {
	final := map[string]interface{}{}
	unflattedContent := d.dialect.unflat(content, final)

	jsonDataAsBytes, err := json.Marshal(unflattedContent)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	outputFile, err := os.Create(*path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer outputFile.Close()

	outputFile.Write(jsonDataAsBytes)

	return nil
}

func (d *jsondialect) Read(path *string, separator *string) (map[string]string, error) {
	contents, err := os.ReadFile(*path)
	if err != nil {
		return nil, err
	}

	var anonMap map[string]interface{}
	err = json.Unmarshal(contents, &anonMap)
	if err != nil {
		return nil, err
	}

	anonMapFlatted := map[string]string{}
	anonMapFlatted = d.dialect.flat(anonMap, "", anonMapFlatted)

	anonMapWithoutEmptyKeys := map[string]string{}
	for key, value := range anonMapFlatted {
		if key != "" {
			anonMapWithoutEmptyKeys[key] = value
		}
	}

	return anonMapWithoutEmptyKeys, nil
}

func NewJSON() IDialect {
	return &jsondialect{
		dialect: dialect{},
	}
}
