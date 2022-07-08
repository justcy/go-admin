package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type DesignTask struct {
	api.Api
}

// GetPage 获取DesignTask列表
// @Summary 获取DesignTask列表
// @Description 获取DesignTask列表
// @Tags DesignTask
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.DesignTask}} "{"code": 200, "data": [...]}"
// @Router /api/v1/design-task [get]
// @Security Bearer
func (e DesignTask) GetPage(c *gin.Context) {
    req := dto.DesignTaskGetPageReq{}
    s := service.DesignTask{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.DesignTask, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取DesignTask 失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取DesignTask
// @Summary 获取DesignTask
// @Description 获取DesignTask
// @Tags DesignTask
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.DesignTask} "{"code": 200, "data": [...]}"
// @Router /api/v1/design-task/{id} [get]
// @Security Bearer
func (e DesignTask) Get(c *gin.Context) {
	req := dto.DesignTaskGetReq{}
	s := service.DesignTask{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.DesignTask

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取DesignTask失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建DesignTask
// @Summary 创建DesignTask
// @Description 创建DesignTask
// @Tags DesignTask
// @Accept application/json
// @Product application/json
// @Param data body dto.DesignTaskInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/design-task [post]
// @Security Bearer
func (e DesignTask) Insert(c *gin.Context) {
    req := dto.DesignTaskInsertReq{}
    s := service.DesignTask{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建DesignTask  失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改DesignTask
// @Summary 修改DesignTask
// @Description 修改DesignTask
// @Tags DesignTask
// @Accept application/json
// @Product application/json
// @Param data body dto.DesignTaskUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/design-task/{id} [put]
// @Security Bearer
func (e DesignTask) Update(c *gin.Context) {
    req := dto.DesignTaskUpdateReq{}
    s := service.DesignTask{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改DesignTask 失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除DesignTask
// @Summary 删除DesignTask
// @Description 删除DesignTask
// @Tags DesignTask
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/design-task [delete]
// @Security Bearer
func (e DesignTask) Delete(c *gin.Context) {
    s := service.DesignTask{}
    req := dto.DesignTaskDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除DesignTask失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}