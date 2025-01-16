package services

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"strconv"
)

func GetStudents() (results []map[string]interface{}) {
	results, err := GetDb().Query("select * from students")
	if err != nil {
		return
	}
	return
}

func GetIndexStudentCoursesNums(sId int) []map[string]types.InfoItem {
	var sql string
	var data []map[string]types.InfoItem
	if sId == 0 {
		sql = "select COUNT(*) as nums,name FROM students as s join select_course as sa on sa.std_id = s.id GROUP BY s.id ORDER BY nums DESC limit 10"
	} else {
		sql = "select COUNT(*) as nums,name FROM students as s join select_course as sa on sa.std_id = s.id where s.id = " + strconv.Itoa(sId) + " GROUP BY s.id ORDER BY nums DESC limit 10"
	}
	query, err := GetDb().Query(sql)
	if err != nil {
		return data
	}
	for i, _ := range query {
		item := map[string]types.InfoItem{
			"姓名": {
				Content: TansTmp(query[i]["name"]),
			},
			"参加次数": {
				Content: "<span class=\"label label-info\">" + TansTmp(query[i]["nums"]) + "</span>",
			},
		}
		data = append(data, item)
	}
	return data
}
