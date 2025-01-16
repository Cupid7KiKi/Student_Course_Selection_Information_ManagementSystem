package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetStudentsTable(ctx *context.Context) table.Table {

	students := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := students.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int)
	info.AddField("姓名", "name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("性别", "gender", db.Varchar).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).
		FieldFilterOptions(types.FieldOptions{{Value: "男", Text: "男"}, {Value: "女", Text: "女"}})
	info.AddField("专业", "major", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldHide()
	info.AddField("学院", "faculty", db.Varchar).
		FieldHide()
	info.AddField("班级", "class", db.Varchar)
	info.AddField("学号", "std_num", db.Varchar).
		FieldHide()
	info.AddField("创建时间", "created_at", db.Timestamp).
		FieldHide()

	info.SetTable("students").SetTitle("学生管理").SetDescription("管理学生信息")

	detail := students.GetDetail()
	detail.AddField("ID", "id", db.Int)
	detail.AddField("姓名", "name", db.Varchar)
	detail.AddField("性别", "gender", db.Varchar)
	detail.AddField("专业", "major", db.Varchar)
	detail.AddField("学院", "faculty", db.Varchar)
	detail.AddField("班级", "class", db.Varchar)
	detail.AddField("学号", "std_num", db.Varchar)

	formList := students.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldHide()
	formList.AddField("姓名", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("性别", "gender", db.Varchar, form.SelectSingle).
		// 单选的选项，text代表显示内容，value代表对应值
		FieldOptions(types.FieldOptions{
			{Text: "男", Value: "男"},
			{Text: "女", Value: "女"},
		}).
		// 设置默认值
		FieldDefault("男").FieldMust()
	formList.AddField("专业", "major", db.Varchar, form.Text).FieldMust()
	formList.AddField("学院", "faculty", db.Varchar, form.Text).FieldMust()
	formList.AddField("班级", "class", db.Varchar, form.Text).FieldMust()
	formList.AddField("学号", "std_num", db.Varchar, form.Text).FieldMust()
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("students").SetTitle("学生管理").SetDescription("更新学生信息")

	return students
}
