package dto

import ("time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type DesignTaskGetPageReq struct {
	dto.Pagination     `search:"-"`
    DesignTaskOrder
}

type DesignTaskOrder struct {Id int `form:"idOrder"  search:"type:order;column:id;table:tb_design_task"`
    Type string `form:"typeOrder"  search:"type:order;column:type;table:tb_design_task"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:tb_design_task"`
    Tags string `form:"tagsOrder"  search:"type:order;column:tags;table:tb_design_task"`
    Comment string `form:"commentOrder"  search:"type:order;column:comment;table:tb_design_task"`
    Params string `form:"paramsOrder"  search:"type:order;column:params;table:tb_design_task"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:tb_design_task"`
    CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_design_task"`
    UpdatedAt time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_design_task"`
    DeletedAt time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_design_task"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_design_task"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_design_task"`
    
}

func (m *DesignTaskGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type DesignTaskInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    Type string `json:"type" comment:"任务类型"`
    Name string `json:"name" comment:"名称"`
    Tags string `json:"tags" comment:"任务标签"`
    Comment string `json:"comment" comment:"任务备注"`
    Params string `json:"params" comment:"任务限制参数"`
    Status string `json:"status" comment:"任务状态 0 等待处理 1 处理完成待交付 2 已交付 "`
    common.ControlBy
}

func (s *DesignTaskInsertReq) Generate(model *models.DesignTask)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Type = s.Type
    model.Name = s.Name
    model.Tags = s.Tags
    model.Comment = s.Comment
    model.Params = s.Params
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *DesignTaskInsertReq) GetId() interface{} {
	return s.Id
}

type DesignTaskUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    Type string `json:"type" comment:"任务类型"`
    Name string `json:"name" comment:"名称"`
    Tags string `json:"tags" comment:"任务标签"`
    Comment string `json:"comment" comment:"任务备注"`
    Params string `json:"params" comment:"任务限制参数"`
    Status string `json:"status" comment:"任务状态 0 等待处理 1 处理完成待交付 2 已交付 "`
    common.ControlBy
}

func (s *DesignTaskUpdateReq) Generate(model *models.DesignTask)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Type = s.Type
    model.Name = s.Name
    model.Tags = s.Tags
    model.Comment = s.Comment
    model.Params = s.Params
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *DesignTaskUpdateReq) GetId() interface{} {
	return s.Id
}

// DesignTaskGetReq 功能获取请求参数
type DesignTaskGetReq struct {
     Id int `uri:"id"`
}
func (s *DesignTaskGetReq) GetId() interface{} {
	return s.Id
}

// DesignTaskDeleteReq 功能删除请求参数
type DesignTaskDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *DesignTaskDeleteReq) GetId() interface{} {
	return s.Ids
}