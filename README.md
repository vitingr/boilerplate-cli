# boilerplate-cli (`bplt`)

Interactive CLI to scaffold projects from your boilerplate GitHub repo.

## Install

```bash
go install github.com/vitingr/boilerplate-cli@latest
```

Or build locally:

```bash
git clone https://github.com/vitingr/boilerplate-cli
cd boilerplate-cli
go build -o bplt .
sudo mv bplt /usr/local/bin/
```

## Usage

```bash
# fully interactive
bplt create

# supply project name up-front
bplt create my-api
```

The CLI will ask:

1. **Language** — node / golang / python  
2. **Framework** — nest / fastify / fiber / gin / …  
3. **Architecture** — ddd / hexagonal / clean-architecture / …  
4. **Output directory** (defaults to `./project-name`)  
5. **git init?**  
6. **Install dependencies?**

## Adding new boilerplates

Edit `internal/config/config.go` → `Catalog` map.  
The path must mirror your GitHub repo: `/<language>/<framework>/<template>/`.

## Requirements

- Git ≥ 2.25 (sparse-checkout support)
- Go ≥ 1.25
