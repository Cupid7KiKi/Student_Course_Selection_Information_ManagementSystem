package pages

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"regexp"
)

func GetTeachersTable(ctx *context.Context) table.Table {

	teachers := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := teachers.GetInfo()

	info.AddField("ID", "id", db.Int)
	info.AddField("姓名", "name", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("性别", "gender", db.Varchar).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).
		FieldFilterOptions(types.FieldOptions{{Value: "男", Text: "男"}, {Value: "女", Text: "女"}})
	info.AddField("职称", "position", db.Varchar)
	info.AddField("手机号码", "phone_number", db.Varchar).
		FieldHide()
	info.AddField("创建时间", "created_at", db.Timestamp).
		FieldHide()

	info.SetTable("teachers").SetTitle("教师管理").SetDescription("管理教师信息")

	formList := teachers.GetForm()
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
	formList.AddField("职称", "position", db.Varchar, form.Text).FieldMust()
	//定义正则表达式
	regex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	formList.AddField("手机号码", "phone_number", db.Varchar, form.Text).FieldMust().
		SetPostValidator(func(values form2.Values) error {
			if !regex.MatchString(values.Get("phone_number")) {
				return fmt.Errorf("您输入的手机号码有误！！！")
			}
			return nil
		})
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("teachers").SetTitle("教师管理").SetDescription("更新教师信息")

	return teachers
}
