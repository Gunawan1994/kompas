package log

// ILogResource interface
type ILogResource interface {
	GetAll(search string, page int, limit int) ([]LogModel, error)
}

// Log class
type Log struct {
	log ILogResource
}

// New will create the audit_log object
func New(l ILogResource) *Log {
	return &Log{
		log: l,
	}
}

// GetAllLog will return all audit_log
func (l *Log) GetAllLog(search string, page int, limit int) ([]LogModel, error) {

	return l.log.GetAll(search, page, limit)
}
