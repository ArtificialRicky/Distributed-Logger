package grep

import (
	"fmt"
	"os/exec"
	"strings"
)

type GrepFormat struct {
	Flags, Pattern, File string
}

type FileMatchLines struct {
	File   string
	N_line int
}

type ReplyFormat struct {
	MatchStats []FileMatchLines
}

type GrepService int

func (grep_service *GrepService) Grep(grep_format *GrepFormat, reply *ReplyFormat) error {
	fmt.Println("pattern =", grep_format.Pattern)
	fmt.Println("file =", grep_format.File)

	// get all files that contain macthing lines
	// bash neccessary here as wildcard matching for files are implemented in bash env
	// grep -H show result with filename
	// grep -c show # of matching lines
	cmd := exec.Command("bash", "-c",
		"grep"+" -H"+grep_format.Flags+" \""+grep_format.Pattern+"\" "+grep_format.File)
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("CMD %s\n", cmd.String())
		fmt.Println(string(stdout))
		fmt.Println(err.Error())
		return err
	}

	// Fidls split on white space
	for _, match := range strings.Fields(string(stdout)) {
		var match_lines FileMatchLines
		// format "filename:line_count" -> "filename line_count"
		// Replace : with space because Sscanf matches %s greedily and consumes :
		match = strings.Replace(match, ":", " ", 1)
		_, err := fmt.Sscanf(match, "%s %d", &match_lines.File, &match_lines.N_line)
		if err != nil {
			fmt.Println(err.Error())
		}
		reply.MatchStats = append(reply.MatchStats, match_lines)
	}

	return nil
}
