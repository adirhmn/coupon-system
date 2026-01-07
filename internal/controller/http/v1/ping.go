package v1

import (
	serverctrl "coupon-system/internal/controller/http"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	modeltransport "coupon-system/internal/entity/transport"
)

// PING godoc
// @Summary Ping the server
// @Description Ping the server
// @Tags ping
// @Produce json
// @Success 200 {object} transport.StandardResponse
// @Router /v1/ping [get]
func (v1 *v1Controller) Ping(c *gin.Context) {
	ctx := c.Request.Context()

	pingPong, err := v1.pingService.Ping(ctx)
	if err != nil {
		serverctrl.ResponseHandler(c, http.StatusInternalServerError, modeltransport.PingResponse{
			ServerSays: pingPong.Message,
		}, fmt.Errorf("internal server error"))
		return
	}

	serverctrl.ResponseHandler(c, http.StatusOK, modeltransport.PingResponse{
		ServerSays: pingPong.Message,
	}, nil)
}