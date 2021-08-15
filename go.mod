module github.com/n-guitar/build-ops-container

go 1.16

require (
	github.com/gofiber/fiber/v2 v2.17.0
	gorm.io/gorm v1.20.7
	gorm.io/driver/sqlite v1.1.4
)

replace (
	github.com/n-guitar/build-ops-container/pkg/buildapi => ./pkg/buildapi
	github.com/n-guitar/build-ops-container/pkg/database => ./pkg/database
	github.com/n-guitar/build-ops-container/pkg/gitcmd => ./pkg/gitcmd
)
