package route

import (
	"github.com/golang191119/nc_crm/handler"
	"github.com/labstack/echo/v4"
)

func All(e *echo.Echo) {
	Private(e)
	Public(e)
}

func Private(e *echo.Echo) {

}

// mình sẽ hiện thực hợp đồng cho thuê xe nguyên chuyến như sau :
// - full price = base price + fee
// - base price tính theo khoảng cách, loại xe, khoảng cách lại chi theo nhiều mức khác nhau
// vd :
// 100 km đầu xe 10 tấn 17.000vnd/km, xe 5T 16.000vnd/km
// 200-500: xe 10T 16.500/km, xe 5T 15.400/km
// 500- 1000: xe 10T 16.000/km, xe 5T 15.000/km

// - fee không phụ thuộc khoảng cách và có nhiều cách tính khác nhau
// vd :
// phí cod 3% : thu thêm 3% * cod
// phí nâng hạ 1000vnd -> 1000 * weight
// phí đồng kiểm 2000vnd -> 2000 * số kiện hàng
// phí khai giá 2% : 2% & giá trị hàng hoá

// restful api, chỉnh sửa hợp đồng - bảng giá, tính giá
// -add/update contract
// -add/update/delete rate-card
// -get level
// -get truck type
// -get price

func Public(e *echo.Echo) {
	g := e.Group("/api/v1/public")
	g.GET("/health", handler.HealthCheck)
	// =================================== DONE =================================== //
	g.GET("/truck-levels", handler.GetTruckLevel)
	g.POST("/truck-levels", handler.AddTruckLevel)
	g.PUT("/truck-levels", handler.UpdateTruckLevel)
	g.GET("/truct-types", handler.GetTruckType)

	g.POST("/contract", handler.AddContract)
	g.GET("/contracts", handler.GetAllContracts)
	g.PUT("/contract", handler.UpdateContract)
	g.POST("/rate-card", handler.AddRateCard)
	g.PUT("/rate-card", handler.UpdateRateCard)
	g.DELETE("/rate-card", handler.DeleteRateCard)
	g.POST("/rate-card-level", handler.AddRateCardLevel)
	g.DELETE("/rate-card-level", handler.DeleteRateCardLevel)
	g.PUT("/rate-card-level", handler.UpdateRateCardLevel)

	// =================================== DOING ================================== //

	g.GET("/price", handler.GetPrice)

}
