package stats

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
)

type GitStats struct {
	Add     int
	Del     int
	Author  string
	Since   string
	Dir     string
	ScanNum int
}

func NewGitStats(author, since, dir string) *GitStats {
	return &GitStats{Author: author, Since: since, Dir: dir}
}

// GetStats gets the git stats of all git repositories under the path.
func (stats *GitStats) GetStats() error {
	err := filepath.Walk(stats.Dir, func(path string, file os.FileInfo, err error) error {
		if file.Name() == ".git" {
			stats.ScanNum++
			return stats.CalculateLinesStats(path)
		}
		return nil
	})

	if err != nil {
		return err
	}
	pterm.Success.Println(fmt.Sprintf("Scanned repositories:%d  ", stats.ScanNum) +
		pterm.LightGreen(fmt.Sprintf("Add lines(+):%d  ", stats.Add)) +
		pterm.Red(fmt.Sprintf("Del lines(-):%d  ", stats.Del)) +
		pterm.Yellow(fmt.Sprintf("Total lines:%d", stats.Add-stats.Del)))

	return nil

}

// CalculateLinesStats calculates the lines added and deleted in the git repository.
func (stats *GitStats) CalculateLinesStats(path string) error {
	spinner := NewInfoSpinner("Calculating " + path)
	pathArgs := ` --git-dir="` + path + `" `
	var authorArgs string = ""
	if stats.Author != "" {
		authorArgs = `--author="` + stats.Author + `" `
	}
	var sinceArgs string = ""
	if stats.Since != "" {
		sinceArgs = ` --since="` + stats.Since + `" `
	}

	cmd :=
		`git` + pathArgs + ` log ` + authorArgs + sinceArgs + ` --pretty=tformat: --numstat | 
        awk '{ add += $1; del += $2; } 
        function defaultVal(var){ return var == "" ? 0 : var } 
        END { printf "%s,%s", defaultVal(add), defaultVal(del) }'`

	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		pterm.Error.Printf("%s: %s\n", err, out)
		return err
	}

	str := strings.Split(string((out)), ",")
	add, _ := strconv.Atoi(str[0])
	stats.Add += add
	del, _ := strconv.Atoi(str[1])
	stats.Del += del

	spinner.Info(fmt.Sprintf("Path:%s  ", path) +
		pterm.LightGreen(fmt.Sprintf("Add lines(+):%d  ", add)) +
		pterm.Red(fmt.Sprintf("Del lines(-):%d", del)))

	return nil
}
