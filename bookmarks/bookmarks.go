package bookmarks

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/glebarez/sqlite"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

var (
	Db *gorm.DB

	browsers = map[string]string{
		"chrome":  "",
		"firefox": "",
		"edge":    "",
	}

	Command = &cli.Command{
		Name:    "bookmarks",
		Usage:   "将书签导出到文件中",
		Aliases: []string{"b"},
		Flags:   []cli.Flag{
			&cli.StringFlag{
				Name:     "browser",
				Aliases:  []string{"b"},
				Required: true,
				Usage:    "支持浏览器 chrome、edge 或 firefox",
			},
		},
		Action:  Action,
	}
)

type (
	BookmarkRoot struct {
		Roots struct {
			BookmarkBar Bookmark `json:"bookmark_bar"`
		} `json:"roots"`
	}

	Bookmark struct {
		Name     string     `json:"name"`
		URL      string     `json:"url,omitempty"`
		Children []Bookmark `json:"children,omitempty"`
	}

	MozBookmarks struct {
		Name string `gorm:"name" json:"name"`
		Url  string `gorm:"url" json:"url"`
	}
)

func Action(c *cli.Context) error {
	browser := c.String("browser")
	if _, ok := browsers[browser]; !ok {
		return fmt.Errorf(color.RedString("不支持该浏览器"))
	}

	filename := fmt.Sprintf("%s_bookmarks.json", browser)
	f, _ := os.Create(filename)
	defer f.Close()

	switch browser {
	case "chrome":
		bookmarks := getChromeBookmarks()
		if len(bookmarks) == 0 {
			return fmt.Errorf(color.RedString("未找到书签文件"))
		}

		b, err := ioutil.ReadFile(bookmarks)
		if err != nil {
			return fmt.Errorf(color.RedString("读取书签文件失败：%w", err))
		}

		var bookmarkRoot BookmarkRoot
		if err = json.Unmarshal(b, &bookmarkRoot); err != nil {
			return fmt.Errorf(color.RedString("解析书签文件失败：%w", err))
		}

		b, _ = json.Marshal(bookmarkRoot.Roots.BookmarkBar)
		_, _ = f.Write(b)

	case "edge":
		bookmarks := getEdgeBookmarks()
		if len(bookmarks) == 0 {
			return fmt.Errorf(color.RedString("未找到书签文件"))
		}

		b, err := ioutil.ReadFile(bookmarks)
		if err != nil {
			return fmt.Errorf(color.RedString("读取书签文件失败：%w", err))
		}

		var bookmarkRoot BookmarkRoot
		if err = json.Unmarshal(b, &bookmarkRoot); err != nil {
			return fmt.Errorf(color.RedString("解析书签文件失败：%w", err))
		}

		b, _ = json.Marshal(bookmarkRoot.Roots.BookmarkBar)
		_, _ = f.Write(b)

	case "firefox":
		bookmarks := getFirefoxBookmarks()
		if len(bookmarks) == 0 {
			return fmt.Errorf(color.RedString("未找到书签文件"))
		}

		initDb(bookmarks)
		defer closeDb()

		var bookmark []MozBookmarks
		if err := Db.Raw(`SELECT moz_bookmarks.title as name,
			   moz_places.url
		  FROM moz_bookmarks
			   LEFT JOIN
			   moz_places ON moz_bookmarks.fk = moz_places.id
		 WHERE type = 1 AND 
			   name NOT IN ('帮助和教程', '自定义 Firefox', '加入进来', '关于我们', '最近使用的标签', '获取帮助','参与进来')
		`).Scan(&bookmark).Error; err != nil {
			return fmt.Errorf(color.RedString("查询书签失败：%w", err))
		}

		b, _ := json.Marshal(bookmark)
		_, _ = f.Write(b)
	}

	fmt.Println(color.GreenString("书签文件已导出到当前目录下的 %s 文件中", filename))
	return nil

}

func getChromeBookmarks() string {
	switch runtime.GOOS {
	case "windows":
		bookmarks, _ := homedir.Expand("~/AppData/Local/Google/Chrome/User Data/Default/Bookmarks")
		return bookmarks
	case "darwin":
		bookmarks, _ := homedir.Expand("~/Library/Application Support/Google/Chrome/Default/Bookmarks")
		return bookmarks
	case "linux":
		bookmarks, _ := homedir.Expand("~/.config/google-chrome/Default/Bookmarks")
		return bookmarks
	default:
		return ""
	}
}

func getEdgeBookmarks() string {
	switch runtime.GOOS {
	case "windows":
		bookmarks, _ := homedir.Expand("~/AppData/Local/Microsoft/Edge/User Data/Default/Bookmarks")
		return bookmarks
	case "darwin":
		bookmarks, _ := homedir.Expand("~/Library/Application Support/Microsoft Edge/Default/Bookmarks")
		return bookmarks
	default:
		return ""
	}
}

func getFirefoxBookmarks() string {
	var bookmarks string
	switch runtime.GOOS {
	case "windows":
		bookmarks, _ = homedir.Expand("~/AppData/Roaming/Mozilla/Firefox/Profiles")
	case "darwin":
		bookmarks, _ = homedir.Expand("~/Library/Application Support/Firefox/Profiles")
	case "linux":
		bookmarks, _ = homedir.Expand("~/.mozilla/firefox")
	default:
		return ""
	}

	_ = filepath.Walk(bookmarks, func(path string, info os.FileInfo, err error) error {
		if info.Name() == "places.sqlite" {
			bookmarks = path
		}
		return nil
	})

	return bookmarks
}

func initDb(dsn string) {
	db, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		panic(fmt.Errorf(color.RedString("连接数据库失败：%w", err)))
	}

	db.Logger = logger.Default.LogMode(logger.Silent)
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	Db = db
}

func closeDb() {
	sqliteDb, err := Db.DB()
	if err != nil {
		panic(fmt.Errorf(color.RedString("关闭数据库失败：%w", err)))
	}

	if err = sqliteDb.Close(); err != nil {
		panic(fmt.Errorf(color.RedString("关闭数据库失败：%w", err)))
	}
}
