package pages

import (
	"Student_Course_Selection_Information_ManagementSystem/services"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/color"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
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
	//info.AddButton(ctx, "选择课程", "fa-plus", action.Jump("/admin/info/select_course/new?__page=1&__pageSize=10&__sort=id&__sort_type=desc"), color.Red)
	//
	//info.HideEditButton()
	//info.AddField("选择课程", "选择课程", db.Varchar).
	//	FieldDisplay(func(value types.FieldModel) interface{} {
	//		return template.Default().
	//			Link().
	//			SetURL("/admin/info/select_course/new").
	//			SetContent(template2.HTML(value.Value)).
	//			OpenInNewTab().
	//			SetTabTitle(template.HTML("Author Detail(" + value.Value + ")")).
	//			GetContent()
	//	})
	// 第一个参数为标题，第二个参数为对应的操作
	//info.AddActionButton(ctx, "操作", action.Jump("https://baidu.com"))
	//if services.GetUserRole(user) == "operator" {
	//
	//}

	//成功案例
	//info.AddButton(ctx, "测试", icon.Android, action.Ajax("/admin/ajax",
	//	func(ctx *context.Context) (success bool, msg string, data interface{}) {
	//		return true, "请求成功，奥利给", ""
	//	}))

	//第一个参数为标题，第二个参数为对应的操作
	//info.AddColumnButtons(ctx, "选择课程", types.GetColumnButton("新建课程", icon.Info,
	//	action.Jump("select_course/new"))).FieldHide()
	//info.Where("tea_id", "=", "1")

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
	iface := services.GetInterfaceByName("WLAN")
	ip := services.GetIPv4Addresses(iface)
	//fmt.Println("是否拥有学生权限,", user.CheckRole("student"))
	if user.CheckRole("student") {
		//detail.AddColumnButtons(ctx, "选择课程", types.GetColumnButton("新建课程", icon.Info,
		//	action.Jump("http://127.0.0.1:9022/admin/info/select_course/new")))

		detail.AddButton(ctx, "选择课程", "fa-plus", action.Jump("/admin/info/select_course/new?__page=1&__pageSize=10&__sort=id&__sort_type=desc"), color.Red)
		detail.HideEditButton()
		//component := tmpl.Template.Col().SetSize(types.SizeMD(9)),)
	}

	detail.AddButton(ctx, "选择课程", "fa-plus", action.Jump("/admin/info/select_course/new?__page=1&__pageSize=10&__sort=id&__sort_type=desc"), color.Red)
	detail.HideNewButton()

	fmt.Println(ip)
	//action.Jump("http://"+ip+":9022/admin/info/select_course/new")
	//detail.AddField("选择课程","choose",db)
	//detail.AddField("123", "123", db.Int)

	if user.CheckRole("teacher") {
		tea_name := services.GetTeacherID(user)
		fmt.Println(tea_name)
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
