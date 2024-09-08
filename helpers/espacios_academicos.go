package helpers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func GetDocente(docenteId string) (map[string]interface{}, error) {
	urlDocente := "http://" + beego.AppConfig.String("TercerosService") + "tercero/" + docenteId
	fmt.Println(urlDocente)
	var docente map[string]interface{}
	if err := request.GetJson(urlDocente, &docente); err != nil {
		return nil, fmt.Errorf("error en el servicio de terceros")
	}
	return docente, nil
}

func GetPoyectoAcademico(proyectoId string) (map[string]interface{}, error) {
	urlProyecto := "http://" + beego.AppConfig.String("ProyectoAcademicoService") + "proyecto_academico_institucion/" + proyectoId
	var proyecto map[string]interface{}
	if err := request.GetJson(urlProyecto, &proyecto); err != nil {
		return nil, fmt.Errorf("error en el servicio de proyecto académico")
	}
	return proyecto, nil
}

func DesactivarGrupoEspacioAcademico(grupoId string) (map[string]interface{}, error) {
	urlGrupo := "http://" + beego.AppConfig.String("EspaciosAcademicosService") + "espacio-academico?query=_id:" + grupoId
	var grupo map[string]interface{}
	if err := request.GetJson(urlGrupo, &grupo); err != nil {
		return nil, fmt.Errorf("error en el servicio de espacios academicos: %v", err)
	}

	grupoData := grupo["Data"].([]interface{})[0].(map[string]interface{})
	grupoData["activo"] = false

	urlGrupoPut := "http://" + beego.AppConfig.String("EspaciosAcademicosService") + "espacio-academico/" + grupoId
	var grupoPut map[string]interface{}
	if err := request.SendJson(urlGrupoPut, "PUT", &grupoPut, grupoData); err != nil {
		return nil, fmt.Errorf("error en el servicio de horario: %v", err)
	}

	return grupoPut["Data"].(map[string]interface{}), nil
}
