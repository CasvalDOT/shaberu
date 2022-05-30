package dialects

import (
	"encoding/csv"
	"os"
)

type csvdialect struct {
	dialect
}

func (d *csvdialect) Write(path *string, content map[string]string) error {
	outputFile, err := os.Create(*path)
	if err != nil {
		return err
	}

	defer outputFile.Close()

	contendEncoded := [][]string{}
	for key, value := range content {
		contendEncoded = append(contendEncoded, []string{key, value})
	}

	writer := csv.NewWriter(outputFile)
	writer.Comma = ';'

	err = writer.WriteAll(contendEncoded)
	if err != nil {
		return err
	}

	return nil
}

func (d *csvdialect) Read(path *string, separator *string) (map[string]string, error) {
	file, err := os.Open(*path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = []rune(*separator)[0]

	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	dataWithoutEmptyKeys := [][]string{}
	for _, item := range data {
		if item[0] != "" {
			dataWithoutEmptyKeys = append(dataWithoutEmptyKeys, item)
		}
	}

	output := map[string]string{}
	for _, item := range dataWithoutEmptyKeys {
		output[item[0]] = item[1]
	}

	return output, nil

}

func NewCSV() IDialect {
	return &csvdialect{
		dialect: dialect{},
	}
}
