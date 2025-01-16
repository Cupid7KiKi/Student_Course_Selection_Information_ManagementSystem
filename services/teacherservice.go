package services

func GetTeachers() (results []map[string]interface{}) {
	results, err := GetDb().Query("select * from teachers")
	if err != nil {
		return
	}
	return
}
