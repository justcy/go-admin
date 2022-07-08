package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type DesignTask struct {
	service.Service
}

// GetPage 获取DesignTask列表
func (e *DesignTask) GetPage(c *dto.DesignTaskGetPageReq, p *actions.DataPermission, list *[]models.DesignTask, count *int64) error {
	var err error
	var data models.DesignTask

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("DesignTaskService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取DesignTask对象
func (e *DesignTask) Get(d *dto.DesignTaskGetReq, p *actions.DataPermission, model *models.DesignTask) error {
	var data models.DesignTask

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetDesignTask error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建DesignTask对象
func (e *DesignTask) Insert(c *dto.DesignTaskInsertReq) error {
    var err error
    var data models.DesignTask
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("DesignTaskService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改DesignTask对象
func (e *DesignTask) Update(c *dto.DesignTaskUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.DesignTask{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if db.Error != nil {
        e.Log.Errorf("DesignTaskService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除DesignTask
func (e *DesignTask) Remove(d *dto.DesignTaskDeleteReq, p *actions.DataPermission) error {
	var data models.DesignTask

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveDesignTask error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}