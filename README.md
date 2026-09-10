## GESTION DE EMPLEADOS

Una aplicación web simple de CRUD (Create, Read, Update, Delete) para gestionar empleados, desarrollada en Go con base de datos MySQL.

## Tecnologías Utilizadas

- **Go 1.21** - Lenguaje de programación
- **MySQL 8.0** - Base de datos
- **Docker & Docker Compose** - Contenedorización
- **HTML Templates** - Plantillas para la interfaz web
- **Bootstrap 4.3** - Framework CSS para el diseño

## Características

- ✅ Crear empleados
- ✅ Leer/listar empleados
- ✅ Actualizar empleados
- ✅ Eliminar empleados
- ✅ Interfaz web responsive
- ✅ Contenedorización con Docker

## Requisitos Previos

- Docker Desktop instalado y ejecutándose
- Docker Compose

## Instalación y Configuración

### Opción 1: Usando Docker (Recomendado)

1. Clona el repositorio:
   ```bash
   git clone <url-del-repositorio>
   cd crud-golang
   ```

2. Construye y ejecuta los contenedores:
   ```bash
   docker-compose up --build
   ```

3. Accede a la aplicación en: http://localhost:8080

### Opción 2: Ejecución Local

1. Instala Go 1.21+ y MySQL 8.0

2. Configura la base de datos MySQL:
   - Crea una base de datos llamada `golang_crud`
   - Ejecuta el script `init.sql` para crear la tabla

3. Instala las dependencias:
   ```bash
   go mod tidy
   ```

4. Ejecuta la aplicación:
   ```bash
   go run main.go
   ```

5. Accede en: http://localhost:8080

## Estructura del Proyecto

```
crud-golang/
├── main.go              # Código principal de la aplicación
├── go.mod               # Módulos de Go
├── go.sum               # Suma de verificación de dependencias
├── Dockerfile           # Configuración para construir la imagen de la app
├── docker-compose.yml   # Orquestación de contenedores
├── init.sql            # Script de inicialización de la base de datos
├── .dockerignore       # Archivos a ignorar en Docker
└── templates/          # Plantillas HTML
    ├── header-tmpl.html
    ├── footer-tmpl.html
    ├── start-tmpl.html
    ├── create-tmpl.html
    └── edit-tmpl.html
```

## Endpoints de la API

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/` | Lista todos los empleados |
| GET | `/create` | Muestra formulario para crear empleado |
| POST | `/insert` | Crea un nuevo empleado |
| GET | `/edit?id={id}` | Muestra formulario para editar empleado |
| POST | `/update` | Actualiza un empleado existente |
| GET | `/delete?id={id}` | Elimina un empleado |

## Base de Datos

### Tabla: employees

```sql
CREATE TABLE employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);
```

### Datos de Ejemplo

La aplicación incluye datos de ejemplo que se insertan automáticamente al iniciar con Docker.

## Variables de Entorno

| Variable | Valor por Defecto | Descripción |
|----------|-------------------|-------------|
| DB_HOST | localhost | Host de la base de datos |
| DB_USER | root | Usuario de MySQL |
| DB_PASSWORD | root | Contraseña de MySQL |
| DB_NAME | golang_crud | Nombre de la base de datos |

## Desarrollo

### Comandos Útiles

```bash
# Ejecutar en modo desarrollo
go run main.go

# Construir la aplicación
go build -o main .

# Ejecutar tests (si los hay)
go test ./...

# Limpiar contenedores Docker
docker-compose down
docker system prune -f
```

### Agregar Nuevas Funcionalidades

1. Modifica `main.go` para agregar nuevos endpoints
2. Crea nuevas plantillas HTML en la carpeta `templates/`
3. Actualiza la base de datos modificando `init.sql`

## Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

## Autor

**Tu Nombre** - [Tu GitHub](https://github.com/tu-usuario)

## Agradecimientos

- [Go SQL Driver](https://github.com/go-sql-driver/mysql) - Driver MySQL para Go
- [Bootstrap](https://getbootstrap.com/) - Framework CSS
- [Docker](https://www.docker.com/) - Plataforma de contenedorización
