# REST API de Ejemplo en Go

Este es un proyecto de ejemplo que demuestra cómo crear un REST API en Go utilizando:

- **Gin**: Framework web rápido y minimalista
- **Ginkgo**: Framework de testing BDD para Go  
- **Patrón MVC**: Organización clara del código en Model-View-Controller

## 🚀 Características

- ✅ CRUD completo para usuarios
- ✅ Validación de datos
- ✅ Manejo de errores
- ✅ Tests completos con Ginkgo/Gomega
- ✅ Arquitectura MVC
- ✅ Respuestas JSON estructuradas
- ✅ CORS habilitado para desarrollo
- ✅ Health check endpoint

## 📁 Estructura del Proyecto

```
github-workflow-showcase/
├── controllers/        # Controladores (lógica de manejo de requests)
│   └── user_controller.go
├── models/            # Modelos de datos y repositorios
│   └── user.go
├── routes/            # Configuración de rutas (Vista en MVC)
│   └── routes.go
├── tests/             # Tests con Ginkgo
│   ├── tests_suite_test.go
│   └── user_test.go
├── main.go            # Punto de entrada de la aplicación
├── go.mod             # Dependencias del proyecto
└── README.md          # Este archivo
```

## 🛠️ Instalación y Ejecución

### Prerrequisitos
- Go 1.21 o superior

### Pasos para ejecutar

1. **Clonar el repositorio:**
   ```bash
   git clone <repository-url>
   cd github-workflow-showcase
   ```

2. **Instalar dependencias:**
   ```bash
   go mod tidy
   ```

3. **Ejecutar la aplicación:**
   ```bash
   go run main.go
   ```

4. **La API estará disponible en:**
   - Servidor: `http://localhost:8080`
   - Health check: `http://localhost:8080/health`
   - Documentación: `http://localhost:8080/`

## 🧪 Ejecutar Tests

```bash
# Ejecutar todos los tests con Ginkgo
go test ./tests/

# O usando ginkgo directamente (si está instalado)
ginkgo ./tests/

# Para instalar ginkgo CLI globalmente
go install github.com/onsi/ginkgo/v2/ginkgo
```

## 📚 API Endpoints

### Health Check
- **GET** `/health` - Verificar estado de la API

### Root
- **GET** `/` - Información de la API y endpoints disponibles

### Usuarios
- **GET** `/api/v1/users` - Obtener todos los usuarios
- **GET** `/api/v1/users/:id` - Obtener usuario por ID
- **POST** `/api/v1/users` - Crear nuevo usuario
- **PUT** `/api/v1/users/:id` - Actualizar usuario existente
- **DELETE** `/api/v1/users/:id` - Eliminar usuario

## 💡 Ejemplos de Uso

### Crear un usuario
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Juan Pérez",
    "email": "juan@example.com",
    "age": 30
  }'
```

### Obtener todos los usuarios
```bash
curl http://localhost:8080/api/v1/users
```

### Obtener usuario por ID
```bash
curl http://localhost:8080/api/v1/users/1
```

### Actualizar usuario
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Juan Pérez Actualizado",
    "email": "juan.updated@example.com",
    "age": 31
  }'
```

### Eliminar usuario
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## 🏗️ Arquitectura MVC

### Model (Modelo)
- **Ubicación**: `models/user.go`
- **Responsabilidad**: Definir la estructura de datos, validaciones y lógica de persistencia
- **Características**:
  - Struct `User` con validaciones
  - Interface `UserRepository` para abstracción de datos
  - Implementación `InMemoryUserRepository` para almacenamiento en memoria

### View (Vista)
- **Ubicación**: `routes/routes.go`
- **Responsabilidad**: Configurar rutas y middlewares
- **Características**:
  - Configuración de router con Gin
  - Middlewares de CORS, logging y recovery
  - Agrupación de rutas por versión de API

### Controller (Controlador)
- **Ubicación**: `controllers/user_controller.go`
- **Responsabilidad**: Manejar requests HTTP y lógica de negocio
- **Características**:
  - Métodos para cada operación CRUD
  - Validación de datos de entrada
  - Manejo de errores y respuestas JSON estructuradas

## 🔧 Configuración

### Variables de Entorno
- `PORT`: Puerto del servidor (por defecto: 8080)

### Personalización
- Para cambiar a base de datos real, implementar nueva versión de `UserRepository`
- Para añadir autenticación, agregar middleware en `routes/routes.go`
- Para cambiar modo de producción, descomentar `gin.SetMode(gin.ReleaseMode)` en routes

## 🧪 Testing con Ginkgo

Los tests están escritos usando Ginkgo (BDD) y Gomega (assertions). Incluyen:

- Tests de health check
- Tests de endpoints de información
- Tests completos de CRUD para usuarios
- Tests de validación de datos
- Tests de manejo de errores

### Estructura de Tests
```go
Describe("User API", func() {
  Describe("GET /users", func() {
    It("should return users list", func() {
      // Test implementation
    })
  })
})
```

## 🤝 Contribuir

1. Fork el proyecto
2. Crear una rama para tu feature (`git checkout -b feature/nueva-caracteristica`)
3. Commit tus cambios (`git commit -am 'Agregar nueva característica'`)
4. Push a la rama (`git push origin feature/nueva-caracteristica`)
5. Abrir un Pull Request

## 📝 Licencia

Este proyecto es de ejemplo y está disponible bajo la licencia MIT.

## 📝 Semantic PR Titles

Este proyecto requiere que todos los Pull Requests sigan el formato de [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>[optional scope]: <description>
```

### Tipos Permitidos:
- `feat:` - Nueva funcionalidad
- `fix:` - Corrección de errores
- `docs:` - Cambios en documentación
- `style:` - Formato de código
- `refactor:` - Refactorización
- `perf:` - Optimización de rendimiento
- `test:` - Cambios en pruebas
- `build:` - Sistema de build
- `ci:` - CI/CD
- `chore:` - Mantenimiento

### Ejemplos:
```
feat: add user authentication endpoint
fix: resolve database connection timeout
docs: update API documentation
refactor: simplify error handling
```

### 🤖 Validación Automática
- ✅ Validación automática del formato de título
- 💡 Sugerencias inteligentes basadas en archivos modificados
- 🏷️ Labels automáticos según el tipo
- 💬 Comentarios de ayuda para títulos inválidos

📚 **[Guía Completa de Títulos Semánticos](./.github/SEMANTIC_PR_GUIDE.md)**

## 🔗 Recursos Útiles

- [Gin Framework](https://gin-gonic.com/)
- [Ginkgo Testing Framework](https://onsi.github.io/ginkgo/)
- [Gomega Assertion Library](https://onsi.github.io/gomega/)
- [Go Documentation](https://golang.org/doc/)
- [Conventional Commits](https://www.conventionalcommits.org/)
