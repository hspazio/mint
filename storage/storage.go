package storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

// Note is initialized using Name and Path of the file
type Note struct {
	Name string
	Path string
}

// Store contains the applications's data
type Store struct {
	Dir    string
	Editor string
}

// NewStore creates a Store
func NewStore(editor string) (*Store, error) {
	store := &Store{
		Dir:    filepath.Join(rootdir(), ".mint"),
		Editor: editor,
	}
	if err := store.setup(); err != nil {
		return nil, err
	}
	return store, nil
}

// Notes lists all available notes in the store
func (s Store) Notes() ([]Note, error) {
	var notes []Note
	files, err := ioutil.ReadDir(s.Dir)
	if err != nil {
		return notes, err
	}
	for _, f := range files {
		note := s.noteFromFile(f)
		notes = append(notes, note)
	}
	return notes, nil
}

// EditNote will open a Note with the default $EDITOR
func (s Store) EditNote(note Note) error {
	file := filepath.Join(s.Dir, note.Path)
	edit := exec.Command(s.Editor, file)
	edit.Stdin = os.Stdin
	edit.Stdout = os.Stdout
	edit.Stderr = os.Stderr

	if err := edit.Start(); err != nil {
		return err
	}
	if err := edit.Wait(); err != nil {
		return err
	}
	return nil
}

// RemoveNote will take in input a name and will remove the Note with the given name
func (s Store) RemoveNote(name string) error {
	note := s.NoteFromName(name)
	file := filepath.Join(s.Dir, note.Path)

	err := os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}

// WriteNote will save the content to a Note
func (s Store) WriteNote(note Note, b []byte) error {
	file := filepath.Join(s.Dir, note.Path)
	return ioutil.WriteFile(file, b, os.ModePerm)
}

// NoteFromName creates a Note given a name. It does not guarantee that the note already exists
func (s Store) NoteFromName(name string) Note {
	path := fmt.Sprintf("%s.%s", name, "md")
	return Note{Name: name, Path: path}
}

func (s Store) noteFromFile(f os.FileInfo) Note {
	name := strings.TrimSuffix(f.Name(), ".md")
	return Note{Name: name, Path: f.Name()}
}

func (s Store) setup() error {
	return os.MkdirAll(s.Dir, os.ModePerm)
}

func rootdir() string {
	usr, err := user.Current()
	if err == nil {
		return usr.HomeDir
	}
	return "."
}
