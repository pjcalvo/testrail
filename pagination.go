package testrail

// PaginatedSuites represents a Test Suites response
// from the TestRail API (suites w/ pagination)
type PaginatedResults struct {
	Offset int             `json:"offset"`
	Limit  int             `json:"limit"`
	Size   int             `json:"size"`
	Suites []Suite         `json:"suites"`
	Links  PaginationLinks `json:"_links"`
}

type PaginationLinks struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
