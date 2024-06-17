package book

import (
	"github.com/labstack/echo/v4"
)

func (r *Book) GetBookListHandler(c echo.Context) error {
	result := r.bookSrv.GetList()

	// Mapping to response mapper
	data := []BookResponsePayloadData{}
	for _, b := range result {
		data = append(data, BookResponsePayloadData{
			Code:  b.Code,
			Title: b.Title,
			Desc:  b.Description,
			Price: b.Price,
		})
	}

	return c.JSON(200, BookResponsePayload{
		Status:  200,
		Message: "SUCCESS",
		Data:    data,
	})
}
