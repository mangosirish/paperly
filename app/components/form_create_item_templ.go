// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func CreateItemForm(requested_table string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"form-create\" x-show=\"open_form_create\" x-transition @keydown.escape.window=\"confirmClose()\" @click.away=\"confirmClose()\"><div class=\"form-content\" @click.stop><span class=\"close\" x-on:click=\"confirmClose()\">&times;</span><h2>Create ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(requested_table)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/form_create_item.templ`, Line: 7, Col: 31}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><div class=\"form-scrollable\"><!-- Aquí se selecciona el formulario correspondiente según la tabla solicitada -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		switch requested_table {
		case "articles":
			templ_7745c5c3_Err = FormArticles().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case "journals":
			templ_7745c5c3_Err = FormJournals().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case "authors":
			templ_7745c5c3_Err = FormAuthors().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case "social-service":
			templ_7745c5c3_Err = FormSocialService().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		default:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p>Formulario no disponible.</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div><script>\n    function createFormHandler() {\n        return {\n            open_form_create: false,\n            confirmClose() {\n                if (confirm(\"¿Realmente desea salir sin guardar?\")) {\n                    this.open_form_create = false;\n                }\n            }\n        };\n    }\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func FormArticles() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form action=\"/create_article\" method=\"POST\"><label for=\"title\">Título del artículo:</label> <input type=\"text\" id=\"title\" name=\"title\" placeholder=\"Título del artículo\" required> <label for=\"type\">Tipo de artículo:</label> <input type=\"text\" id=\"type\" name=\"type\" placeholder=\"Tipo de artículo\" required> <label for=\"reception_date\">Fecha de recepción:</label> <input type=\"date\" id=\"reception_date\" name=\"reception_date\" required> <label for=\"status\">Estado:</label> <input type=\"text\" id=\"status\" name=\"status\" placeholder=\"Estado del artículo\" required> <button type=\"submit\">Crear Artículo</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func FormAuthors() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form action=\"/create_author\" method=\"POST\"><label for=\"first_name\">Nombre(s):</label> <input type=\"text\" id=\"first_name\" name=\"first_name\" placeholder=\"Nombre del autor\" required> <label for=\"first_surname\">Primer apellido:</label> <input type=\"text\" id=\"first_surname\" name=\"first_surname\" placeholder=\"Primer apellido\" required> <label for=\"second_surname\">Segundo apellido:</label> <input type=\"text\" id=\"second_surname\" name=\"second_surname\" placeholder=\"Segundo apellido\" required> <label for=\"personal_email\">Correo personal:</label> <input type=\"email\" id=\"personal_email\" name=\"personal_email\" placeholder=\"Correo personal del autor\" required> <label for=\"institutional_email\">Correo institucional:</label> <input type=\"email\" id=\"institutional_email\" name=\"institutional_email\" placeholder=\"Correo institucional del autor\"> <label for=\"notes\">Notas sobre el autor:</label> <textarea id=\"notes\" name=\"notes\" placeholder=\"Notas adicionales\"></textarea> <button type=\"submit\">Crear Autor</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func FormJournals() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form action=\"/create_journal\" method=\"POST\"><label for=\"status\">Estado:</label> <input type=\"text\" id=\"status\" name=\"status\" placeholder=\"Estado del número\" required> <label for=\"age\">Antigüedad (años):</label> <input type=\"number\" id=\"age\" name=\"age\" placeholder=\"Años de antigüedad\" required> <label for=\"publication_date\">Fecha de publicación:</label> <input type=\"date\" id=\"publication_date\" name=\"publication_date\" required> <label for=\"start_month_period\">Mes de inicio del periodo:</label> <input type=\"text\" id=\"start_month_period\" name=\"start_month_period\" placeholder=\"Mes de inicio\" required> <label for=\"end_month_period\">Mes de término del periodo:</label> <input type=\"text\" id=\"end_month_period\" name=\"end_month_period\" placeholder=\"Mes de término\" required> <label for=\"number\">Número del ejemplar:</label> <input type=\"number\" id=\"number\" name=\"number\" placeholder=\"Número del ejemplar\" required> <label for=\"volume_number\">Número de volumen:</label> <input type=\"number\" id=\"volume_number\" name=\"volume_number\" placeholder=\"Número de volumen\" required> <label for=\"edition_number\">Número de edición:</label> <input type=\"number\" id=\"edition_number\" name=\"edition_number\" required> <label for=\"special_number\">Número especial (si aplica):</label> <input type=\"number\" id=\"special_number\" name=\"special_number\"> <label for=\"online_link\">Enlace en línea:</label> <input type=\"url\" id=\"online_link\" name=\"online_link\" placeholder=\"URL del ejemplar\"> <label for=\"reserve_number\">Número de reserva:</label> <input type=\"text\" id=\"reserve_number\" name=\"reserve_number\" placeholder=\"Número de reserva\"> <button type=\"submit\">Crear Número</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func FormSocialService() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form action=\"/create_student_service\" method=\"POST\"><label for=\"person_id\">ID de persona:</label> <input type=\"number\" id=\"person_id\" name=\"person_id\" placeholder=\"ID de la persona (estudiante)\" required> <label for=\"start_date\">Fecha de inicio:</label> <input type=\"date\" id=\"start_date\" name=\"start_date\" required> <label for=\"end_date\">Fecha de término:</label> <input type=\"date\" id=\"end_date\" name=\"end_date\" required> <label for=\"documentation\">Documentación:</label> <textarea id=\"documentation\" name=\"documentation\" placeholder=\"Documentación adicional\"></textarea> <label for=\"status\">Estado del servicio social:</label> <input type=\"text\" id=\"status\" name=\"status\" placeholder=\"Estado\" required> <label for=\"division\">División:</label> <input type=\"text\" id=\"division\" name=\"division\" placeholder=\"División del servicio social\" required> <label for=\"institution\">Institución:</label> <input type=\"text\" id=\"institution\" name=\"institution\" placeholder=\"Institución del servicio social\" required><fieldset><legend>Documentación presentada:</legend> <label><input type=\"checkbox\" name=\"is_enrollment_request_submitted\"> Solicitud de inscripción</label> <label><input type=\"checkbox\" name=\"is_presentation_letter_submitted\"> Carta de presentación</label> <label><input type=\"checkbox\" name=\"is_acceptance_letter_submitted\"> Carta de aceptación</label> <label><input type=\"checkbox\" name=\"is_advisor_appointment_submitted\"> Nombramiento de asesor</label> <label><input type=\"checkbox\" name=\"is_commitment_letter_submitted\"> Carta de compromiso</label> <label><input type=\"checkbox\" name=\"is_intermediate_report_submitted\"> Informe intermedio</label> <label><input type=\"checkbox\" name=\"is_intermediate_report_validation_submitted\"> Validación del informe intermedio</label> <label><input type=\"checkbox\" name=\"is_final_report_submitted\"> Informe final</label> <label><input type=\"checkbox\" name=\"is_completion_letter_submitted\"> Carta de terminación</label></fieldset><button type=\"submit\">Crear Servicio Social</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
