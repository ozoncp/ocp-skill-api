package internal

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-skill-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-skill-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/saver_mock.go -package=mocks github.com/ozoncp/ocp-skill-api/internal/saver Saver
