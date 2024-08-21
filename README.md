# Paperly

**Paperly** es una plataforma diseñada para la gestión de registros de artículos académicos para la revista Tiempo UAM de la Universidad Autónoma Metropolitana, Unidad Azcapotzalco. Este sistema permite a los editores y colaboradores manejar el flujo de trabajo de los artículos, desde su envío hasta su publicación.

## Estructura del Proyecto

El proyecto está organizado en tres componentes principales:

1. **db**: Esta carpeta contiene los scripts y configuraciones necesarios para la base de datos del proyecto, la cual está implementada en PostgreSQL.
2. **api**: En esta carpeta se encuentra el código de la API, que está desarrollada en Go. La API se encarga de manejar las solicitudes del frontend y gestionar la lógica de negocio.
3. **web**: Esta carpeta contiene el frontend del proyecto, desarrollado en React. Aquí se gestiona la interacción del usuario con el sistema.

## Tecnologías Utilizadas

- **Base de Datos**: PostgreSQL
- **Backend**: Go (con Go's net/http para el servidor web y manejo de rutas)
- **Frontend**: React (utilizando React Router para la navegación y Axios para las solicitudes a la API)

## Instalación y Configuración

### Requisitos Previos

Asegúrate de tener instalados los siguientes componentes en tu sistema:

- **Go** (1.18 o superior)
- **Node.js** (14.x o superior)
- **PostgreSQL** (13.x o superior)

## Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo `LICENSE` para más detalles.

## Contacto

Para cualquier consulta o sugerencia, puedes contactar al equipo de desarrollo en [tiempouam@gmail.com](mailto:tiempouam@gmail.com).