package main

import (
	"os"
)

// Usage example:
//     export EGET_GITHUB_PROXY="example.com/artifactory/github"
//     export EGET_GITHUB_API_PROXY="example.com/artifactory/github-api"

var GithubProxy = "github.com"
var GithubApiProxy = "api.github.com"

func init() {
	if value, ok := os.LookupEnv("EGET_GITHUB_PROXY"); ok {
		GithubProxy = value
	}
	if value, ok := os.LookupEnv("EGET_GITHUB_API_PROXY"); ok {
		GithubApiProxy = value
	}
}
