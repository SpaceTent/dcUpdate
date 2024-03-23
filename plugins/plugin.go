package plugins

var Plugin = make(map[string]any)

func Truncate(s string) string {
	if len(s) > 20 {
		return s[:20] + "..."
	}
	return s
}

func init() {

	Plugin["Truncate"] = func(in string) string { return Truncate(in) }
}
