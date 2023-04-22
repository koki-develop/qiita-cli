package notify

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/google/go-github/v51/github"
	"github.com/koki-develop/qiita-cli/internal/util"
)

type cache struct {
	Version    string    `json:"version,omitempty"`
	Expiration time.Time `json:"expiration,omitempty"`
}

func NotifyNewRelease(w io.Writer, version string) error {
	c, err := loadCache()
	if err != nil {
		return err
	}

	// if the latest release version is cached and not expired, return
	if c != nil && c.Expiration.After(time.Now()) {
		return nil
	}

	// fetch latest release version
	cl := github.NewClient(nil)
	release, _, err := cl.Repositories.GetLatestRelease(context.Background(), "koki-develop", "qiita-cli")
	if err != nil {
		return err
	}
	// save latest release version
	c = &cache{Version: *release.TagName, Expiration: time.Now().Add(24 * time.Hour)}
	if err := saveCache(c); err != nil {
		return err
	}

	// if the latest release version is newer than the current version, print a message
	if util.Version(c.Version).Newer(util.Version(version)) {
		color.New(color.Bold).Fprintf(w, "A new version (%s) of qiita-cli is available!\n", c.Version)
		fmt.Fprint(w, "See: ")
		color.New(color.Underline).Fprintf(w, "https://github.com/koki-develop/qiita-cli/releases/tag/%s\n", c.Version)
		fmt.Fprintln(w)
	}

	return nil
}

func cachePath() (string, error) {
	d, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(d, "qiita-cli", "release.json"), nil
}

func saveCache(c *cache) error {
	p, err := cachePath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(c); err != nil {
		return err
	}

	return nil
}

func loadCache() (*cache, error) {
	p, err := cachePath()
	if err != nil {
		return nil, err
	}

	ext, err := util.Exists(p)
	if err != nil {
		return nil, err
	}
	if !ext {
		return nil, nil
	}

	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var c cache
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
