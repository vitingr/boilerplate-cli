package config

var Catalog = map[string]map[string][]string{
	"node": {
		"nest": {
			"ddd",
			"hexagonal",
			"clean-architecture",
			"monolith",
			"microservice",
		},
		"fastify": {
			"rest-api",
			"graphql",
			"microservice",
			"monorepo",
		},
	},
	"golang": {
		"fiber": {
			"rest-api",
			"hexagonal",
			"clean-architecture",
			"cli",
		},
		"gin": {
			"rest-api",
			"crud",
			"microservice",
		},
		"echo": {
			"rest-api",
			"hexagonal",
		},
		"stdlib": {
			"rest-api",
			"cli",
		},
	},
	"python": {
		"fastapi": {
			"rest-api",
			"hexagonal",
			"async-worker",
		},
		"django": {
			"monolith",
			"rest-api",
			"ddd",
		},
		"flask": {
			"rest-api",
			"microservice",
		},
	},
}

const RepoBaseURL = "https://github.com/yourusername/boilerplates"

type ScaffoldConfig struct {
	ProjectName string
	Language    string
	Framework   string
	Template    string
	OutputDir   string
	InitGit     bool
	InstallDeps bool
}
