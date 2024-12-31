package maputils

import (
	"fmt"
	"reflect"
)

//	StructToMap 将结构体转换为 map[string]interface{} 类型的映射
//
// 参数:
//   - obj: 需要转换的结构体对象
//
// 返回值:
//   - map[string]interface{}: 返回一个包含结构体字段名和对应值的映射
func StructToMap(obj interface{}) map[string]interface{} {
	// 获取传入对象的类型和值
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	// 创建一个空的 map 用于存储结果
	data := make(map[string]interface{})

	// 遍历结构体的每个字段
	for i := 0; i < objType.NumField(); i++ {
		// 获取字段的标签，如果存在 "mapstructure" 标签则使用该标签作为键
		if objType.Field(i).Tag.Get("mapstructure") != "" {
			data[objType.Field(i).Tag.Get("mapstructure")] = objValue.Field(i).Interface()
		} else {
			// 否则使用字段名作为键
			data[objType.Field(i).Name] = objValue.Field(i).Interface()
		}
	}
	return data // 返回包含结构体字段及其值的映射
}

// StructToMapOmitEmpty 将结构体转换为 map，忽略空字段
//
// 参数：
//   - s: 待转换的结构体
//
// 返回值：
//   - result: 转换之后的 map，包含非空字段
//   - err: 错误信息，如果传入的不是结构体则返回相应的错误
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

		// 检查字段是否为布尔类型或非零值
		if field.Kind() == reflect.Bool || !field.IsZero() {
			mapKey := fieldType.Tag.Get(tagName)
			// 如果没有设置 mapstructure 标签，则使用字段名作为键
			if mapKey == "" {
				mapKey = fieldType.Name
			}
			result[mapKey] = field.Interface()
		}
	}

	return result, nil
}
