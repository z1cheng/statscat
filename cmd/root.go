/*
Copyright © 2022 Chen Chen imchench@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
    "github.com/spf13/cobra"
    "github.com/z1cheng/statscat/stats"
)

var rootDir string
var author string
var since string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
    Use:   "statscat [-d dir] [-a author] [--since since] ",
    Short: "Stats Cat🐈 is a CLI tool to get statistics of your all git repositories",
    Example: `
    statscat  # get the statistics of all repositories in current directory
    statscat -d /directory -a author --since 1.week  # get the statistics of all repositories under /directory, author is author name, since is from 1 week ago`,
    DisableFlagsInUseLine: true,
    Run: func(cmd *cobra.Command, args []string) {
        stats := stats.NewGitStats(author, since, rootDir)
        stats.GetStats()
    },
}

func init() {
    RootCmd.Flags().StringVarP(&rootDir, "dir", "d", ".", "directory to be calculated, statscat will search recursively, default is current directory")
    RootCmd.Flags().StringVarP(&author, "author", "a", "", "author name to be calculated, default is all authors")
    RootCmd.Flags().StringVar(&since, "since", "", "show stats more recent than a specific date")
}
