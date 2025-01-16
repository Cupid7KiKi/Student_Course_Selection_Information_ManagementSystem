package services

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	selection "github.com/GoAdminGroup/go-admin/template/types/form/select"
	template2 "html/template"
	"reflect"
	"strconv"
)

var myDb db.Connection

func SetDb(appDb db.Connection) {
	myDb = appDb
}
func GetDb() db.Connection {
	return myDb
}
func TransFieldOptions(result interface{}, text, value string) types.FieldOptions {
	rV := reflect.ValueOf(result)
	rT := reflect.TypeOf(result)
	var myOptions types.FieldOptions
	if rT.Kind() == reflect.Slice {
		for i := 0; i < rV.Len(); i++ {
			rVItem := rV.Index(i)
			switch rVItem.Kind() {
			case reflect.Struct:
				textValue := rVItem.FieldByName(text)
				valueValue := rVItem.FieldByName(value)
				myOptions = append(myOptions, types.FieldOption{
					Text:  TransStr(textValue),
					Value: TransStr(valueValue),
				})
			case reflect.Map:
				textValue := rVItem.MapIndex(reflect.ValueOf(text))
				valueValue := rVItem.MapIndex(reflect.ValueOf(value))
				myOptions = append(myOptions, types.FieldOption{
					Text:  TransStr(textValue),
					Value: TransStr(valueValue),
				})
			default:
				return myOptions
			}

		}
	}
	return myOptions
}

func TransSelectionOptions(result interface{}, text, value string) selection.Options {
	rV := reflect.ValueOf(result)
	rT := reflect.TypeOf(result)
	var myOptions selection.Options
	if rT.Kind() == reflect.Slice {
		for i := 0; i < rV.Len(); i++ {
			rVItem := rV.Index(i)
			switch rVItem.Kind() {
			case reflect.Struct:
				textValue := rVItem.FieldByName(text)
				valueValue := rVItem.FieldByName(value)
				myOptions = append(myOptions, selection.Option{
					Text: TransStr(textValue),
					ID:   TransStr(valueValue),
				})
			case reflect.Map:
				textValue := rVItem.MapIndex(reflect.ValueOf(text))
				valueValue := rVItem.MapIndex(reflect.ValueOf(value))
				myOptions = append(myOptions, selection.Option{
					Text: TransStr(textValue),
					ID:   TransStr(valueValue),
				})
			default:
				return myOptions
			}

		}
	}
	return myOptions
}
func TransStr(t reflect.Value) string {
	switch t.Kind() {
	case reflect.Int64, reflect.Int:
		return strconv.FormatInt(t.Int(), 10)
	case reflect.String:
		return t.String()
	case reflect.Interface:
		return TransStr(reflect.ValueOf(t.Interface()))
	default:
		return ""
	}
}

// TransStr 将 reflect.Value 转换为 string，如果值是有效的且可以转换为 string，则返回转换后的字符串，否则返回空字符串。
//
//	func TransStr1(v reflect.Value) string {
//		if v.IsValid() && v.CanInterface() {
//			return fmt.Sprintf("%v", v.Interface())
//		}
//		return ""
//	}
func IntToTmp(v int) template2.HTML {
	return template.HTML(strconv.Itoa(v))
}

func TansTmp(v interface{}) template2.HTML {
	kind := reflect.TypeOf(v).Kind()
	switch kind {
	case reflect.String:
		return template2.HTML(v.(string))
	case reflect.Int:
		return IntToTmp(v.(int))
	case reflect.Int64:
		return template2.HTML(strconv.FormatInt(v.(int64), 10))
	default:
		fmt.Printf("当前类型不支持%v\n", kind)
	}
	return ""
}
