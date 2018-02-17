package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/hspazio/mint/storage"
	"github.com/spf13/cobra"
)

var name string

func init() {
	rootCmd.AddCommand(noteCmd)
}

var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Create or edit a note",
	Long:  "Create or edit a note, create a new one if it doesn't exist",
	Run: func(cmd *cobra.Command, args []string) {
		var note storage.Note
		var err error

		if len(args) < 1 {
			n, err := noteFromList()
			if err != nil {
				exit("error while listing notes", err)
			}
			note = *n
		} else {
			note = store.NoteFromName(args[0])
		}

		stdinInfo, err := os.Stdin.Stat()
		if err != nil {
			exit("could not read stats from Stdin", err)
		}
		if stdinInfo.Size() > 0 {
			content, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				exit("could not read from Stdin", err)
			}
			err = store.WriteNote(note, content)
			if err != nil {
				exit("could not save file", err)
			}
		} else {
			fmt.Println(store.Dir)
			if err := store.EditNote(note); err != nil {
				exit("could not start editor", err)
			}
		}
	},
}

func noteFromList() (*storage.Note, error) {
	notes, err := store.Notes()
	if err != nil {
		return nil, err
	}
	printList(notes)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Which note to edit?: ")
	scanner.Scan()
	if scanner.Err() != nil {
		return nil, fmt.Errorf("could not read from Stdin: %v", scanner.Err())
	}
	index, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, fmt.Errorf("answer is not a number")
	}
	if index > (len(notes)-1) || index < 0 {
		return nil, fmt.Errorf("number is not in the list")
	}
	return &notes[index], nil
}
