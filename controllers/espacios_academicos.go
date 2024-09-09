package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/sga_espacio_academico_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

// EspaciosAcademicosController operations for Espacios_academicos
type EspaciosAcademicosController struct {
	beego.Controller
}

// URLMapping ...
func (c *EspaciosAcademicosController) URLMapping() {
	c.Mapping("GetAcademicSpacesByProject", c.GetAcademicSpacesByProject)
	c.Mapping("PutAcademicSpaceAssignPeriod", c.PutAcademicSpaceAssignPeriod)
	c.Mapping("GetEspacioAcademico", c.GetEspacioAcademico)
	c.Mapping("GetGruposDeEspacioAcademicoPorPeriodo", c.GetGruposDeEspacioAcademicoPorPeriodo)
	c.Mapping("DeleteGrupoEspacioAcademico", c.DeleteGrupoEspacioAcademico)
	c.Mapping("ActivarGrupoEspacioAcademico", c.ActivarGrupoEspacioAcademico)
}

// GetAcademicSpacesByProject ...
// @Title GetAcademicSpacesByProject
// @Description get Espacios_academicos for Plan Estudios
// @Param	proyecto_id		path	int	true	"Id del proyecto"
// @Success 200 {}
// @Failure 404 not found resource
// @router /proyectos/:proyecto_id [get]
func (c *EspaciosAcademicosController) GetAcademicSpacesByProject() {
	defer errorhandler.HandlePanic(&c.Controller)
	/*
		check validez de id proyecto
	*/
	idProyectoStr := c.Ctx.Input.Param(":proyecto_id")
	idProyecto, errId := strconv.ParseInt(idProyectoStr, 10, 64)
	if errId != nil || idProyecto <= 0 {
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, "Error en el formato del id del proyecto")
		c.Ctx.Output.SetStatus(400)
	} else {
		resultado := services.GetAcademicSpacesByProject(idProyecto)
		c.Data["json"] = resultado
		c.Ctx.Output.SetStatus(resultado.Status)
	}
	c.ServeJSON()
}

// PutAcademicSpaceAssignPeriod ...
// @Title PutAcademicSpaceAssignPeriod
// @Description Asigna el periodo a los grupos/espacios académicos indicados
// @Param   body        body    {}  true        "Asignar periodo a los espacios académicos"
// @Success 200 {}
// @Failure 400 the request contains incorrect syntaxis
// @router /hijos/asignar-periodo [put]
func (c *EspaciosAcademicosController) PutAcademicSpaceAssignPeriod() {
	defer errorhandler.HandlePanic(&c.Controller)
	/*
		{
			"grupo": ["Grupo 1", "Grupo 3"],
			"periodo_id": 36,
			"padre": "649cf98ecf8adba537ca9052"
		}
	*/
	dataBody := c.Ctx.Input.RequestBody
	resultado := services.PutAcademicSpaceAssignPeriod(dataBody)
	c.Data["json"] = resultado
	c.Ctx.Output.SetStatus(resultado.Status)
	c.ServeJSON()
}

// GetEspacioAcademico...
// @Title GetEspacioAcademico
// @Description obtener espacio academico con mas detalles por id
// @Param	espacio_id		path	string	true	"Id del espacio academico"
// @Success 200 {}
// @Failure 404 not found resource
// @router /:espacio_id [get]
func (c *EspaciosAcademicosController) GetEspacioAcademico() {
	defer errorhandler.HandlePanic(&c.Controller)

	espacioAcademicoId := c.Ctx.Input.Param(":espacio_id")

	resultado := services.GetEspacioAcademico(espacioAcademicoId)
	c.Data["json"] = resultado
	c.Ctx.Output.SetStatus(resultado.Status)
	c.ServeJSON()
}

// GetGruposDeEspacioAcademico...
// @Title GetGruposDeEspacioAcademico
// @Description obtener los grupos de espacio academico padre con mas detalles por periodo
// @Param	espacio-academico-id		query	string	false	"Se recibe parametro: id del espacio academico padre"
// @Param	periodo-id					query	string	false	"Se recibe parametro: id del periodo"
// @Success 200 {}
// @Failure 404 not found resource
// @router /grupos [get]
func (c *EspaciosAcademicosController) GetGruposDeEspacioAcademicoPorPeriodo() {
	defer errorhandler.HandlePanic(&c.Controller)

	espacioAcademicoId := c.GetString("espacio-academico-id")
	periodoId := c.GetString("periodo-id")

	resultado := services.GetGruposDeEspacioAcademicoPorPeriodo(espacioAcademicoId, periodoId)
	c.Data["json"] = resultado
	c.Ctx.Output.SetStatus(resultado.Status)
	c.ServeJSON()
}

// DeleteGrupoEspacioAcademico...
// @Title DeleteGrupoEspacioAcademico
// @Description eliminar grupo teniendo en cuenta que no tenga colocaciones
// @Param	grupo_id		path	string	true	"Id del espacio academico"
// @Success 200 {}
// @Failure 404 not found resource
// @router /grupo/:grupo_id [delete]
func (c *EspaciosAcademicosController) DeleteGrupoEspacioAcademico() {
	defer errorhandler.HandlePanic(&c.Controller)

	grupoId := c.Ctx.Input.Param(":grupo_id")

	resultado := services.DeleteGrupoEspacioAcademico(grupoId)
	c.Data["json"] = resultado
	c.Ctx.Output.SetStatus(resultado.Status)
	c.ServeJSON()
}

// ActivarGrupoEspacioAcademico...
// @Title ActivarGrupoEspacioAcademico
// @Description cambiar el estado de activo a true de un grupo de espacio academico
// @Param	grupo_id		path	string	true	"Id del espacio academico"
// @Success 200 {}
// @Failure 404 not found resource
// @router /activar/:grupo_id [get]
func (c *EspaciosAcademicosController) ActivarGrupoEspacioAcademico() {
	defer errorhandler.HandlePanic(&c.Controller)

	grupoId := c.Ctx.Input.Param(":grupo_id")

	resultado := services.ActivarGrupoEspacioAcademico(grupoId)
	c.Data["json"] = resultado
	c.Ctx.Output.SetStatus(resultado.Status)
	c.ServeJSON()
}
