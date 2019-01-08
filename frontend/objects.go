package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ObjectPrefixes []string

// ObjectsResponse returned on a successful GET of objects
// swagger:response
type ObjectsResponse struct {
	// in: body
	Body ObjectPrefixes
}

func (f *Frontend) InitObjectsApi() {
	// swagger:route GET /objects Objects listObjects
	//
	// Lists the object types in the system
	//
	//     Responses:
	//       200: ObjectsResponse
	//       401: NoContentResponse
	//       403: NoContentResponse
	f.ApiGroup.GET("/objects",
		func(c *gin.Context) {
			objPrefixes := f.dt.GetObjectTypes()
			c.JSON(http.StatusOK, objPrefixes)
		})
}
