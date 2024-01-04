package mysql

const (
	sqlReleaseTable = "releases"
)

type SqlRelease struct {
	ID          string `db:"id"`
	Title       string `db:"title"`
	Released    string `db:"released"`
	ResourceUrl string `db:"resource_url"`
	Uri         string `db:"uri"`
	Year        string `db:"year"`
}
