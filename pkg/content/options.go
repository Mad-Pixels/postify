package content

const metafileName = "meta.json"
const urlPrefix = "content/"

// fType is an enumeration of supported content format.
type fType int

const (
	HTML fType = iota
	Telegram
)

// Option ...
type Option func(Content)
