package v1

import (
	"gin-blog/common/config"
	"gin-blog/common/response"
	"gin-blog/models"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取多个文章标签
func GetTags(c *gin.Context) {

	// 请求参数
	name := c.Query("name")
	maps := make(map[string]interface{}, 2)
	if name != "" {
		maps["name"] = name
	}
	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	// 查询数据
	data := map[string]interface{}{
		"lists": models.GetTags(0, 10, maps),
		"total": models.GetTagTotal(maps),
	}

	// 返回数据
	response.ResponseSuccess(c, data)
}

//新增文章标签
func AddTag(c *gin.Context) {

	// 请求参数
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	// 校验参数
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	// 参数有误
	if valid.HasErrors() {
		response.ResponseError(c, config.INVALID_PARAMS)
		return
	}
	// 标签已存在
	tag, err := models.GetTagInfo(name, 0)
	if tag.ID > 0 || err != nil {
		response.ResponseError(c, config.ERROR_EXIST_TAG)
		return
	}
	// 添加标签
	res := models.AddTag(name, state, createdBy)
	if !res {
		response.ResponseError(c, config.ERROR)
		return
	}

	// 返回数据
	response.ResponseSuccess(c, nil)
}

//修改文章标签
func EditTag(c *gin.Context) {

	// 请求参数
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	// 检验数据
	valid := validation.Validation{}
	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	// 参数有误
	if valid.HasErrors() {
		response.ResponseError(c, config.INVALID_PARAMS)
		return
	}

	// 查询标签
	tag, err := models.GetTagInfo(name, 0)
	if tag == nil || err != nil {
		response.ResponseError(c, config.ERROR_NOT_EXIST_TAG)
		return
	}

	// 更新数据
	data := make(map[string]interface{}, 3)
	data["modified_by"] = modifiedBy
	if name != "" {
		data["name"] = name
	}
	if state != -1 {
		data["state"] = state
	}
	err = models.EditTag(id, data)
	if err != nil {
		response.ResponseError(c, config.ERROR)
		return
	}

	// 返回数据
	response.ResponseSuccess(c, nil)

}

//删除文章标签
func DeleteTag(c *gin.Context) {

	// 请求参数
	id := com.StrTo(c.Param("id")).MustInt()

	// 校验参数
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	// 参数有误
	if valid.HasErrors() {
		response.ResponseError(c, config.INVALID_PARAMS)
		return
	}

	// 查询标签
	tag, err := models.GetTagInfo("", id)
	if tag == nil || err != nil {
		response.ResponseError(c, config.ERROR_NOT_EXIST_TAG)
		return
	}

	// 删除标签
	err = models.DeleteTag(id)
	if err != nil {
		response.ResponseError(c, config.ERROR)
		return
	}

	// 返回数据
	response.ResponseSuccess(c, nil)
}
