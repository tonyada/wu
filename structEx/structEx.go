package structEx

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// 获取结构体中字段的名称
// get struct all field names
func GetFieldNames(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

//获取结构体中Tag的值，如果没有tag则返回字段值
// get struct all 1 tags
func GetTagsByArrayPos(tagArrPos int, structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		tagName := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[tagArrPos]
		}
		result = append(result, tagName)
	}
	return result
}

// Get field Tag value
// struct { MarketCode int `em_clist:"f12"` }
// GetFieldTagValue("MarketCode", "em_clist", struct)return f12
func GetFieldTagValue(fieldName, tagName string, structName interface{}) string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return ""
	}
	fieldNum := t.NumField()
	result := ""
	for i := 0; i < fieldNum; i++ {
		if t.Field(i).Name == fieldName {
			// log.Println(t.Field(i).Name)
			tags := strings.Split(string(t.Field(i).Tag), "\"")
			// log.Println(tags)
			for k, v := range tags {
				tagLeft := strings.Trim(v, " ") // del 第二个tag前面的空格[ sec:] `json:"aa" sec:"bb"`
				// log.Println(k, tagLeft)
				if tagLeft == tagName+":" { // add tag: [sec:]
					result = tags[k+1]
					break
				}
			}
			break
		}
	}
	return result
}

// struct { MarketCode int `json:"market_code"` }
// GetFieldTagValue("MarketCode", struct) return market_code
func GetFieldJsonValue(fieldName string, structName interface{}) string {
	return GetFieldTagValue(fieldName, "json", structName)
}
func Print(st interface{}) {
	Explicit(reflect.ValueOf(st), 0)
}

func Explicit(v reflect.Value, depth int) {
	if v.CanInterface() {
		t := v.Type()
		switch v.Kind() {
		case reflect.Ptr:
			Explicit(v.Elem(), depth)
		case reflect.Struct:
			fmt.Printf(strings.Repeat("\t", depth)+"%v %v {\n", t.Name(), t.Kind())
			for i := 0; i < v.NumField(); i++ {
				f := v.Field(i)
				if f.Kind() == reflect.Struct || f.Kind() == reflect.Ptr {
					fmt.Printf(strings.Repeat("\t", depth+1)+"%s %s : \n", t.Field(i).Name, f.Type())
					Explicit(f, depth+2)
				} else {
					if f.CanInterface() {
						fieldType := fmt.Sprintf("%v", f.Type())
						if fieldType == "string" {
							fmt.Printf(strings.Repeat("\t", depth+1)+"%s \t%s \t: \"%v\" \n", t.Field(i).Name, f.Type(), f.Interface())
						} else {
							fmt.Printf(strings.Repeat("\t", depth+1)+"%s \t%s \t: %v \n", t.Field(i).Name, f.Type(), f.Interface())
						}
					} else {
						fieldType := fmt.Sprintf("%v", f.Type())
						if fieldType == "string" {
							fmt.Printf(strings.Repeat("\t", depth+1)+"%s \t%s \t: \"%v\" \n", t.Field(i).Name, f.Type(), f)
						} else {
							fmt.Printf(strings.Repeat("\t", depth+1)+"%s \t%s \t: %v \n", t.Field(i).Name, f.Type(), f)
						}
					}
				}
			}
			fmt.Println(strings.Repeat("\t", depth) + "}")
		}
	} else {
		fmt.Printf(strings.Repeat("\t", depth)+"%+v\n", v)
	}
}
