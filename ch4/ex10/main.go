// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/daviddengcn/go-colortext"

	"github.com/chris-gilmore/gopl/ch4/ex10/github"
)

// $ ./main repo:golang/go is:open json decoder
func main() {
	// Search issues and sort by created date in descending order
	result, err := github.SearchIssues(os.Args[1:], "created", "desc")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	aYearAgo := time.Now().UTC().AddDate(-1, 0, 0)
	aMonthAgo := time.Now().UTC().AddDate(0, -1, 0)
	for _, item := range result.Items {
		// Alter text color based on Age category
		switch {
		case item.CreatedAt.Before(aYearAgo):
			ct.Foreground(ct.White, true)
		case item.CreatedAt.Before(aMonthAgo):
			ct.Foreground(ct.Cyan, true)
		default:
			ct.Foreground(ct.Blue, true)
		}

		fmt.Printf("%s #%-5d %9.9s %.55s",
			item.CreatedAt, item.Number, item.User.Login, item.Title)

		// Reset text color back to normal
		ct.ResetColor()

		fmt.Println()
	}
}
