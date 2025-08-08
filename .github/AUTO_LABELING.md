# 🏷️ Auto-Labeling System

This repository uses an intelligent auto-labeling system that automatically categorizes pull requests based on the code changes, patterns, and content. This helps maintain organization and improves the development workflow.

## 📋 How It Works

The auto-labeling system consists of three main workflows:

### 1. 🎯 Area-Based Labeling (`auto-labeler.yml`)

Automatically adds labels based on **which files were modified**:

| Files Changed | Label Applied | Description |
|---------------|---------------|-------------|
| `models/` | `area: models` | Changes to data models and business logic |
| `controllers/` | `area: controllers` | Changes to HTTP request handlers |
| `routes/` | `area: routes` | Changes to API route configuration |
| `tests/` or `*_test.go` | `area: tests` | Changes to test files |
| `.github/` | `area: ci/cd` | Changes to CI/CD workflows |
| `*.md`, `*.txt` | `area: documentation` | Changes to documentation |
| `go.mod`, `go.sum` | `area: dependencies` | Changes to dependencies |
| `*.yml`, `*.yaml`, `*.json` | `area: configuration` | Changes to config files |
| `main.go` | `area: core` | Changes to core application |
| `Dockerfile` | `area: docker` | Changes to Docker configuration |

### 2. 📏 Size-Based Labeling

Automatically categorizes PRs by size:

| Lines Changed | Label | Color | Description |
|---------------|-------|-------|-------------|
| < 10 | `size: XS` | 🟢 Green | Extra small PR |
| < 30 | `size: S` | 🟡 Yellow-Green | Small PR |
| < 100 | `size: M` | 🟡 Gold | Medium PR |
| < 500 | `size: L` | 🟠 Orange | Large PR |
| ≥ 500 | `size: XL` | 🔴 Red | Extra large PR |

### 3. 🔍 Type Detection (`pr-type-detector.yml`)

Intelligently detects PR type based on **title, description, and change patterns**:

#### Detection Patterns:

| Type | Triggers | Examples |
|------|----------|----------|
| `type: feature` | `feat:`, `add:`, `implement`, `new` | "feat: add user authentication" |
| `type: bugfix` | `fix:`, `bug:`, `resolve`, `closes #` | "fix: resolve login issue" |
| `type: enhancement` | `improve:`, `enhance:`, `update:` | "improve: better error handling" |
| `type: refactor` | `refactor:`, `cleanup`, `restructure` | "refactor: reorganize models" |
| `type: performance` | `perf:`, `optimize`, `speed up` | "perf: optimize database queries" |
| `type: security` | `security:`, `vulnerability`, `secure` | "security: fix XSS vulnerability" |
| `type: breaking` | `breaking:`, API changes, major deletions | "breaking: change API structure" |

## 🎨 Label Color Scheme

Our labels follow a consistent color scheme for easy visual identification:

- **🟢 Green**: Positive/Safe changes (XS size, models, ready to merge)
- **🔵 Blue**: Features and controllers  
- **🔴 Red**: Critical/Breaking changes (XL size, breaking, security)
- **🟡 Yellow**: Tests, documentation, needs attention
- **🟣 Purple**: CI/CD, refactoring, workflow changes
- **🟠 Orange**: Large changes, performance, enhancements

## 🔄 Automatic Label Management

### Label Creation
- Labels are automatically created if they don't exist
- Color and description are standardized
- No manual setup required

### Label Updates
- Old size labels are automatically removed when new ones are added
- Labels are updated based on the latest analysis
- Duplicate labels are prevented

### Label Synchronization
- `sync-labels.yml` workflow maintains label consistency
- Runs when `.github/labels.yml` is modified
- Ensures all labels match the configuration

## 📊 PR Analysis Features

### Automatic Comments
When a PR is opened, the system adds a comprehensive analysis comment:

```markdown
🤖 Auto-Labeler Analysis

🟡 PR Size: Medium (87 total line changes)

📊 Changes Summary:
- 📁 Files changed: 5
- ➕ Lines added: 45
- ➖ Lines deleted: 42

🎯 Areas Modified:
- 🎮 Controllers
- 🧪 Tests
- 📚 Documentation
```

### Breaking Change Detection
Special handling for potentially breaking changes:
- Automatic detection based on patterns
- Warning comments on PRs
- Checklist for breaking change requirements
- Red label for high visibility

## 🛠️ Advanced Features

### Smart Detection Rules

1. **Test Coverage Check**: Flags PRs with Go code changes but no test updates
2. **Documentation Updates**: Detects when docs should be updated
3. **Dependency Analysis**: Tracks dependency changes and security implications  
4. **File Pattern Recognition**: Uses intelligent patterns, not just folder names

### Context-Aware Labeling
- Considers PR title, description, and code changes together
- Uses regex patterns for flexible matching
- Handles edge cases and special scenarios

## 📝 Best Practices for Contributors

### PR Titles
Use conventional commit format for better detection:
- ✅ `feat: add user registration endpoint`
- ✅ `fix: resolve authentication bug #123`  
- ✅ `docs: update API documentation`
- ✅ `refactor: improve error handling`

### PR Descriptions
Include relevant keywords:
- For features: "implement", "new functionality", "add capability"
- For fixes: "resolve", "fix issue", "closes #123"
- For breaking changes: explicitly mention "breaking change"

### File Organization
Keep changes focused on specific areas when possible:
- Model changes in `models/`
- Controller logic in `controllers/`
- Route definitions in `routes/`
- Tests in `tests/` or `*_test.go`

## 🔧 Customization

### Adding New Labels
1. Update `.github/labels.yml`
2. Add detection logic in workflow files
3. Run the sync-labels workflow

### Modifying Detection Patterns
Edit the pattern arrays in:
- `auto-labeler.yml` for file-based detection
- `pr-type-detector.yml` for content-based detection

### Custom Rules
The system supports:
- Custom regex patterns
- File path matching
- Content analysis
- Combination rules

## 🚀 Benefits

- **🔍 Automatic Organization**: No manual labeling required
- **👀 Quick Overview**: Instant visual categorization of PRs
- **📈 Better Metrics**: Track areas of development activity
- **⚡ Faster Reviews**: Reviewers can prioritize by size and type
- **🎯 Focused Attention**: Breaking changes get special handling
- **📊 Project Insights**: Analyze development patterns over time

## 🎯 Example Scenarios

### Scenario 1: New Feature
```bash
# PR Title: "feat: implement user profile management"
# Files: controllers/profile.go, models/user.go, tests/profile_test.go
# Result: area: controllers, area: models, area: tests, type: feature, size: M
```

### Scenario 2: Bug Fix  
```bash
# PR Title: "fix: resolve database connection timeout"
# Files: models/database.go
# Result: area: models, type: bugfix, size: S
```

### Scenario 3: Documentation Update
```bash
# PR Title: "docs: add API endpoint documentation" 
# Files: README.md, docs/api.md
# Result: area: documentation, type: enhancement, size: XS
```

This system ensures consistent, accurate, and helpful labeling without any manual intervention! 🎉
