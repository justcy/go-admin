package models

import (
	"go-admin/common/models"

)

type DesignTask struct {
    models.Model
    
    Type string `json:"type" gorm:"type:int;comment:任务类型"` 
    Name string `json:"name" gorm:"type:varchar(128);comment:名称"` 
    Tags string `json:"tags" gorm:"type:varchar(128);comment:任务标签"` 
    Comment string `json:"comment" gorm:"type:varchar(128);comment:任务备注"` 
    Params string `json:"params" gorm:"type:text;comment:任务限制参数"` 
    Status string `json:"status" gorm:"type:int;comment:任务状态 "` 
    models.ModelTime
    models.ControlBy
}

func (DesignTask) TableName() string {
    return "tb_design_task"
}

func (e *DesignTask) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *DesignTask) GetId() interface{} {
	return e.Id
}