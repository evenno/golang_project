package blog

import (
	//"fmt"

	//"fmt"

	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
	options      map[string]string
	right string
	page int 
	pagesize int
}

//  在所有方法之前加载
func (this *baseController) Prepare(){
	this.Data["hidejs"] = `<!--[if lt IE 9]>
	<script src="/static/js/html5shiv.min.js"></script>
	<![endif]-->`

	var (
		pagesize int 
		err error 
		page int 
	)
	if page,err = strconv.Atoi(this.Ctx.Input.Param(":page")); err !=nil || page < 1 {
		page = 1
	}

	if pagesize, err = strconv.Atoi(this.getOption("pagesize")); err != nil || pagesize < 1 {
		pagesize = 10
	}
	this.page=page 
	this.pagesize=pagesize
}

func (this *baseController)getOption(name string)string {
	if v,ok:=this.options[name];ok{
		return v
	}else {
		return ""
	}
}


func (this *baseController) setHeadMetas(params ...string) {
	title_buf := make([]string, 0, 3)
	if len(params) == 0 && this.getOption("sitename") != "" {
		title_buf = append(title_buf, this.getOption("sitename"))
	}
	if len(params) > 0 {
		title_buf = append(title_buf, params[0])
	}
	title_buf = append(title_buf, this.getOption("subtitle"))
	this.Data["title"] = strings.Join(title_buf, " - ")

	if len(params) > 1 {
		this.Data["keywords"] = params[1]
	} else {
		this.Data["keywords"] = this.getOption("keywords")
	}

	if len(params) > 2 {
		this.Data["description"] = params[2]
	} else {
		this.Data["description"] = this.getOption("description")
	}
}

func (this *baseController) display(tpl string){
	theme := "double"
	// 类型断言
	if v, ok := this.options["theme"]; ok && v != "" {
		theme = v
	}
	this.Layout=theme + "/layout.html"
	this.Data["root"]="/"+beego.BConfig.WebConfig.ViewsPath + "/" + theme + "/"
	this.TplName=theme+"/" + tpl +".html"
	this.LayoutSections=make(map[string]string)
	this.LayoutSections["head"]=theme+"/head.html"
	if tpl == "index" {
		this.LayoutSections["banner"] = theme + "/banner.html"
	}
	if this.right != "" {
		this.LayoutSections["right"]=theme + "/" +this.right
	}
	this.LayoutSections["foot"]=theme+"/foot.html"
}