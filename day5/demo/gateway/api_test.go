package gateway_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"lab05/domain/mocks"
	"lab05/gateway"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/emicklei/go-restful/v3"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var api *gateway.API
var my_api *gateway.MyAPI

func newJSONRequest(method, url string, obj interface{}) *http.Request {
	var payloadReader io.Reader

	if obj != nil {
		payload, err := json.Marshal(obj)
		Expect(err).To(BeNil())
		payloadReader = bytes.NewBuffer(payload)
	}
	req, err := http.NewRequest(method, url, payloadReader)
	Expect(err).To(BeNil())
	req.Header.Set("content-type", "application/json")
	return req
}

func readJSONResponse(in io.Reader, v interface{}) {
	respBody, err := ioutil.ReadAll(in)
	Expect(err).To(BeNil())
	fmt.Println(string(respBody))
	Expect(json.Unmarshal(respBody, v)).To(BeNil())
}

var _ = Describe("MyAPI Handler", func() {
	var (
		ws             *restful.WebService
		container      *restful.Container
		recorder       *httptest.ResponseRecorder
		my_storageMock *mocks.MockMyStorage
	)

	BeforeEach(func() {
		my_storageMock = mocks.NewMockMyStorage(gomock.NewController(&testing.B{}))
		my_api = gateway.NewAPIMyStorage(my_storageMock)
		ws = new(restful.WebService)
		container = restful.NewContainer()
		container.Add(ws)
		recorder = httptest.NewRecorder()
		my_api.RegisterMyRoutes(ws)
	})

	// frumos avem aici un context
	Context("When getting a my book", func() {
		// daca n-avem autor e belea
		It("does not exists should return 404", func() {
			// facem aici un get
			request, _ := http.NewRequest("GET", "/my-app/mybooks/5", nil)

			// si servim frumos aici
			container.ServeHTTP(recorder, request)
			// si sa vedem ce ne arata
			Expect(recorder.Code).To(Equal(http.StatusNotFound))
		})

		// daca n-avem titlu e belea
		It("should fail if does not respect endpoint signature", func() {
			// facem aici frumos un request
			request, _ := http.NewRequest("GET", "/my-app/mybooks?author=Becali", nil)

			// si servim frumos aici
			container.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(405))
		})
	})

	// salvam frumos o carte
	Context("When saving a new book", func() {
		// trb sa dea fail daca nu merge pe storage
		It("Should fail if store fails", func() {
			// body de il vrem aici frumos sa-l adaugam
			bodyContent := "asdfg"
			// request frumos de put
			request, err := http.NewRequest("PUT", "/my-app/mybooks/store/1", strings.NewReader(bodyContent))
			// da asa e
			Expect(err).To(BeNil())

			// acum mockuim storage ul sa ne dea fail
			my_storageMock.EXPECT().WriteMyContent("1", bodyContent).Return(fmt.Errorf("fail"))
			// si servim
			container.ServeHTTP(recorder, request)
			// si trb sa ne dea eroare
			Expect(recorder.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})

var _ = Describe("API Handler", func() {
	var (
		ws          *restful.WebService
		container   *restful.Container
		recorder    *httptest.ResponseRecorder
		storageMock *mocks.MockStorage
	)

	BeforeEach(func() {
		storageMock = mocks.NewMockStorage(gomock.NewController(&testing.B{}))
		api = gateway.NewAPI(storageMock)
		ws = new(restful.WebService)
		container = restful.NewContainer()
		container.Add(ws)
		recorder = httptest.NewRecorder()
		api.RegisterRoutes(ws)
	})

	Context("When adding a new book", func() {
		It("should fail as no author is mentioned in query parameters", func() {
			request, _ := http.NewRequest("GET", "/book-app/books", nil)

			container.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(http.StatusBadRequest))
		})

		It("should fail as no title is mentioned in query parameters", func() {
			request, _ := http.NewRequest("GET", "/book-app/books?author=Becali", nil)

			container.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(http.StatusBadRequest))
		})
	})

	Context("When saving a new book", func() {
		It("Should fail if store fails", func() {
			bodyContent := "asdfg"
			request, err := http.NewRequest("PUT", "/book-app/books/store/1", strings.NewReader(bodyContent))
			Expect(err).To(BeNil())

			storageMock.EXPECT().WriteContent("1", bodyContent).Return(fmt.Errorf("fail"))
			container.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
