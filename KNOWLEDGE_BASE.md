# Polyglot Codebase Knowledge Graph

> Generated offline by **readmenator**. Supports C, C++, Python, Go, Rust, JS/TS, Java, C#, Shell, PHP, Dart, GDScript, Nim, ASM.
> No LLMs. No tokens. Pure static analysis.

**Total Files Parsed:** 4 | **Total Symbols Extracted:** 3 | **Total Imports:** 3

## Structural Knowledge Map
```mermaid
graph TD
    classDef mod fill:#1e1e1e,stroke:#ff6666,stroke-width:2px,color:#fff;
    classDef cls fill:#2d2d2d,stroke:#4ec9b0,stroke-width:2px,color:#fff;
    classDef fn fill:#333,stroke:#dcdcaa,stroke-width:1px,color:#dcdcaa;
    classDef ext fill:#111,stroke:#666,stroke-dasharray: 5 5,color:#aaa;
    main_go["main.go (go)"]
    class main_go mod;
    main_go_add["add"]
    class main_go_add fn;
    main_go --> main_go_add
    main_go_main["main"]
    class main_go_main fn;
    main_go --> main_go_main
    shell_go["shell.go (go)"]
    class shell_go mod;
    shell_go_main["main"]
    class shell_go_main fn;
    shell_go --> shell_go_main
    app_py["app.py (py)"]
    class app_py mod;
    install_sh["install.sh (sh)"]
    class install_sh mod;
    ext_os["os"]
    class ext_os ext;
    app_py -.->|imports| ext_os
    ext_fmt["fmt"]
    class ext_fmt ext;
    main_go -.->|imports| ext_fmt
    ext_C["C"]
    class ext_C ext;
    shell_go -.->|imports| ext_C
```

---

## Architecture Reference

### GO (2 files)

#### `main.go`
**Path:** `main.go`

**Functions:**
- `add` (line 6) - *go:noinline*
- `main` (line 8)

#### `shell.go`
**Path:** `shell.go`

**Functions:**
- `main` (line 36)

### PY (1 files)

#### `app.py`
**Path:** `app.py`

*No symbols extracted*

### SH (1 files)

#### `install.sh`
**Path:** `install.sh`

*No symbols extracted*
