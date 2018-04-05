package filesystem

var List = make(map[string]FileSystem)

type FileSystem interface {
	// Create returns a command followed by its arguments, that will
	// create path as a directory.
	Create(path string) []string

	// Clone returns a command followed by its arguments, that will
	// clone src to dst.
	Clone(src, dst string) []string

	// Remove removes path and its children.
	Remove(path string) error
}