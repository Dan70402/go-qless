package qless

//easyjson:json
type jobDataArray []jobData

//easyjson:json
type jobData struct {
	JID          string      `json:"jid"`
	Class        string      `json:"klass"`
	State        string      `json:"state"`
	Queue        string      `json:"queue"`
	Worker       string      `json:"worker"`
	Tracked      bool        `json:"tracked"`
	Priority     int         `json:"priority"`
	Expires      int64       `json:"expires"`
	Retries      int         `json:"retries"`
	Remaining    int         `json:"remaining"`
	Data         string      `json:"data"`
	Tags         StringSlice `json:"tags"`
	History      []History   `json:"history"`
	Failure      *Failure    `json:"failure,omitempty"`
	Dependents   StringSlice `json:"dependents"`
	Dependencies StringSlice `json:"dependencies"`
}

//easyjson:json
type History struct {
	When   int64  `json:"when"`
	Queue  string `json:"q"`
	What   string `json:"what"`
	Worker string `json:"worker"`
}

//easyjson:json
type QueueInfo struct {
	Name      string `json:"name"`
	Paused    bool   `json:"paused"`
	Waiting   int    `json:"waiting"`
	Running   int    `json:"running"`
	Stalled   int    `json:"stalled"`
	Scheduled int    `json:"scheduled"`
	Recurring int    `json:"recurring"`
	Depends   int    `json:"depends"`
}

//easyjson:json
type Failure struct {
	Group   string `json:"group"`
	Message string `json:"message"`
	When    int64  `json:"when"`
	Worker  string `json:"worker"`
}

//easyjson:json
type StatData struct {
	Count     int64   `json:"count"`
	Histogram []int64 `json:"histogram"`
	Mean      float64 `json:"mean"`
	Std       float64 `json:"std"`
}

//easyjson:json
type QueueStatistics struct {
	Failed   int64    `json:"failed"`
	Failures int64    `json:"failures"`
	Retries  int64    `json:"retries"`
	Run      StatData `json:"run"`
	Wait     StatData `json:"wait"`
}
