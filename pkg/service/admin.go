package service

import (
	"github.com/la4zen/balance-platform-hackaton/pkg/models"
	"github.com/labstack/echo"
)

func (s *Service) AddTask(c echo.Context) error {
	user := c.Get("user").(*models.User)
	if !user.IsAdmin {
		return echo.ErrForbidden.Internal
	}
	task := &models.Task{}
	c.Bind(task)
	if err := task.Validate(); err != nil {
		return c.String(400, err.Error())
	}
	if err := s.Store.AddTask(task); err != nil {
		return err
	}
	return c.NoContent(200)
}

func (s *Service) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*models.User)
	if !user.IsAdmin {
		return echo.ErrForbidden.Internal
	}
	task := &models.Task{}
	c.Bind(task)
	if task.ID == 0 {
		return c.String(400, "task id required")
	}
	if err := s.Store.DeleteTask(task); err != nil {
		return err
	}
	return c.NoContent(200)
}

func (s *Service) GetTasks(c echo.Context) error {
	user := c.Get("user").(*models.User)
	if !user.IsAdmin {
		return echo.ErrForbidden.Internal
	}
	tasks, err := s.Store.GetTasks()
	if err != nil {
		return err
	}
	return c.JSON(200, tasks)
}
