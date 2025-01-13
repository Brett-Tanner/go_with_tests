package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

const DbFileName = "game.db.json"

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(DbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", DbFileName, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		return nil, nil, fmt.Errorf("problem creating player store %v", err)
	}

	return store, closeFunc, nil
}

func initializePlayerDBFile(database *os.File) error {
	database.Seek(0, 0)

	info, err := database.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", database.Name(), err)
	}

	if info.Size() == 0 {
		database.Write([]byte("[]"))
		database.Seek(0, 0)
	}

	return nil
}

func NewFileSystemPlayerStore(database *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(database)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player DB file %s, %v", database.Name(), err)
	}

	league, err := NewLeague(database)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from %s, %v", database.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{database}), league: league,
	}, nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Score > f.league[j].Score
	})

	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Score
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Score++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)
}
