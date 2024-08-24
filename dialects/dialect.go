package dialects

import (
	"strconv"
	"strings"
)

type dialect struct{}

type IDialect interface {
	Read(*string, *string) (map[string]string, error)
	Write(*string, map[string]string) error
}

func (d *dialect) flat(data map[string]interface{}, startKey string, final map[string]string) map[string]string {
	for key, value := range data {
		keyToUse := key
		if startKey != "" {
			keyToUse = startKey + "." + key
		}

		switch value := value.(type) {
		case string:
			final[keyToUse] = value
		case []interface{}:
			tmp := map[string]interface{}{}
			for i, v := range value {
				tmp[strconv.Itoa(i)] = v
			}
			d.flat(tmp, keyToUse, final)
		default:
			d.flat(value.(map[string]interface{}), keyToUse, final)
		}
	}

	return final
}

func (d *dialect) unflat(data map[string]string, final map[string]interface{}) map[string]interface{} {
	for key, value := range data {
		keySplitted := strings.Split(key, ".")

		if len(keySplitted) == 1 {
			final[key] = value
		} else {
			newKey := strings.Join(keySplitted[1:], ".")
			if final[keySplitted[0]] == nil {
				final[keySplitted[0]] = make(map[string]interface{})
			}

			if len(keySplitted) == 2 {
				final[keySplitted[0]].(map[string]interface{})[keySplitted[1]] = value
			} else {
				fakeObj := map[string]string{}
				fakeObj[newKey] = value
				final[keySplitted[0]] = d.unflat(fakeObj, final[keySplitted[0]].(map[string]interface{}))
			}
		}
	}

	return final
}

func New(format string) IDialect {
	switch format {
	case "csv":
		return NewCSV()
	case "json":
		return NewJSON()
	case "php":
		return NewPHP()
	default:
		return nil
	}
}
