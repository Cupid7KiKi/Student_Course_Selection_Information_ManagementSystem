package services

import (
	"fmt"
	"strings"
)

func GetCourseT(value string) (results []map[string]interface{}) {
	results, err := GetDb().Query("select tea_id from course_teacher where course_id" + "=" + value)
	if err != nil {
		return
	}
	return
}

func GetTeacherName(t []interface{}) (results []map[string]interface{}) {
	strs := make([]string, len(t))
	for i, n := range t {
		strs[i] = fmt.Sprintf("%d", n)
	}
	result := strings.Join(strs, ",")
	results, err := GetDb().Query("SELECT * FROM teachers WHERE id " + "IN" + "(" + result + ")")
	if err != nil {
		return
	}
	return
}
