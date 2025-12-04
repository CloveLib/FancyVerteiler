package git

type GithubEvent struct {
	HeadCommit HeadCommit `json:"head_commit"`
}

type HeadCommit struct {
	ID       string `json:"id"`
	Messagae string `json:"message"`
}
