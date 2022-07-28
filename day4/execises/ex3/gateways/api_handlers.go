package gateways

import (
	"github.com/emicklei/go-restful/v3"
	"http-rest/v/daos"
	"http-rest/v/domain"
	"io/ioutil"
	"log"
	"net/http"
)

var bookPath = "/books"

type API struct {
	users map[int]interface{}
}

func NewAPI() *API {
	return &API{
		users: make(map[int]interface{}),
	}
}

func (api *API) RegisterRoutes(ws *restful.WebService) {
	ws.Path("/my-app")
	ws.Route(ws.GET(bookPath).To(api.echoGETHandler).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON))
	ws.Route(ws.GET(bookPath + "/{name}").To(api.echoGETHandlerBookName).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON))
	ws.Route(ws.PUT(bookPath).To(api.echoPUTHandler).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON))
	ws.Route(ws.POST(bookPath).To(api.echoPOSTHandler).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON))
	ws.Route(ws.PATCH(bookPath).To(api.echoPATCHHandler).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON))
	ws.Route(ws.DELETE(bookPath).To(api.echoDELETEHandler).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON))
}

func (api *API) echoGETHandler(req *restful.Request, resp *restful.Response) {
	answer := daos.GetAll()
	var data []domain.Book
	i := 0
	for _, v := range answer {
		data[i] = v
		i++
	}
	resp.WriteAsJson(data)
}

func (api *API) echoGETHandlerBookName(req *restful.Request, resp *restful.Response) {
	answer, err := daos.GetBook(req.PathParameter("name"))
	if err != nil {
		resp.WriteAsJson(answer)
	} else {
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
	}
}

func (api *API) echoPUTHandlerBookName(req *restful.Request, resp *restful.Response) {
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

	/*daos.AddBook(data)

	if err != nil {
		resp.WriteAsJson(answer)
	} else {
		resp.WriteServiceError(http.StatusInternalServerError, restful.NewError(http.StatusInternalServerError, err.Error()))
	}*/
}
