package buildapi

/*
Container imageのbuildをapi経由で行うための関数を定義する
build用に利用するデータのdatabase操作を行う
*/

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n-guitar/build-ops-container/pkg/database"
	"gorm.io/gorm"
)

// build用に利用するデータの構造体
type BuildData struct {
	gorm.Model
	BuildName string `json:"build_name"`
	GitRepo   string `json:"git_repo"`
	ImgTag    string `json:"img_tag"`
}

func GetBuildDataSet(c *fiber.Ctx) error {
	db := database.DBConn
	var dataset []BuildData
	db.Find(&dataset)
	// build 配列をすばやく簡単に JSON 文字列にシリアル化して応答
	return c.JSON(dataset)
}

// ToDo buildがないときの処理
func GetBuildData(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var build BuildData
	db.Find(&build, id)
	return c.JSON(build)
}

func NewBuildData(c *fiber.Ctx) error {
	db := database.DBConn
	var build BuildData
	// formデータをセット
	buildname := c.FormValue("buildname")
	gitrepo := c.FormValue("gitrepo")
	imgtag := c.FormValue("imgtag")
	build.BuildName = buildname
	build.GitRepo = gitrepo
	build.ImgTag = imgtag
	// databaseへレコード追加
	db.Create(&build)
	return c.JSON(build)
}

func DeleteBuildData(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var build BuildData
	db.First(&build, id)
	if build.BuildName == "" {
		return c.Status(500).SendString("No Build Found with ID")
	}
	db.Delete(&build)
	return c.SendString("Build Seccessfully Deleted")
}
