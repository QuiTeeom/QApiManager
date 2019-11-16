package api_manager

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CreateHandler(mux *http.ServeMux) {
	addController(mux, NewUserController())
	addController(mux, NewApiController())
}

func addController(mux *http.ServeMux, controller HttpController) {
	for _, method := range controller.methods {
		mux.HandleFunc(controller.path+method.path, func(writer http.ResponseWriter, request *http.Request) {
			defer func() {
				err := recover()
				if err != nil {
					fmt.Println(err)
					json, _ := json.Marshal(&HttpResponse{
						Success: false,
						Message: err.(string),
					})
					writer.Header().Set("Content-Type", "application/json")
					writer.Write(json)
				}
			}()

			fn := method.methodMap[strings.ToLower(request.Method)]
			if fn != nil {
				ctx := &QtRequestContext{
					ResponseWriter: writer,
					Request:        request,
				}

				fn(ctx)

				json, _ := json.Marshal(&HttpResponse{
					Success: true,
					Data:    ctx.responseObj,
				})
				writer.Header().Set("Content-Type", "application/json")
				writer.Write(json)
			}
		})
	}
}

type HttpController struct {
	path    string
	methods []HttpControllerMethod
}
type HttpControllerMethod struct {
	path      string
	method    string
	methodMap map[string]func(context *QtRequestContext)
}
type QtRequestContext struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	responseObj    interface{}
}

func (cxt *QtRequestContext) getJson(o interface{}) {
	bodyLength := cxt.Request.ContentLength
	body := make([]byte, bodyLength)
	r := cxt.Request.Body
	r.Read(body)
	json.Unmarshal(body, o)
}

func (cxt *QtRequestContext) returnObj(o interface{}) {
	cxt.responseObj = o
}

type HttpResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewUserController() HttpController {
	return HttpController{
		path: "/users",
		methods: []HttpControllerMethod{
			{
				path: "",
				methodMap: map[string]func(context *QtRequestContext){
					"get": func(context *QtRequestContext) {
						//user := User{}
						//res := make([]User,1)
						//db.Select(user,&res)
						//fmt.Fprint(writer,json.Marshal(&res))
					},
					"post": func(context *QtRequestContext) {
						//request.ParseForm()
						//userName := request.FormValue("userName")
						//
						//user := User{}
						//var count int
						//db.Find(&user,&User{UserName:userName}).Count(&count)
						//
						//if count==0 {
						//	user = User{
						//		UserName:request.FormValue("userName"),
						//	}
						//	db.Create(&user)
						//}
					},
				},
			},
		},
	}
}

func NewApiController() HttpController {
	return HttpController{
		path: "/apis",
		methods: []HttpControllerMethod{
			{
				path: "",
				methodMap: map[string]func(*QtRequestContext){
					"post": func(cxt *QtRequestContext) {
						api := Api{}
						cxt.getJson(&api)

						var count int
						db.Table("apis").Where("code = '" + api.Code + "'").Count(&count)
						if count > 0 {
							panic("code 已经存在")
						}
						db.Create(api)
					},
					"put": func(cxt *QtRequestContext) {
						api := Api{}
						cxt.getJson(&api)
						db.Table("apis").Where("code = '" + api.Code + "'").Update(&api)
					},
					"get": func(cxt *QtRequestContext) {
						collectionCode := cxt.Request.FormValue("collectionCode")
						var apis []Api
						db.Where("collection_code='" + collectionCode + "'").Find(&apis)
						cxt.returnObj(apis)
					},
				},
			},
		},
	}
}
