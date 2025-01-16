package services

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/template/types"
	"strconv"
)

func GetCourses() (results []map[string]interface{}) {
	results, err := GetDb().Query("select * from courses")
	if err != nil {
		return
	}
	return
}

func GetCoursesTable() []map[string]types.InfoItem {

	var tableData []map[string]types.InfoItem
	query, err := GetDb().Query("SELECT\n\tt.name,\n\tc.title,\n\tc.start_date,\n\tc.end_date, ct.created_at FROM\n\tcourse_teacher AS ct\n\tJOIN courses AS c ON ct.course_id = c.id\n\tJOIN teachers AS t ON ct.tea_id = t.id\n\tORDER BY c.start_date asc  limit 10")
	if err != nil {
		fmt.Println(err)
		return tableData
	}

	for i, _ := range query {
		item := map[string]types.InfoItem{
			"授课教师": {Content: "<span class=\"label label-info\">" + TansTmp(query[i]["name"]) + "</span>"},
			"课程名称": {Content: TansTmp(query[i]["title"])},
			"课程时间": {Content: TansTmp(fmt.Sprintf("%s-%s", query[i]["start_date"], query[i]["end_date"]))},
			"创建时间": {Content: TansTmp(query[i]["created_at"])},
		}
		tableData = append(tableData, item)
	}
	return tableData
}

func GetIndexCoursesNums(cId int) []map[string]types.InfoItem {
	var sql string
	var data []map[string]types.InfoItem
	if cId == 0 {
		sql = "select count(*) as nums,title FROM courses as c join select_course as a on a.course_id = c.id  GROUP BY c.id ORDER BY nums DESC  limit 10"
	} else {
		sql = "select count(*) as nums,title FROM courses as c join select_course as a on a.course_id = c.id where c.id = " + strconv.Itoa(cId) + " GROUP BY c.id ORDER BY nums DESC  limit 10"
	}
	query, err := GetDb().Query(sql)
	if err != nil {
		return data
	}
	for i, _ := range query {
		item := map[string]types.InfoItem{
			"课程": {
				Content: TansTmp(query[i]["title"]),
			},
			"被选次数": {
				Content: "<span class=\"label label-info\">" + TansTmp(query[i]["nums"]) + "</span>",
			},
		}
		data = append(data, item)
	}
	return data
}
