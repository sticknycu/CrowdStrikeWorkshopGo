package gateway

import (
	"fmt"
	"io/ioutil"
	"lab05/domain"
	"net/http"

	"github.com/emicklei/go-restful/v3"

	log "github.com/sirupsen/logrus"
)

const (
	booksPath        = "/books"
	booksPathStore   = "/books/store/{id}"
	myBooksPath      = "/mybooks"
	myBooksPathStore = "/mybooks/store/{id}"
)

type API struct {
	books   map[int]domain.Book
	storage domain.Storage
}

type MyAPI struct {
	books      map[int]domain.Book
	my_storage domain.MyStorage
}

func NewAPI(storage domain.Storage) *API {
	return &API{
		books:   make(map[int]domain.Book),
		storage: storage}
}

func NewAPIMyStorage(my_storage domain.MyStorage) *MyAPI {
	return &MyAPI{
		books:      make(map[int]domain.Book),
		my_storage: my_storage,
	}
}

func (my_api *MyAPI) RegisterMyRoutes(ws *restful.WebService) {
	ws.Path("/my-app")
	ws.Route(ws.POST(myBooksPath).To(my_api.addMyBook).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON).Doc("Adding a new book in the database"))
	ws.Route(ws.GET(myBooksPathStore).To(my_api.getMyBook).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON).Doc("Getting a book from database"))
	ws.Route(ws.PUT(myBooksPathStore).To(my_api.saveMyBook).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON).Doc("Adding a new book in the database"))
}

func (api *API) RegisterRoutes(ws *restful.WebService) {
	ws.Path("/book-app")
	ws.Route(ws.POST(booksPath).To(api.addBook).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON).Doc("Adding a new book in the database"))
	ws.Route(ws.GET(booksPath).To(api.getBook).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON).Doc("Getting a book from database"))
	ws.Route(ws.PUT(booksPathStore).To(api.saveBook).Reads(restful.MIME_JSON).Writes(restful.MIME_JSON).Doc("Adding a new book in the database"))
}

func (my_api *MyAPI) saveMyBook(req *restful.Request, resp *restful.Response) {
	// avem aici in endpoint un path parameter boy
	id := req.PathParameter("id")
	// daca n avem e belea
	if id == "" {
		resp.WriteError(http.StatusBadRequest, fmt.Errorf("a valid id must be provided"))
		log.Errorf("Bad request. No valid id provided")
		return
	}
	// luam body ul frumos de l-am transmis
	dataReader := req.Request.Body
	// nasol daca nu i nici d-ala
	if dataReader == nil {
		resp.WriteError(http.StatusBadRequest, fmt.Errorf("no body provided"))
		log.Errorf("Bad request. No body provided")
		return
	}

	// si facem defer pe frumosu' de l-am luam
	defer dataReader.Close()

	// asa acum citim tot ce am luat
	data, err := ioutil.ReadAll(dataReader)
	// si hopa daca nu merge
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, fmt.Errorf("internal error"))
		log.WithError(err).Errorf("Failed to read body")
		return
	}
	// si daca e totul frumos
	log.Infof("Writing book with ID=%s and content=%v", id, data)

	// scriem la noi in storage
	err = my_api.my_storage.WriteMyContent(id, string(data))
	// si daca nu merge e belea iar
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, fmt.Errorf("internal error"))
		log.WithError(err).Errorf("Failed to write book to store")
		return
	}

	// scriem frumos raspunsul la ei in raspuns ca a mers
	_, err = resp.Write([]byte("Book saved"))
	// si belea daca nu merge
	if err != nil {
		log.WithError(err).Error("Failed to write response")
		_ = resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

func (api *API) saveBook(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	if id == "" {
		resp.WriteError(http.StatusBadRequest, fmt.Errorf("a valid id must be provided"))
		log.Errorf("Bad request. No valid id provided")
		return
	}
	dataReader := req.Request.Body
	if dataReader == nil {
		resp.WriteError(http.StatusBadRequest, fmt.Errorf("no body provided"))
		log.Errorf("Bad request. No body provided")
		return
	}
	defer dataReader.Close()

	data, err := ioutil.ReadAll(dataReader)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, fmt.Errorf("internal error"))
		log.WithError(err).Errorf("Failed to read body")
		return
	}
	log.Infof("Writing book with ID=%s and content=%v", id, data)

	err = api.storage.WriteContent(id, string(data))
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, fmt.Errorf("internal error"))
		log.WithError(err).Errorf("Failed to write book to store")
		return
	}

	_, err = resp.Write([]byte("Book saved"))
	if err != nil {
		log.WithError(err).Error("Failed to write response")
		_ = resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

func (my_api *MyAPI) addMyBook(req *restful.Request, resp *restful.Response) {
	// luam frumos obiectul nostru preferat
	book := &domain.Book{}
	// si citim frumos entitatea ca obiect
	err := req.ReadEntity(book)

	// si daca nu merge e belea
	if err != nil {
		log.WithError(err).Error("Failed to parse book json")
		_ = resp.WriteError(http.StatusBadRequest, err)
		return
	}

	// dupa accesam sa vedem daca o avem
	_, err = my_api.my_storage.GetMyContent(string(rune(book.GetBookHash())))
	// si daca e, e belea
	if err != nil {
		log.Error("Book already exists in the database")
		_ = resp.WriteError(http.StatusConflict, fmt.Errorf("book already exists"))
		return
	}

	// si daca nu e, frumos o punem
	my_api.books[book.GetBookHash()] = *book
	log.Info("Book added successfully in database")
}

func (api *API) addBook(req *restful.Request, resp *restful.Response) {
	book := &domain.Book{}
	err := req.ReadEntity(book)

	if err != nil {
		log.WithError(err).Error("Failed to parse book json")
		_ = resp.WriteError(http.StatusBadRequest, err)
		return
	}

	_, exists := api.books[book.GetBookHash()]
	if exists {
		log.Error("Book already exists in the database")
		_ = resp.WriteError(http.StatusConflict, fmt.Errorf("book already exists"))
		return
	}

	api.books[book.GetBookHash()] = *book
	log.Info("Book added successfully in database")
}

func (my_api *MyAPI) getMyBook(req *restful.Request, resp *restful.Response) {
	// avem aici niste query params frumosi de trb sa-i luam in seama
	id := req.PathParameter("id")

	// ca sa citim la noi in storage
	b, err := my_api.my_storage.GetMyContent(id)
	// si daca nu e, e belea
	if err != nil {
		log.Error("Book not found")
		_ = resp.WriteError(http.StatusNotFound, fmt.Errorf("book not found"))
		return
	}
	// si daca e, e de bine
	err = resp.WriteAsJson(b)
	// dar e belea daca nu i dam write response, sau pica el ca nu vrea el acolo ceva
	if err != nil {
		log.WithError(err).Error("Failed to write response")
		_ = resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

func (api *API) getBook(req *restful.Request, resp *restful.Response) {
	author := req.QueryParameter("author")
	title := req.QueryParameter("title")

	if author == "" {
		log.Error("Failed to read author")
		_ = resp.WriteError(http.StatusBadRequest, fmt.Errorf("book author must be provided"))
		return
	}

	if title == "" {
		log.Error("Failed to read title")
		_ = resp.WriteError(http.StatusBadRequest, fmt.Errorf("book title must be provided"))
		return
	}

	book := &domain.Book{
		Title:  title,
		Author: author,
	}

	b, ok := api.books[book.GetBookHash()]
	if !ok {
		log.Error("Book not found")
		_ = resp.WriteError(http.StatusNotFound, fmt.Errorf("book not found"))
		return
	}

	err := resp.WriteAsJson(b)
	if err != nil {
		log.WithError(err).Error("Failed to write response")
		_ = resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}
