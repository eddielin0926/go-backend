package routers_test

import (
	"encoding/json"
	"go-backend/routers"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	var app *gin.Engine
	var rec *httptest.ResponseRecorder
	gin.SetMode(gin.ReleaseMode)

	BeforeEach(func() {
		app = gin.New()
		routers.SetupRouter(app)
		rec = httptest.NewRecorder()
	})

	It("should return user", func() {
		var data string
		req, _ := http.NewRequest("GET", "/api/v1/user/", nil)

		app.ServeHTTP(rec, req)
		json.Unmarshal(rec.Body.Bytes(), &data)

		Expect(rec.Code).To(Equal(http.StatusOK))
		Expect(data).To(Equal("user"))
	})
})
