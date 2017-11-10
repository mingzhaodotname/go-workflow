package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:WfTemplateController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:WfTemplateController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:WfTemplateController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:WfTemplateController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:WfTemplateController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:WfTemplateController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:WfTemplateController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:WfTemplateController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:WfTemplateController"] = append(beego.GlobalControllerRouter["github.com/mingzhaodotname/go-workflow/wf-server/controllers:WfTemplateController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
