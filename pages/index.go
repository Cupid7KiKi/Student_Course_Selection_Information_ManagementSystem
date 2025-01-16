package pages

import (
	"Student_Course_Selection_Information_ManagementSystem/pkg"
	"Student_Course_Selection_Information_ManagementSystem/services"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/adminlte/components/smallbox"
	template2 "html/template"
)

func GetDashBoard(ctx *context.Context) (types.Panel, error) {
	tmp := template.Default(ctx)

	rows := []pkg.BaseComponent{getRow1(tmp), getRow2(tmp)}
	return types.Panel{
		Title: "工作台",
		Content: func() template2.HTML {
			var html template2.HTML
			for i, _ := range rows {
				html += rows[i].GetContent()
			}
			return html
		}(),
	}, nil
}
func getRow1(tmp template.Template) types.RowAttribute {
	col := tmp.Col()
	col1 := col.SetSize(types.SizeMD(4)).SetContent(smallbox.New().SetTitle("课程数量").SetUrl("/admin/info/courses").SetValue(services.IntToTmp(len(services.GetCourses()))).SetColor("green").SetIcon(icon.Book).GetContent()).GetContent()
	col2 := col.SetSize(types.SizeMD(4)).SetContent(smallbox.New().SetTitle("学生数量").SetUrl("/admin/info/students").SetValue(services.IntToTmp(len(services.GetStudents()))).SetColor("green").SetIcon(icon.User).GetContent()).GetContent()
	col3 := col.SetSize(types.SizeMD(4)).SetContent(smallbox.New().SetTitle("教师数量").SetUrl("/admin/info/teachers").SetValue(services.IntToTmp(len(services.GetTeachers()))).SetColor("green").SetIcon(icon.User).GetContent()).GetContent()
	return tmp.Row().SetContent(col1 + col2 + col3)
}

func getRow2(tmp template.Template) types.RowAttribute {

	table := tmp.Table().SetType("table").SetInfoList(services.GetCoursesTable()).SetThead(types.Thead{
		{Head: "课程名称"},
		{Head: "课程时间"},
		{Head: "授课教师"},
		{Head: "创建时间"},
	}).GetContent()

	warpTable := template.HTML(fmt.Sprintf("<div class=\"table-responsive\">  %s</div>", table))

	tabs := template.Default().Tabs().
		SetData([]map[string]template2.HTML{
			{
				"title": "学生榜",
				"content": tmp.Table().SetStyle("table").SetInfoList(services.GetIndexStudentCoursesNums(0)).SetThead(types.Thead{
					{Head: "姓名", Width: "140px"},
					{Head: "参加次数"},
				}).GetContent(),
			}, {
				"title": "课程榜",
				"content": tmp.Table().SetStyle("table").SetThead(types.Thead{
					{Head: "课程", Width: "140px"},
					{Head: "被选次数"},
				}).SetInfoList(services.GetIndexCoursesNums(0)).GetContent(),
			},
		}).GetContent()

	newActivity := tmp.Col().SetSize(types.SizeMD(9)).SetContent(
		tmp.Box().WithHeadBorder().SetHeader("热门课程").SetBody(warpTable).GetContent()).GetContent()
	top := tmp.Col().SetSize(types.SizeMD(3)).SetContent(tabs).GetContent()
	return tmp.Row().SetContent(newActivity + top)
}
