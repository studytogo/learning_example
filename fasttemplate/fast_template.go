package fasttemplate

import (
	"fmt"
	"regexp"
)

func StudyFastTemplate() {
	//
	//json := `{ {{name}} }`
	//
	//var m = map[string]interface{}{
	//	"name": "zzg",
	//}
	//
	//newString := fasttemplate.New(json, "{{", "}}").ExecuteString(m)
	//
	//fmt.Println(newString)
	withParam := `{{tasks.split.outputs.result}}`
	regexp, _ := regexp.Compile(`{{tasks\.([a-z0-9\-_]+)\.outputs.result}}`)
	result := regexp.FindStringSubmatch(withParam)
	fmt.Println("--", result)
}
