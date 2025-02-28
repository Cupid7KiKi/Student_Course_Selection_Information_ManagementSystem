package pages

import (
	"Student_Course_Selection_Information_ManagementSystem/services"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	tmpl "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCourseteacherTable(ctx *context.Context) table.Table {

	courseTeacher := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	user := auth.Auth(ctx)

	info := courseTeacher.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int)
	info.AddField("教师姓名", "name", db.Int).
		FieldJoin(types.Join{
			Table:     "teachers", // 连表的表名
			Field:     "tea_id",   // 要连表的字段
			JoinField: "id",       // 连表的表的字段
		}).
		//筛选时支持模糊查询
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return value.Row["teachers_goadmin_join_name"]
		})
	info.AddField("课程名称", "title", db.Int).
		FieldJoin(types.Join{
			Table:     "courses",   // 连表的表名
			Field:     "course_id", // 要连表的字段
			JoinField: "id",        // 连表的表的字段
		}).
		//筛选时支持模糊查询
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return value.Row["courses_goadmin_join_title"]
		})

	detail := courseTeacher.GetDetail()
	detail.AddField("ID", "id", db.Int)
	detail.AddField("教师姓名", "name", db.Int).
		FieldJoin(types.Join{
			Table:     "teachers", // 连表的表名
			Field:     "tea_id",   // 要连表的字段
			JoinField: "id",       // 连表的表的字段
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return value.Row["teachers_goadmin_join_name"]
		})
	detail.AddField("课程名称", "title", db.Int).
		FieldJoin(types.Join{
			Table:     "courses",   // 连表的表名
			Field:     "course_id", // 要连表的字段
			JoinField: "id",        // 连表的表的字段
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return value.Row["courses_goadmin_join_title"]
		})
	detail.AddField("课程介绍", "description", db.Longtext).
		FieldJoin(types.Join{
			Table:     "courses",
			Field:     "description",
			JoinField: "id",
		}).FieldDisplay(func(value types.FieldModel) interface{} {
		return value.Row["courses_goadmin_join_description"]
	})
	if user.CheckRole("student") {
		components := tmpl.Default(ctx)
		lHtml := components.Col().SetSize(types.SizeMD(2)).SetContent("").GetContent()
		rHtml := components.Col().SetSize(types.SizeMD(10)).SetContent("&nbsp;&nbsp;&nbsp;&nbsp;" + "<a href=\"/admin/info/select_course/new\" class=\"btn btn-primary\">选择课程</a>\n").GetContent()
		components.Col().SetContent(lHtml + rHtml).GetContent()
		detail.SetFooterHtml(components.Row().SetContent(lHtml + rHtml).GetContent())
	}
	//iface := services.GetInterfaceByName("WLAN")
	//ip := services.GetIPv4Addresses(iface)
	//fmt.Println(ip)

	if user.CheckRole("teacher") {
		tea_name := services.GetTeacherID(user)
		//fmt.Println(tea_name)
		info.Where("tea_id", "=", services.TransItoStr(tea_name))
	}

	info.SetTable("course_teacher").SetTitle("课程与教师").SetDescription("管理课程与教师间的信息").SetActionButtonFold()

	formList := courseTeacher.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("教师姓名", "tea_id", db.Int, form.SelectSingle).FieldOptions(services.TransFieldOptions(services.GetTeachers(), "name", "id")).FieldMust()
	formList.AddField("课程名称", "course_id", db.Int, form.SelectSingle).FieldOptions(services.TransFieldOptions(services.GetCourses(), "title", "id")).FieldMust()
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()
	formList.SetTable("course_teacher").SetTitle("课程与教师").SetDescription("更新课程与教师间的信息")

	return courseTeacher
}
