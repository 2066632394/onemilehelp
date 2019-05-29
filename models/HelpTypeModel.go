package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

//互助类型表
type HelpType struct {
	HelpTypeId            	int64
	HelpTypeName      	string    `orm:"unique;size(64)" form:"HelpTypeName"  valid:"Required"`
}

func (u *HelpType) TableName() string {
	return beego.AppConfig.String("help_type_table")
}

func (u *HelpType) Valid(v *validation.Validation) {

}

func init() {
	orm.RegisterModel(new(HelpType))
}

/************************************************************/

//get type list
func GetHelpTypelist(page int64, page_size int64, sort string) (helps []orm.Params, count int64) {
	o := orm.NewOrm()
	h := new(HelpType)
	qs := o.QueryTable(h)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&helps)
	count, _ = qs.Count()
	return helps, count
}

//添加类型
func AddHelpType(u *HelpType) (int64, error) {
	o := orm.NewOrm()
	h := new(HelpType)
	h.HelpTypeName = u.HelpTypeName

	id, err := o.Insert(h)
	return id, err
}

//更新用户
func UpdateTypeName(u *HelpType) (int64, error) {
	o := orm.NewOrm()
	h := make(orm.Params)
	if len(u.HelpTypeName) > 0 {
		h["HelpTypeName"] = u.HelpTypeName
	}

	var table HelpType
	num, err := o.QueryTable(table).Filter("HelpTypeId", u.HelpTypeId).Update(h)
	return num, err
}

func DelUserById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&HelpType{HelpTypeId: Id})
	return status, err
}

func GetUserById(id int64) (help HelpType) {
	help = HelpType{HelpTypeId: id}
	o := orm.NewOrm()
	o.Read(&help, "HelpTypeId")
	return help
}

