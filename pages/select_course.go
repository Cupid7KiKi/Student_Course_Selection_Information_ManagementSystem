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
	selection "github.com/GoAdminGroup/go-admin/template/types/form/select"
)

func GetSelectcourseTable(ctx *context.Context) table.Table {

	selectCourse := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	user := auth.Auth(ctx)

	fmt.Println(user.Name)
	fmt.Println(user.Permissions[0].Slug)
	// 验证其权限
	if !user.CheckPermission("Select_Course") && !user.CheckPermission("*") {
		fmt.Println("No Permission")
	} else {
		fmt.Println("获取权限成功")
	}

	info := selectCourse.GetInfo().HideFilterArea()
	fmt.Println(user.CheckRole("administrator"))
	fmt.Println(user.Roles[0].Slug)
	fmt.Println(user.Id)
	if user.CheckRole("student") {
		//预查询仅能查看自己
		std_name := services.GetStudentName(user)
		fmt.Println(std_name)

		info.Where("std_id", "=", services.TransItoStr(std_name))

	}

	info.AddButton(ctx, "选课", "fa-plus", action.Jump("/admin/info/select_course/new?__page=1&__pageSize=10&__sort=id&__sort_type=desc"), color.LightGreen)
	info.HideNewButton()
	info.AddField("ID", "id", db.Int)
	info.AddField("学生姓名", "name", db.Int).
		FieldJoin(types.Join{
			Table:     "students", // 连表的表名
			Field:     "std_id",   // 要连表的字段
			JoinField: "id",       // 连表的表的字段
		}).
		//筛选时支持模糊查询
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return value.Row["students_goadmin_join_name"]
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
	info.AddField("授课教师", "name", db.Int).
		FieldJoin(types.Join{
			Table:     "teachers",
			Field:     "t_id",
			JoinField: "id",
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return value.Row["teachers_goadmin_join_name"]
		})
	info.AddField("创建时间", "created_at", db.Timestamp).
		FieldHide()

	//成功案例
	info.AddActionButton(ctx, "选课", action.Jump("select_course/new"), color.LightGreen)

	info.SetTable("select_course").SetTitle("学生选课管理").SetDescription("管理学生选课")

	//detail := selectCourse.GetDetail()
	//detail.AddField("ID", "id", db.Int)
	//detail.AddField("学生姓名", "name", db.Int)
	//detail.AddField("课程名称", "title", db.Int)
	//detail.AddField("授课教师", "t_id", db.Int)
	//detail.AddField("创建时间", "created_at", db.Timestamp)
	//detail := selectCourse.GetDetailFromInfo()
	//detail.AddField("授课教师", "name", db.Varchar)
	formList := selectCourse.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldHide()
	if user.CheckRole("administrator") {
		formList.AddField("学生姓名", "std_id", db.Int, form.SelectSingle).FieldOptions(services.TransFieldOptions(services.GetStudents(), "name", "id")).FieldMust()
	}
	fmt.Println("测试：", user.CheckRole("student"))
	if user.CheckRole("student") {
		formList.AddField("学生姓名", "std_id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDefault(user.Name)
	}
	formList.AddField("课程名称", "course_id", db.Varchar, form.SelectSingle).FieldOptions(services.TransFieldOptions(services.GetCourses(), "title", "id")).FieldMust().FieldOnChooseAjax("t_id", "choose/course_id",
		func(ctx *context.Context) (bool, string, interface{}) {
			c_id := ctx.FormValue("value")
			tea := services.GetCourseT(c_id)
			var teas []interface{}
			// 遍历切片
			for _, item := range tea {
				// 获取每个 map 的键和值
				for _, value := range item {
					teas = append(teas, value)
				}
			}
			var data = make(selection.Options, len(teas))
			data = services.TransSelectionOptions(services.GetTeacherName(teas), "name", "id")
			return true, "ok", data
		})
	formList.AddField("授课教师", "t_id", db.Int, form.SelectSingle).
		FieldOptionInitFn(func(val types.FieldModel) types.FieldOptions {
			return types.FieldOptions{
				{Value: val.Value, Text: val.Value, Selected: true},
			}
		}).FieldMust()
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("select_course").SetTitle("学生选课管理").SetDescription("管理学生选课")

	return selectCourse
}
