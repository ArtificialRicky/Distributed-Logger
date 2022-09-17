package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"

	grep "cs425.com/mp1/grepService"
)

func ShowResult(replies []grep.ReplyFormat, servers []string) {
	total_match := 0
	for i, reply := range replies {
		for _, match := range reply.MatchStats {
			fmt.Printf("%s %s:%d\n", servers[i], match.File, match.N_line)
			total_match += match.N_line
		}
	}
	fmt.Println("Total Match:", total_match)
}

func main() {
	flags := flag.String("flags", "c", "Flags for grep. --flags=Ec for regular expression.")
	pattern := flag.String("pattern", "", "pattern to search.")
	file_prefix := flag.String("prefix", "vm", "prefix of files to search.")
	n_machine := flag.Int("n", 10, "number of machines")
	logging := flag.Bool("log", false, "Logging errors")
	flag.Parse()

	port := 1234

	var replies []grep.ReplyFormat
	var replied_servers []string
	for i := 1; i <= *n_machine; i++ {
		host := fmt.Sprintf("fa22-cs425-80%02d.cs.illinois.edu", i)
		host_port := fmt.Sprintf("%s:%d", host, port)
		client, err := rpc.Dial("tcp", host_port)

		if err != nil {
			if *logging {
				log.Print("Dialing ", host_port)
				log.Println(err)
			}
			continue
		}

		args := &grep.GrepFormat{Flags: *flags, Pattern: *pattern, File: *file_prefix + "*"}

		var reply grep.ReplyFormat
		err = client.Call("GrepService.Grep", args, &reply)
		if err != nil {
			if *logging {
				log.Print("Calling GrepService.Grep", host_port)
				log.Println(err)
			}
			continue
		}

		replied_servers = append(replied_servers, host)
		replies = append(replies, reply)
	}
	ShowResult(replies, replied_servers)

}
