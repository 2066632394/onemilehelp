package models

import (
	"errors"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	. "github.com/beego/admin/src/lib"
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
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
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
	if err := checkUser(u); err != nil {
		return 0, err
	}
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
		h["Username"] = u.Username
	}
	if len(u.Nickname) > 0 {
		user["Nickname"] = u.Nickname
	}
	if len(u.Email) > 0 {
		user["Email"] = u.Email
	}
	if len(u.Remark) > 0 {
		user["Remark"] = u.Remark
	}
	if len(u.Password) > 0 {
		user["Password"] = Strtomd5(u.Password)
	}
	if u.Status != 0 {
		user["Status"] = u.Status
	}
	if len(user) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table User
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
	return num, err
}

func DelUserById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&User{Id: Id})
	return status, err
}

func GetUserByUsername(username string) (user User) {
	user = User{Username: username}
	o := orm.NewOrm()
	o.Read(&user, "Username")
	return user
}

func GetUserById(id int64) (user User) {
	user = User{Id: id}
	o := orm.NewOrm()
	o.Read(&user, "Id")
	return user
}

