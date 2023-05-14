package input

import (
	log "github.com/pterm/pterm"
	"path/filepath"
)

func GetImageFilenames(args []string) []string {
	out := make([]string, 0)
	for _, arg := range args {
		matches, err := filepath.Glob(arg)
		if err != nil {
			log.Error.Printf("Could not parse argument %s: %v\n", arg, err)
			continue
		}
		out = append(out, matches...)
	}

	if len(out) == 0 {
		log.Warning.Printf("No files match the given name, folder or wildcards (%s).\n", args)
	}

	return out
}
