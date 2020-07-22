package github

import "time"

// github 패키지는 GitHub 이슈 트래커에 대한 Go API를 제공한다.
// https://developer.github.com/v3/search/#search-issues 참조

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct{
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct{
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"Created_at"`
	Body string // 마크다운 포멧
}

type User struct{
	Login string
	HTMLURL string `json:"html_url"`
}