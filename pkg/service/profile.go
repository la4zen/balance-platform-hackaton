package service

import (
	"io"
	"os"
	"strconv"

	"github.com/la4zen/balance-platform-hackaton/pkg/models"
	"github.com/labstack/echo"
)

func (s *Service) UpdateProfile(c echo.Context) error {
	user := c.Get("user").(*models.User)
	c.Bind(user)
	file, err := c.FormFile("avatar")
	if err != nil {
		return err
	}
	if file != nil {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		dst, err := os.Create("static/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		user.AvatarURI = "/img/" + file.Filename
	}
	if err := s.Store.UpdateUser(user); err != nil {
		_ = os.Remove("static/" + file.Filename)
		return err
	}
	return c.NoContent(200)
}

func (s *Service) GetProfile(c echo.Context) error {
	_id := c.Param("id")
	if _id == "" {
		return c.NoContent(400)
	}
	id, err := strconv.Atoi(_id)
	if err != nil {
		return c.NoContent(400)
	}
	user := &models.User{
		ID: uint(id),
	}
	if s.Store.GetUser(user) != nil {
		return c.NoContent(500)
	}
	return c.JSON(200, user)
}
