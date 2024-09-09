package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/sga_espacio_academico_mid/controllers:EspaciosAcademicosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_espacio_academico_mid/controllers:EspaciosAcademicosController"],
        beego.ControllerComments{
            Method: "GetEspacioAcademico",
            Router: "/:espacio_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_espacio_academico_mid/controllers:EspaciosAcademicosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_espacio_academico_mid/controllers:EspaciosAcademicosController"],
        beego.ControllerComments{
            Method: "ActivarGrupoEspacioAcademico",
            Router: "/activar/:grupo_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_espacio_academico_mid/controllers:EspaciosAcademicosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_espacio_academico_mid/controllers:EspaciosAcademicosController"],
        beego.ControllerComments{
            Method: "DeleteGrupoEspacioAcademico",
            Router: "/grupo/:grupo_id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_espacio_academico_mid/controllers:EspaciosAcademicosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_espacio_academico_mid/controllers:EspaciosAcademicosController"],
        beego.ControllerComments{
            Method: "GetGruposDeEspacioAcademicoPorPeriodo",
            Router: "/grupos",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_espacio_academico_mid/controllers:EspaciosAcademicosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_espacio_academico_mid/controllers:EspaciosAcademicosController"],
        beego.ControllerComments{
            Method: "PutAcademicSpaceAssignPeriod",
            Router: "/hijos/asignar-periodo",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_espacio_academico_mid/controllers:EspaciosAcademicosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_espacio_academico_mid/controllers:EspaciosAcademicosController"],
        beego.ControllerComments{
            Method: "GetAcademicSpacesByProject",
            Router: "/proyectos/:proyecto_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
