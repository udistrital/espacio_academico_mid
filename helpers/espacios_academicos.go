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
