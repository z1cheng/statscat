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
    Add         int
    Del         int
    Author      string
    Since       string
    Dir         string
    ScanNum     int
    CommitCount int
    gitArguments
}

type gitArguments struct {
    pathArgs   string
    authorArgs string
    sinceArgs  string
}

func NewGitStats(author, since, dir string) *GitStats {
    return &GitStats{Author: author, Since: since, Dir: dir}
}

// GetStats gets the git stats of all git repositories under the path.
func (stats *GitStats) GetStats() error {
    // expand git args
    if stats.Author != "" {
        stats.gitArguments.authorArgs = `--author="` + stats.Author + `" `
    }
    if stats.Since != "" {
        stats.gitArguments.sinceArgs = ` --since="` + stats.Since + `" `
    }

    // go through all directories and calculate stats
    err := filepath.Walk(stats.Dir, func(path string, file os.FileInfo, err error) error {
        if file.Name() == ".git" {
            stats.ScanNum++
            stats.gitArguments.pathArgs = ` --git-dir="` + path + `" `
            return stats.CalculateStats(path)
        }
        return nil
    })

    if err != nil {
        return err
    }

    pterm.Success.Println(fmt.Sprintf("Scanned repositories:%d  ", stats.ScanNum) +
        pterm.Cyan(fmt.Sprintf("Commits:%d  ", stats.CommitCount)) +
        pterm.LightGreen(fmt.Sprintf("Additions(+):%d  ", stats.Add)) +
        pterm.Red(fmt.Sprintf("Deletions(-):%d  ", stats.Del)) +
        pterm.Yellow(fmt.Sprintf("Total lines:%d  ", stats.Add-stats.Del)),
    )

    return nil
}

func (stats *GitStats) CalculateStats(path string) error {
    spinner := NewInfoSpinner("Calculating " + path)
    args := stats.gitArguments

    cmd :=
        `git` + args.pathArgs + ` log ` + args.authorArgs + args.sinceArgs + ` --pretty=tformat: --numstat | 
        awk '{ add += $1; del += $2; } 
        function defaultVal(var){ return var == "" ? 0 : var } 
        END { printf "%s,%s,", defaultVal(add), defaultVal(del) }' && 
        git` + args.pathArgs + ` rev-list ` + args.authorArgs + args.sinceArgs + ` --all --count `

    out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
    if err != nil {
        pterm.Error.Printf("%s: %s\n", err, out)
        return err
    }
    str := strings.Split(string(out), ",")
    add, _ := strconv.Atoi(strings.TrimSpace(str[0]))
    stats.Add += add
    del, _ := strconv.Atoi(strings.TrimSpace(str[1]))
    stats.Del += del
    commitCount, _ := strconv.Atoi(strings.TrimSpace(str[2]))
    stats.CommitCount += commitCount

    spinner.Info(fmt.Sprintf("Path:%s  ", path) +
        pterm.Cyan(fmt.Sprintf("Commits:%d  ", commitCount)) +
        pterm.LightGreen(fmt.Sprintf("Additions(+):%d  ", add)) +
        pterm.Red(fmt.Sprintf("Deletions(-):%d  ", del)) +
        pterm.Yellow(fmt.Sprintf("Total lines:%d  ", add-del)),
    )

    return nil
}
