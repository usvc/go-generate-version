// +build ignore

package main

import (
	"bytes"
	"html/template"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	appVersion, appCommit := getRepoVersion()
	if versionFile, err := os.Create("./version.go"); err != nil {
		panic(err)
	} else {
		defer versionFile.Close()
		versionTemplate.Execute(versionFile, struct {
			AppVersion Version
			AppCommit  Commit
			Timestamp  string
		}{
			AppVersion: appVersion,
			AppCommit:  appCommit,
			Timestamp:  time.Now().Format("2006-01-02T15:04:05-0700"),
		})
	}
}

// Version contains the semver version
type Version string

// Commit contains the commit SHA hash
type Commit string

func getRepoVersion() (Version, Commit) {
	_, err := exec.LookPath("git")
	if err != nil {
		panic(err)
	}

	var versionOutput bytes.Buffer
	getVersion := exec.Command("git", "describe", "--tags", "--abbrev=0")
	getVersion.Stdout = &versionOutput
	getVersion.Stderr = &versionOutput
	getVersion.Run()

	var commitOutput bytes.Buffer
	getCommit := exec.Command("git", "log", "-n", "1", "--format='%H'")
	getCommit.Stdout = &commitOutput
	getCommit.Stderr = &commitOutput
	getCommit.Run()

	version := strings.Trim(versionOutput.String(), " \n")
	commit := strings.Trim(commitOutput.String(), " -'\n")[:7]
	return Version(version), Commit(commit)
}

var versionTemplate = template.Must(template.New("test").Parse(`
// GENERATED FILE - DO NOT MODIFY
//
// GENERATED BY GO:GENERATE AT {{.Timestamp}}
//
// FILE GENERATED USING ~/generators/versioning/main.go

package main

const Version = "{{.AppVersion}}"

const Commit = "{{.AppCommit}}"

`))