package main

import (
	"coredemo/framework"
)

// 注册路由规则
func registerRouter(core *framework.Core) {
	// 静态路由+HTTP方法匹配
	core.Get("/user/login", UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}

func UserLoginController(c *framework.Context) error {
	err := c.Json(200, "ok, UserLoginController")
	if err != nil {
		return err
	}
	return nil
}

func SubjectListController(c *framework.Context) error {
	err := c.Json(200, "ok, SubjectListController")
	if err != nil {
		return err
	}
	return nil
}

func SubjectDelController(c *framework.Context) error {
	err := c.Json(200, "ok, SubjectDelController")
	if err != nil {
		return err
	}
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	err := c.Json(200, "ok, SubjectUpdateController")
	if err != nil {
		return err
	}
	return nil
}

func SubjectGetController(c *framework.Context) error {
	err := c.Json(200, "ok, SubjectGetController")
	if err != nil {
		return err
	}
	return nil
}

func SubjectNameController(c *framework.Context) error {
	err := c.Json(200, "ok, SubjectNameController")
	if err != nil {
		return err
	}
	return nil
}
