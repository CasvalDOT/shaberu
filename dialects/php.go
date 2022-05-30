package dialects

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
)

type phpdialect struct {
	dialect
}

func (d *phpdialect) prepareContent(content map[string]interface{}, final string, level int) string {
	indent := ""

	for key, value := range content {
		switch value := value.(type) {
		case string:
			valueAsString := value
			doubleQuoteRegex := regexp.MustCompile("\"")
			doubleQuoteRegex.ReplaceAllString(valueAsString, `\\"`)
			final = final + indent + "'" + key + "' => \"" + valueAsString + "\",\n"
		default:
			final = final + "'" + key + "'" + " => [\n"
			final = d.prepareContent(value.(map[string]interface{}), final, level+1)
			final = final + "\n],\n"
		}
	}

	return final
}

func (d *phpdialect) Write(path *string, content map[string]string) error {
	file, err := os.Create(*path)
	if err != nil {
		return err
	}

	contentAsInterface := map[string]interface{}{}
	contentAsInterface = d.dialect.unflat(content, contentAsInterface)

	contentAsString := ""
	contentAsString = d.prepareContent(contentAsInterface, contentAsString, 0)

	defer file.Close()

	file.WriteString("<?php\n\n")
	file.WriteString("return [")
	file.WriteString(contentAsString)
	file.WriteString("];")

	return nil
}

func (d *phpdialect) Read(path *string, separator *string) (map[string]string, error) {
	// The creation of a system
	// that parse a php file dictionary
	// is a little tricky. So please make sure the php file follow the form:
	/*
		<?php
		return [
			'key' => 'value'
			...
		];
	*/

	contents, err := ioutil.ReadFile(*path)
	if err != nil {
		return nil, err
	}

	contents = regexp.MustCompile(`<\?php`).ReplaceAll(contents, []byte(""))
	contents = regexp.MustCompile(`return`).ReplaceAll(contents, []byte(""))
	contents = regexp.MustCompile(`;`).ReplaceAll(contents, []byte(""))
	contents = regexp.MustCompile(`\[`).ReplaceAll(contents, []byte("{"))
	contents = regexp.MustCompile(`=>`).ReplaceAll(contents, []byte(":"))
	contents = regexp.MustCompile(`\]`).ReplaceAll(contents, []byte("\"\":\"\"\n}"))

	var anonMap map[string]string
	err = json.Unmarshal(contents, &anonMap)
	if err != nil {
		return nil, err
	}

	anonMapWithoutEmptyKeys := map[string]string{}
	for key, value := range anonMap {
		if key != "" {
			anonMapWithoutEmptyKeys[key] = value
		}
	}

	return anonMapWithoutEmptyKeys, nil
}

func NewPHP() IDialect {
	return &phpdialect{
		dialect: dialect{},
	}
}
