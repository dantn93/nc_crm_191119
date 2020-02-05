package handler

// func UpdateUser(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*model.UserClaims)

// 	var req model.UserUpdateReq
// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
// 	}

// 	req.ID = claims.UserID

// 	res, err := db.UpdateUser(&req)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
// 	}

// 	return c.JSON(http.StatusOK, res)
// }
