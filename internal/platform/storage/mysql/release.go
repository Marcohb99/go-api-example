package mysql

const (
	sqlReleaseTable = "releases"
)

type sqlRelease struct {
	ID          string `db:"uuid"`
	Title       string `db:"title"`
	Released    string `db:"released"`
	ResourceUrl string `db:"resource_url"`
	Uri         string `db:"uri"`
	Year        string `db:"year"`
}
