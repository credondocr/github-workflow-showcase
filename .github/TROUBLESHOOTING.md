# 🛠️ GitHub Actions Troubleshooting Guide

Esta guía te ayuda a resolver problemas comunes con los workflows de GitHub Actions del proyecto.

## 🔧 Problemas Comunes y Soluciones

### ❌ Error: Deprecated actions/upload-artifact: v3

**Problema:**
```
Error: This request has been automatically failed because it uses a deprecated version of `actions/upload-artifact: v3`. 
Learn more: https://github.blog/changelog/2024-04-16-deprecation-notice-v3-of-the-artifact-actions/
```

**Causa:** 
GitHub Actions v3 para upload/download-artifact están deprecadas desde abril 2024.

**Solución:**
```yaml
# ❌ Deprecado
uses: actions/upload-artifact@v3
uses: actions/download-artifact@v3
uses: actions/create-release@v1

# ✅ Actualizado  
uses: actions/upload-artifact@v4
uses: actions/download-artifact@v4
# create-release@v1 reemplazado con 'gh release create'
```

**Cambios importantes en v4:**
- Usa Node.js 20 en lugar de 16
- Mejor performance y compatibilidad
- APIs mejoradas para manejo de artifacts

### ❌ Error: Invalid format in GITHUB_OUTPUT

**Problema:**
```
Error: Unable to process file command 'output' successfully.
Error: Invalid format '0'
```

**Causa:** 
Sintaxis incorrecta en comandos `echo >> $GITHUB_OUTPUT` con comandos complejos que pueden fallar.

**Problema original:**
```bash
# ❌ Problemático - grep -c puede fallar con exit code 1
echo "count=$(git diff | grep -c pattern || echo 0)" >> $GITHUB_OUTPUT
```

**Solución:**
```bash
# ✅ Correcto - usar variables temporales
COUNT=$(git diff | grep pattern | wc -l || echo 0)
echo "count=$COUNT" >> $GITHUB_OUTPUT

# O usar alternativa más robusta:
COUNT=$(git diff | grep -c pattern 2>/dev/null || echo 0)
echo "count=$COUNT" >> $GITHUB_OUTPUT
```

**Mejores prácticas para GITHUB_OUTPUT:**
- Usar variables temporales para comandos complejos
- Evitar `grep -c` dentro de `$(...)` 
- Usar `wc -l` en lugar de `grep -c`
- Manejar exit codes con `|| echo 0`

### ❌ Error: Cannot find module 'js-yaml'

**Problema:**
```
Error: Cannot find module 'js-yaml'
Require stack:
- /home/runner/work/_actions/actions/github-script/v6/dist/index.js
```

**Causa:** 
`github-script@v6` no incluye la dependencia `js-yaml` por defecto.

**Soluciones:**

#### ✅ Solución 1: Actualizar a github-script@v7
```yaml
# Cambiar de:
uses: actions/github-script@v6

# A:
uses: actions/github-script@v7
```

#### ✅ Solución 2: Parser YAML manual (implementado)
```yaml
- name: Parse YAML manually
  uses: actions/github-script@v7
  with:
    script: |
      const fs = require('fs');
      const yamlContent = fs.readFileSync('.github/labels.yml', 'utf8');
      
      // Simple YAML parser for our specific format
      function parseLabelsYaml(content) {
        const labels = [];
        const lines = content.split('\n');
        let currentLabel = {};
        
        for (const line of lines) {
          const trimmed = line.trim();
          if (trimmed.startsWith('- name:')) {
            if (currentLabel.name) {
              labels.push(currentLabel);
            }
            currentLabel = { name: trimmed.replace('- name:', '').replace(/"/g, '').trim() };
          } else if (trimmed.startsWith('color:')) {
            currentLabel.color = trimmed.replace('color:', '').replace(/"/g, '').trim();
          } else if (trimmed.startsWith('description:')) {
            currentLabel.description = trimmed.replace('description:', '').replace(/"/g, '').trim();
          }
        }
        if (currentLabel.name) {
          labels.push(currentLabel);
        }
        return labels;
      }
      
      const labelsConfig = parseLabelsYaml(yamlContent);
```

#### ✅ Solución 3: Usar yq para parsear YAML
```yaml
- name: Parse YAML with yq
  id: parse-yaml
  uses: mikefarah/yq@master
  with:
    cmd: yq eval -o=json '.[]' .github/labels.yml

- name: Use parsed data
  uses: actions/github-script@v7
  with:
    script: |
      const labelsConfig = `${{ steps.parse-yaml.outputs.result }}`.split('\n')
        .filter(line => line.trim())
        .map(line => JSON.parse(line));
```

#### ✅ Solución 4: Acción dedicada (recomendada)
```yaml
- name: Sync labels
  uses: crazy-max/ghaction-github-labeler@v5
  with:
    github-token: ${{ secrets.GITHUB_TOKEN }}
    yaml-file: .github/labels.yml
```

### ❌ Error: API rate limit exceeded

**Problema:**
```
Error: API rate limit exceeded
```

**Soluciones:**

#### ✅ Añadir delays entre llamadas
```yaml
script: |
  for (const item of items) {
    await processItem(item);
    // Add delay to respect rate limits
    await new Promise(resolve => setTimeout(resolve, 1000));
  }
```

#### ✅ Usar batch operations
```yaml
script: |
  // Process in batches instead of individual calls
  const batchSize = 10;
  for (let i = 0; i < items.length; i += batchSize) {
    const batch = items.slice(i, i + batchSize);
    await Promise.all(batch.map(item => processItem(item)));
    await new Promise(resolve => setTimeout(resolve, 2000));
  }
```

### ❌ Error: Resource not accessible by integration

**Problema:**
```
RequestError [HttpError]: Resource not accessible by integration
status: 403
url: 'https://api.github.com/repos/.../issues/.../comments'
```

**Causa:** 
El `GITHUB_TOKEN` no tiene permisos suficientes para acceder al recurso (comentar en PRs, crear labels, etc.).

**Problema original:**
```yaml
# ❌ Sin permisos definidos
name: My Workflow
on: [pull_request]
jobs:
  my-job:
    runs-on: ubuntu-latest
    steps: # ... intentará comentar pero fallará
```

**Soluciones:**

#### ✅ Definir permisos a nivel de workflow
```yaml
name: My Workflow
on: [pull_request]

permissions:
  contents: read          # Para leer el código
  pull-requests: write    # Para comentar en PRs
  issues: write          # Para comentar en issues
  actions: read          # Para leer resultados de workflows
  checks: write          # Para crear check runs

jobs:
  my-job:
    runs-on: ubuntu-latest
    steps: # ... ahora funcionará
```

#### ✅ Definir permisos a nivel de job
```yaml
name: My Workflow
on: [pull_request]

jobs:
  my-job:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      pull-requests: write
      issues: write
    steps: # ... funcionará solo para este job
```

#### ✅ Permisos específicos según el uso
```yaml
# Para workflows que solo leen código
permissions:
  contents: read

# Para workflows que comentan en PRs
permissions:
  contents: read
  pull-requests: write
  issues: write

# Para workflows que crean releases
permissions:
  contents: write
  pull-requests: write
  actions: read

# Para workflows que manejan labels
permissions:
  contents: read
  issues: write
  pull-requests: write
```

**Permisos disponibles:**
- `contents`: read/write - Acceso al código del repositorio
- `pull-requests`: read/write - Acceso a PRs y comentarios
- `issues`: read/write - Acceso a issues y comentarios  
- `actions`: read/write - Acceso a workflows y artifacts
- `checks`: read/write - Acceso a check runs y status
- `deployments`: read/write - Acceso a deployments
- `packages`: read/write - Acceso a packages/registry

### ❌ Error: Workflow not triggering

**Problema:** El workflow no se ejecuta cuando debería.

**Soluciones:**

#### ✅ Verificar triggers
```yaml
on:
  pull_request:
    types: [opened, synchronize, reopened]  # Especificar tipos
    branches: [ main, develop ]             # Especificar branches
    paths: [ '.github/**' ]                 # Especificar paths si es necesario
```

#### ✅ Verificar permisos del repositorio
- Settings → Actions → General → Workflow permissions
- Debe estar en "Read and write permissions"

### ❌ Error: Invalid YAML syntax

**Problema:**
```
Invalid workflow file: .github/workflows/example.yml
```

**Soluciones:**

#### ✅ Validar YAML online
- Usar [yamllint.com](https://yamllint.com)
- VS Code extension: "YAML"

#### ✅ Común errores de sintaxis:
```yaml
# ❌ Incorrecto - indentación mixta
  - name: Step 1
      run: echo "hello"

# ✅ Correcto - indentación consistente
  - name: Step 1
    run: echo "hello"

# ❌ Incorrecto - quotes no balanceadas  
  - name: 'Step with quote"
    
# ✅ Correcto
  - name: 'Step with quote'
```

### ❌ Error: Step timeout

**Problema:**
```
Error: The operation was canceled.
```

**Solución:**
```yaml
- name: Long running step
  timeout-minutes: 30  # Default is 360 (6 hours)
  run: |
    # Your long running command
```

### ❌ Error: Matrix job failures

**Problema:** Algunos jobs de matrix fallan.

**Solución:**
```yaml
strategy:
  matrix:
    os: [ubuntu-latest, windows-latest, macos-latest]
  fail-fast: false  # Continue other jobs even if one fails
```

## 🔍 Debugging Workflows

### 1. **Enable Debug Logging**
```yaml
env:
  ACTIONS_STEP_DEBUG: true
  ACTIONS_RUNNER_DEBUG: true
```

### 2. **Add Debug Steps**
```yaml
- name: Debug context
  uses: actions/github-script@v7
  with:
    script: |
      console.log('Context:', JSON.stringify(context, null, 2));
      console.log('GitHub:', JSON.stringify(github, null, 2));
```

### 3. **Test Locally**
```bash
# Install act for local testing
curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash

# Run workflow locally
act -n  # Dry run
act     # Run actual workflow
```

## 📋 Mejores Prácticas

### 1. **Version Pinning**
```yaml
# ✅ Bueno - pin to specific version
uses: actions/checkout@v4

# ❌ Evitar - floating versions
uses: actions/checkout@main
```

### 2. **Error Handling**
```yaml
- name: Step with error handling
  id: my-step
  continue-on-error: true
  run: |
    # Command that might fail
    
- name: Handle failure
  if: steps.my-step.outcome == 'failure'
  run: |
    echo "Previous step failed, handling gracefully"
```

### 3. **Secrets Management**
```yaml
# ✅ Correcto
env:
  API_KEY: ${{ secrets.API_KEY }}

# ❌ Nunca hardcodear secrets
env:
  API_KEY: "abc123"
```

### 4. **Cache Optimization**
```yaml
- name: Cache dependencies
  uses: actions/cache@v3
  with:
    path: |
      ~/.cache/go-build
      ~/go/pkg/mod
    key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
    restore-keys: |
      ${{ runner.os }}-go-
```

### 5. **Versiones Actualizadas de Acciones**
```yaml
# ✅ Versiones actuales recomendadas (2024)
uses: actions/checkout@v4
uses: actions/setup-go@v5
uses: actions/upload-artifact@v4
uses: actions/download-artifact@v4
uses: actions/cache@v4
uses: golangci/golangci-lint-action@v6
uses: github/codeql-action/upload-sarif@v3
uses: codecov/codecov-action@v4
uses: actions/github-script@v7
uses: docker/setup-buildx-action@v3
uses: docker/login-action@v3

# ❌ Versiones deprecadas (evitar)
uses: actions/upload-artifact@v3      # Deprecated abril 2024
uses: actions/create-release@v1       # Discontinued
uses: actions/github-script@v6        # Missing dependencies
uses: golangci/golangci-lint-action@v3 # Outdated
```

## 📞 Getting Help

### 1. **GitHub Community**
- [GitHub Community Discussions](https://github.com/community/community/discussions)
- [Actions Marketplace](https://github.com/marketplace?type=actions)

### 2. **Documentation**
- [GitHub Actions Docs](https://docs.github.com/en/actions)
- [Workflow Syntax](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions)

### 3. **Tools**
- [Actions Toolkit](https://github.com/actions/toolkit)
- [Act (Local Testing)](https://github.com/nektos/act)
- [VS Code Extension](https://marketplace.visualstudio.com/items?itemName=github.vscode-github-actions)

## 🚀 Workflow Health Check

Para verificar que todos los workflows estén funcionando:

```bash
# Listar todos los workflows
gh workflow list

# Ver status de runs recientes
gh run list

# Ver detalles de un run específico
gh run view [run-id]

# Re-run failed jobs
gh run rerun [run-id] --failed
```

---

💡 **Tip:** Siempre prueba los workflows en un branch separado antes de mergear a main para evitar problemas en producción.
