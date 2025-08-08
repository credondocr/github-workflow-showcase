# 📝 Semantic PR Title Guide

Este proyecto requiere que todos los títulos de Pull Requests sigan el formato de [Conventional Commits](https://www.conventionalcommits.org/). Esto nos ayuda a mantener un historial claro, generar changelogs automáticos y categorizar cambios eficientemente.

## 📋 Formato Requerido

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Ejemplos Válidos:
- `feat: add user authentication`
- `fix: resolve login timeout issue`
- `docs: update API documentation`
- `feat(auth): implement OAuth2 support`
- `fix(api): handle null responses correctly`

## 🎯 Tipos Permitidos

| Tipo | Descripción | Cuándo Usar | Ejemplo |
|------|-------------|-------------|---------|
| `feat` | Nueva funcionalidad | Añadir características nuevas | `feat: add email notifications` |
| `fix` | Corrección de errores | Solucionar bugs o problemas | `fix: resolve memory leak in parser` |
| `docs` | Documentación | Cambios solo en documentación | `docs: update installation guide` |
| `style` | Formato de código | Cambios de estilo sin afectar lógica | `style: format code with gofmt` |
| `refactor` | Refactorización | Reestructurar código sin cambiar funcionalidad | `refactor: simplify validation logic` |
| `perf` | Optimización | Mejoras de rendimiento | `perf: optimize database queries` |
| `test` | Pruebas | Añadir o corregir tests | `test: add unit tests for auth module` |
| `build` | Sistema de build | Cambios en dependencias o build | `build: update Go to version 1.21` |
| `ci` | Integración continua | Cambios en CI/CD | `ci: add automated deployment` |
| `chore` | Tareas de mantenimiento | Cambios que no afectan src o tests | `chore: update .gitignore` |
| `revert` | Revertir cambios | Deshacer commits anteriores | `revert: "feat: user authentication"` |

## 🔧 Scope (Opcional)

El scope especifica qué parte del código fue afectada:

```
feat(api): add new endpoint
fix(auth): resolve token validation
docs(readme): update examples
```

### Scopes Comunes:
- `api` - Cambios en la API
- `auth` - Autenticación y autorización
- `ui` - Interfaz de usuario
- `db` - Base de datos
- `test` - Archivos de prueba
- `ci` - Configuración de CI/CD

## 📏 Reglas de Formato

### ✅ Hacer:
- Usar minúsculas para el tipo: `feat:` no `Feat:`
- Usar tiempo presente: `add` no `added`
- Ser específico y claro: `fix login validation` no `fix bug`
- Mantener el título bajo 100 caracteres
- No terminar con punto final

### ❌ No hacer:
- `Update code` → Muy vago
- `Fix bug` → No específico
- `feat: Add new feature.` → No usar punto final
- `FEAT: add feature` → No usar mayúsculas
- `added user authentication` → No usar past tense

## 🎯 Ejemplos por Categoría

### 🆕 Nuevas Características (feat)
```
feat: add user registration endpoint
feat: implement dark mode toggle
feat(api): add pagination to user list
feat(auth): integrate OAuth2 providers
```

### 🐛 Corrección de Errores (fix)
```
fix: resolve memory leak in image processor
fix: handle null pointer in user validation
fix(api): correct response status codes
fix(auth): prevent token expiration edge case
```

### 📚 Documentación (docs)
```
docs: add API endpoint documentation
docs: update README with new examples
docs(api): document authentication flow
docs: fix typos in installation guide
```

### 🎨 Estilo y Formato (style)
```
style: format all Go files with gofmt
style: fix indentation in config files
style: remove trailing whitespace
style(css): organize selectors alphabetically
```

### ♻️ Refactorización (refactor)
```
refactor: extract validation logic to separate module
refactor: simplify error handling in controllers
refactor(models): improve user struct design
refactor: reduce cyclomatic complexity in parser
```

### ⚡ Rendimiento (perf)
```
perf: optimize database connection pooling
perf: reduce memory allocation in hot path
perf(api): implement response caching
perf: improve algorithm efficiency by 40%
```

### 🧪 Pruebas (test)
```
test: add integration tests for auth endpoints
test: increase coverage for user validation
test(unit): cover edge cases in parser
test: add performance benchmarks
```

### 🏗️ Build y Dependencias (build)
```
build: update Go version to 1.21
build: add Docker multi-stage build
build(deps): bump gin framework to v1.9.1
build: configure automated security scanning
```

### 🔄 CI/CD (ci)
```
ci: add automated testing pipeline
ci: configure deployment to staging
ci(github): update workflow permissions
ci: add code coverage reporting
```

### 🧹 Mantenimiento (chore)
```
chore: update .gitignore patterns
chore: clean up unused dependencies
chore(release): prepare version 2.1.0
chore: organize project structure
```

## 🚨 Breaking Changes

Para cambios que rompen compatibilidad, usa `!` después del tipo:

```
feat!: change authentication API structure
fix!: remove deprecated user endpoints
refactor!: restructure configuration format
```

## 🤖 Validación Automática

Este proyecto incluye validación automática que:

1. ✅ **Verifica el formato** - Confirma que sigue conventional commits
2. 🏷️ **Añade labels** - Etiqueta automáticamente basado en el tipo
3. 💬 **Proporciona feedback** - Sugiere correcciones si es inválido
4. 🔄 **Actualiza comentarios** - Mantiene ayuda actualizada

### Proceso de Validación:
1. Abres/editas un PR
2. El sistema valida el título automáticamente
3. Si es válido: ✅ Se añade label apropiado
4. Si es inválido: ❌ Se comenta con sugerencias
5. Corriges el título según las sugerencias
6. ✅ Validación pasa automáticamente

## 💡 Consejos para Títulos Efectivos

### 🎯 Sé Específico
```
❌ fix: bug in login
✅ fix: resolve timeout in OAuth callback
```

### 📝 Describe el "Qué", no el "Cómo"
```
❌ feat: add if statement for validation
✅ feat: validate user input before processing
```

### 🔍 Usa el Scope para Claridad
```
❌ fix: validation issue
✅ fix(auth): email validation accepts invalid domains
```

### 📊 Incluye Contexto Relevante
```
❌ perf: optimization
✅ perf: reduce API response time by 60%
```

## 🔗 Integración con Releases

Los títulos semánticos permiten:

- **Generación automática de changelogs**
- **Versionado semántico automático**
- **Categorización de cambios en releases**
- **Detección de breaking changes**

### Ejemplo de Changelog Generado:
```markdown
## [2.1.0] - 2024-03-15

### Features
- feat: add user authentication
- feat(api): implement rate limiting

### Bug Fixes  
- fix: resolve memory leak in parser
- fix(auth): handle token expiration

### Documentation
- docs: update API guide with examples
```

## 🚀 Herramientas Útiles

### Extensiones de VS Code:
- **Conventional Commits** - Ayuda con formato
- **GitLens** - Historial de commits mejorado

### CLI Tools:
```bash
# Commitizen - ayuda interactiva para commits
npm install -g commitizen
npm install -g cz-conventional-changelog

# Usar con:
git cz
```

### Git Hooks:
```bash
# Pre-commit hook para validar formato
#!/bin/sh
commit_regex='^(feat|fix|docs|style|refactor|perf|test|build|ci|chore|revert)(\(.+\))?: .{1,100}'

if ! grep -qE "$commit_regex" "$1"; then
    echo "❌ Invalid commit message format!"
    echo "Use: type(scope): description"
    exit 1
fi
```

## ❓ Preguntas Frecuentes

### Q: ¿Puedo usar mayúsculas en la descripción?
**A:** No, mantén todo en minúsculas excepto nombres propios.

### Q: ¿Qué pasa si mi cambio afecta múltiples áreas?
**A:** Usa el área más significativa o divide en múltiples PRs.

### Q: ¿Cuándo usar `chore` vs `feat`?
**A:** `feat` para funcionalidades que afectan usuarios, `chore` para tareas internas.

### Q: ¿Puedo cambiar el título después de crear el PR?
**A:** Sí, la validación se ejecuta cada vez que editas el título.

### Q: ¿Qué pasa si el validador falla?
**A:** El PR no se puede mergear hasta que el título sea válido.

---

**¿Necesitas ayuda?** El sistema automáticamente te proporcionará sugerencias específicas cuando tu título no sea válido. ¡Simplemente sigue las recomendaciones! 🎉
