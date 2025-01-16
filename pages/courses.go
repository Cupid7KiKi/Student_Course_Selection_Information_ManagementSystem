package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCoursesTable(ctx *context.Context) table.Table {

	courses := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := courses.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int)
	info.AddField("课程名称", "title", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("上课地点", "location", db.Varchar)
	info.AddField("学分", "credit", db.Varchar)
	info.AddField("课程简介", "description", db.Text).
		FieldHide()
	info.AddField("开始时间", "start_date", db.Timestamp).
		FieldHide()
	info.AddField("结束时间", "end_date", db.Timestamp).
		FieldHide()
	info.AddField("创建时间", "created_at", db.Timestamp).
		FieldHide()

	info.SetTable("courses").SetTitle("课程管理").SetDescription("管理课程信息")

	detail := courses.GetDetail()
	detail.AddField("ID", "id", db.Int)
	detail.AddField("课程名称", "title", db.Varchar)
	detail.AddField("上课地点", "location", db.Varchar)
	detail.AddField("学分", "credit", db.Varchar)
	detail.AddField("课程简介", "description", db.Text)
	detail.AddField("开始时间", "start_date", db.Timestamp)
	detail.AddField("结束时间", "end_date", db.Timestamp)

	formList := courses.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldHide()
	formList.AddField("课程名称", "title", db.Varchar, form.Text).FieldMust()
	formList.AddField("上课地点", "location", db.Varchar, form.Text).FieldMust()
	formList.AddField("学分", "credit", db.Varchar, form.Text).FieldMust()
	formList.AddField("课程简介", "description", db.Text, form.RichText)
	formList.AddField("开始时间", "start_date", db.Timestamp, form.Datetime).FieldMust()
	formList.AddField("结束时间", "end_date", db.Timestamp, form.Datetime).FieldMust()
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("courses").SetTitle("课程管理").SetDescription("管理课程信息")

	return courses
}
