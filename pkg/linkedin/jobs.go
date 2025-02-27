package linkedin

type Job struct {
	ID           int
	Title        string
	Description  string
	Requirements string
	Location     string
}

func NewJob(id int, title, des, req, loc string) *Job {
	return &Job{
		ID:           id,
		Title:        title,
		Description:  des,
		Requirements: req,
		Location:     loc,
	}

}
