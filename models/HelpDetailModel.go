package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
	"log"
)

//用户表
type HelpDetail struct {
	HelpId            	int64
	HelpTile      		string    `orm:"unique;size(64)" form:"HelpTile"  valid:"Required;MaxSize(64);MinSize(1)"`
	HelpDesc			string    `orm:"size(2024)" form:"HelpDesc" valid:"Required"`
	HelpUid				int64    `orm:"-" form:"HelpUid" valid:"Required"`
	HelpType			int64	 `valid:"Required"`
	HelpState			int64    `orm:"default(0)" form:"Status" valid:"Range(1,2)"`
	CreateTime    		int64
	UpdateTime    		int64		`valid:"Required"`
}

func (u *HelpDetail) TableName() string {
	return beego.AppConfig.String("help_detail_table")
}

func (u *HelpDetail) Valid(v *validation.Validation) {

}

func init() {
	orm.RegisterModel(new(HelpDetail))
}


//验证信息
func checkHelpDetail(h *HelpDetail) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&h)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

/************************************************************/

//get help list
func GetHelpDetailList(page int64, page_size int64, sort string) (list []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(HelpDetail)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&list)
	count, _ = qs.Count()
	return list, count
}

//添加互助
func AddHelpDetail(u *HelpDetail) (int64, error) {

	o := orm.NewOrm()
	h := new(HelpDetail)
	h.HelpTile = u.HelpTile
	h.HelpDesc = u.HelpDesc
	h.HelpUid = u.HelpUid
	h.HelpType = 1
	h.HelpState = 0
	h.CreateTime = time.Now().Unix()
	h.UpdateTime = time.Now().Unix()
	id, err := o.Insert(h)
	return id, err
}

//更新互助信息
func UpdateHelpDetail(u *HelpDetail) (int64, error) {

	o := orm.NewOrm()
	h := make(orm.Params)

	if len(h) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table HelpDetail
	num, err := o.QueryTable(table).Filter("HelpId", u.HelpId).Update(h)
	return num, err
}

func DelHelpDetailById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&HelpDetail{HelpId: Id})
	return status, err
}

func GetHelpDetailById(id int64) (h HelpDetail) {
	h = HelpDetail{HelpId: id}
	o := orm.NewOrm()
	o.Read(&h, "Id")
	return h
}

