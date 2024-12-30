package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

//	将结构体转换为 map
//
// 参数：
//   - s: 待转换的结构体
//
// 返回值：
//   - result: 转换之后的 map
//   - err: error
func StructToMapOmitEmpty(s interface{}) (result map[string]interface{}, err error) {
	values := reflect.ValueOf(s)
	sType := reflect.TypeOf(s)

	tagName := "mapstructure"

	// 检查传入的参数是否为结构体
	if sType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("错误：%s，传入的不是结构体数据", sType.Kind())
	}

	result = make(map[string]interface{})
	// 遍历结构体的所有字段
	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)
		fieldType := sType.Field(i)

		if field.Kind() == reflect.Bool || !field.IsZero() {
			mapKey := fieldType.Tag.Get(tagName)
			result[mapKey] = field.Interface()
		}
	}

	return result, nil
}

//	将结构体转换为 map[string]string
//
// 参数：
//   - s: 待转换的结构体
//
// 返回值：
//   - result: 转换之后的 map
//   - err: error
func StructToStringMap(s interface{}) (result map[string]string, err error) {
	values := reflect.ValueOf(s)
	sType := reflect.TypeOf(s)

	tagName := "mapstructure"

	// 检查传入的参数是否为结构体
	if sType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("错误：%s，传入的不是结构体数据", sType.Kind())
	}

	result = make(map[string]string)
	// 遍历结构体的所有字段
	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)
		fieldType := sType.Field(i)

		mapKey := fieldType.Tag.Get(tagName)
		result[mapKey] = field.String()
	}

	return result, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func Pointer[T any](in T) (out *T) {
	return &in
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// MaheHump 将字符串转换为驼峰命名
func MaheHump(s string) string {
	words := strings.Split(s, "-")

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	return strings.Join(words, "")
}

// 随机字符串
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[RandomInt(0, len(letters))]
	}
	return string(b)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
