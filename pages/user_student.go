package pages

import (
	"Student_Course_Selection_Information_ManagementSystem/services"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUserstudentTable(ctx *context.Context) table.Table {

	userStudent := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := userStudent.GetInfo()

	info.AddField("ID", "id", db.Int)
	info.AddField("学生姓名", "name", db.Int).
		FieldJoin(types.Join{
			Table:     "students", // 连表的表名
			Field:     "std_id",   // 要连表的字段
			JoinField: "id",       // 连表的表的字段
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return value.Row["students_goadmin_join_name"]
		})
	info.AddField("用户名", "username", db.Int).
		FieldJoin(types.Join{
			Table:     "goadmin_users", // 连表的表名
			Field:     "user_id",       // 要连表的字段
			JoinField: "id",            // 连表的表的字段
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return value.Row["goadmin_users_goadmin_join_username"]
		})
	info.AddField("创建时间", "created_at", db.Timestamp)

	info.SetTable("user_student").SetTitle("学生用户").SetDescription("关联学生和用户的中间表")

	formList := userStudent.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldHide()
	formList.AddField("学生姓名", "std_id", db.Int, form.SelectSingle).FieldOptions(services.TransFieldOptions(services.GetStudents(), "name", "id")).FieldMust()
	formList.AddField("用户名", "user_id", db.Int, form.SelectSingle).FieldOptions(services.TransFieldOptions(services.GetUserName(), "name", "id")).FieldMust()
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("user_student").SetTitle("学生用户").SetDescription("关联学生和用户的中间表")

	return userStudent
}
