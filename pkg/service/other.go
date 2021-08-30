package service

import "github.com/labstack/echo"

func (s *Service) Accesible(c echo.Context) error {
	return c.NoContent(200)
}
