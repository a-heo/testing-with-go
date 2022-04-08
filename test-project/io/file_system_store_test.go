package main

import (
	"testing"
	"io"
	"io/ioutil"
	"os"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		//strings new reader returns us a reader that func filesystemplayerstore uses to read data
		// database := strings.NewReader(`[
		// 	{"Name": "Cleo", "Wins": 10},
		// 	{"Name": "Chris", "Wins": 33}]`)

		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)

		//read again (readseeker interface allows for this)
		got = store.GetLeague()
		assertLeague(t, got, want)

	})

	t.Run("get player score", func(t *testing.T) {
		//strings.reader cannot implement readwriteseeker
		// database := strings.NewReader(`[
		// 	{"Name": "Cleo", "Wins": 10},
		// 	{"Name": "Chris", "Wins": 33}]`)
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)

	})

	t.Run("store wins for existing player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
			defer cleanDatabase()

			store := FileSystemPlayerStore{database}

			store.RecordWin("Pepper")
			
			got := store.GetPlayerScore("Pepper")
			want := 1
			assertScoreEquals(t, got, want)
	})
}

//create temp file to use for testing
func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	//removes temp file once test is finished (prevent leaks)
	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}