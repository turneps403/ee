package yaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadAndMerge(files []string) map[interface{}]interface{} {
	ret := map[interface{}]interface{}{}
	for _, f := range files {
		data, _ := ioutil.ReadFile(f)
		currentMap := map[interface{}]interface{}{}
		if err := yaml.Unmarshal(data, &currentMap); err != nil {
			continue
		}
		ret = mergeMaps(ret, currentMap)
	}
	return ret
}

func mergeMaps(a, b map[interface{}]interface{}) map[interface{}]interface{} {
	out := make(map[interface{}]interface{}, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		if v, ok := v.(map[interface{}]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[interface{}]interface{}); ok {
					out[k] = mergeMaps(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}
