package controllers

import (
	"fmt"
	"math/rand"
	
	//	"github.com/astaxie/beego"
	//	"net/http/httputil"

	"github.com/cuu/select_tags_admin/models"
)

type MainController struct {
	BaseController
}
	
// @router / [get]
func (this *MainController) Index() {
	
	var dishes []*models.Dish
	var dishes_randed []*models.Dish
	qs := models.Dishes().OrderBy("-Created").Limit(100).RelatedSel()

	models.ListObjects(qs,&dishes)

	random_length := 2
	
	if len(dishes) >= random_length {
		randoms := rand.Perm(len(dishes))[:random_length]
		
		for i:=0; i< len(randoms); i++ {
			dishes_randed = append(dishes_randed,dishes[ randoms[i] ] )
		}
		
	}else {
		dishes_randed = dishes
	}

	this.Data["Dishes"] = dishes_randed
	
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	
	this.TplName = "index.tpl"
	this.Render()
	
}

// @router / [post]
func (this *MainController) IndexPost() {
	this.Data["Website"] = "beego.me POST"
	this.Data["Email"] = "astaxie@gmail.com"


	fmt.Println(this.RequestDump())

	PostData := fmt.Sprintf("%s", this.Input())

	this.Data["PostData"] = PostData
	
	this.TplName = "index.tpl"
	this.Render()

}
