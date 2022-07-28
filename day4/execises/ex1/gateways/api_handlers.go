package gateways

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const echoPath = "/echo"
const toolsPath = "/tools/repeater"
const fiboPath = "/tools/fibo/{n}"
const userPath = "/user/{user}"

type API struct {
	users map[int]interface{}
}

func NewAPI() *API {
	return &API{
		users: make(map[int]interface{}),
	}
}

var user_data = make(map[string]interface{})

func (api *API) RegisterRoutes(ws *restful.WebService) {
	// ex 2
	ws.Path("/my-app")
	ws.Route(ws.PUT(echoPath).To(api.echoPUTHandler).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON))
	ws.Route(ws.PATCH(echoPath).To(api.echoPATCHHandler).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON))
	ws.Route(ws.DELETE(echoPath).To(api.echoDELETEHandler).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON))
	ws.Route(ws.POST(echoPath).To(api.echoPOSTHandler).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON))
	ws.Route(ws.GET(echoPath).To(api.echoGETHandler).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON))

	// ex 2
	ws.Route(ws.POST(toolsPath).To(api.echoPOSTHandlerForRepeater).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON))
	ws.Route(ws.GET(fiboPath).To(api.echoGETHandlerForFibo).Param(ws.PathParameter("n", "fibonacci n number").DataType("int")))

	ws.Route(ws.GET(userPath).To(api.echoGETHandlerForUser).Param(ws.PathParameter("user", "user name").DataType("string")))
}

func calculateFibonacci(n int) int {
	x, y := 0, 1
	var i int
	for i = 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func (api *API) echoGETHandlerForUser(req *restful.Request, resp *restful.Response) {
	name := req.PathParameter("user")

	preferences := req.HeaderParameter("USER-IDENTIFIER")

	if preferences == "" {
		id := uuid.New()
		resp.Header().Add("USER-IDENTIFIER", string(id[:]))
		user_data[name] = ""
	} else {
		user_data[name] = preferences
	}

	resp.WriteAsJson(
		user_data[name],
	)
}

func (api *API) echoGETHandlerForFibo(req *restful.Request, resp *restful.Response) {
	number := req.PathParameter("n")

	realNumber, err := strconv.Atoi(number)
	if err != nil {
		log.Printf("[ERROR] An error has occured while converting the number")
	}

	resp.WriteAsJson(
		calculateFibonacci(realNumber),
	)
}

func (api *API) echoPOSTHandlerForRepeater(req *restful.Request, resp *restful.Response) {
	limit := req.QueryParameter("limit")
	strVal := req.QueryParameter("string")

	value, err := strconv.Atoi(limit)
	if err != nil {
		log.Printf("An error has occured while converting limit: %v", err)
	} else {
		log.Printf("Value of limit is: %d", value)
	}

	data := strings.Repeat(strVal, value)

	resp.WriteAsJson(
		data,
	)
}

func (api *API) echoPOSTHandler(req *restful.Request, resp *restful.Response) {
	body := req.Request.Body
	if body == nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusNotFound, restful.NewError(http.StatusInternalServerError, "nil body"))
		return

	}
	defer body.Close()
	var err error
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	resp.WriteAsJson(map[string]string{
		"POST": string(data),
	})
}

func (api *API) echoGETHandler(req *restful.Request, resp *restful.Response) {
	//param := req.QueryParameter("")
	body := req.Request.Body
	if body == nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusNotFound, restful.NewError(http.StatusInternalServerError, "nil body"))
		return

	}
	defer body.Close()
	var err error
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	resp.WriteAsJson(map[string]string{
		"GET": string(data),
	})
}

func (api *API) echoPUTHandler(req *restful.Request, resp *restful.Response) {
	//param := req.QueryParameter("")
	body := req.Request.Body
	if body == nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusNotFound, restful.NewError(http.StatusInternalServerError, "nil body"))
		return

	}
	defer body.Close()
	var err error
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	resp.WriteAsJson(map[string]string{
		"PUT": string(data),
	})
}

func (api *API) echoPATCHHandler(req *restful.Request, resp *restful.Response) {
	//param := req.QueryParameter("")
	body := req.Request.Body
	if body == nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusNotFound, restful.NewError(http.StatusInternalServerError, "nil body"))
		return

	}
	defer body.Close()
	var err error
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	resp.WriteAsJson(map[string]string{
		"PATCH": string(data),
	})
}

func (api *API) echoDELETEHandler(req *restful.Request, resp *restful.Response) {
	//param := req.QueryParameter("")
	body := req.Request.Body
	if body == nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusNotFound, restful.NewError(http.StatusInternalServerError, "nil body"))
		return

	}
	defer body.Close()
	var err error
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("[ERROR] Couldn't read request body")
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	resp.WriteAsJson(map[string]string{
		"DELETE": string(data),
	})
}
