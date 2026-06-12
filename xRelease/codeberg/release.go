package codeberg

import (
	"slices"
	"strings"

	"github.com/unix755/xtools/xEncoding"
)

type Release struct {
	ID         int      `json:"id"`
	TagName    string   `json:"tag_name"`
	Name       string   `json:"name"`
	TarballURL string   `json:"tarball_url"`
	ZipballURL string   `json:"zipball_url"`
	Assets     []Assets `json:"assets"`
}

type Assets struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
	Type               string `json:"type"`
}

func newRelease[T *[]Release | *Release](releaseApiUrl string) (r T, err error) {
	return r, xEncoding.URLToTarget(&r, "json", releaseApiUrl)
}

func GetReleases(repo string) (rs *[]Release, err error) {
	return newRelease[*[]Release]("https://codeberg.org/api/v1/repos/" + repo + "/releases")
}

func GetReleaseLatest(repo string) (r *Release, err error) {
	return newRelease[*Release]("https://codeberg.org/api/v1/repos/" + repo + "/releases/latest")
}

func GetReleaseByTagName(repo string, tagName string) (r *Release, err error) {
	return newRelease[*Release]("https://codeberg.org/api/v1/repos/" + repo + "/releases/tags/" + tagName)
}

func (r *Release) GetAssets(includes []string, excludes []string) (assets []Assets) {
	// 排除不包含
	for _, exclude := range excludes {
		r.Assets = slices.DeleteFunc(r.Assets, func(assets Assets) bool {
			return strings.Contains(assets.Name, exclude)
		})
	}
	// 寻找所有全包含项目
	for _, include := range includes {
		r.Assets = slices.DeleteFunc(r.Assets, func(assets Assets) bool {
			return !strings.Contains(assets.Name, include)
		})
	}
	return r.Assets
}
