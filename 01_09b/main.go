package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`

	albumNum int
	songNum  int
}

type PlaylistHeap []Song

func (p PlaylistHeap) Len() int {
	return len(p)
}

func (p PlaylistHeap) Less(i, j int) bool {
	return p[i].PlayCount > p[j].PlayCount
}

func (p PlaylistHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PlaylistHeap) Push(x any) {
	*p = append(*p, x.(Song))
}

func (p *PlaylistHeap) Pop() any {
	original := *p
	n := len(original) - 1
	songPopped := original[n]
	*p = original[0:n]
	return songPopped
}

// makePlaylist makes the merged sorted list of songs
func makePlaylist(albums [][]Song) []Song {
	//var playlist []Song
	//for _, album := range albums {
	//	for _, song := range album {
	//		playlist = append(playlist, song)
	//	}
	//}
	//
	//sort.SliceStable(playlist, func(i, j int) bool { return playlist[i].PlayCount > playlist[j].PlayCount })
	//return playlist

	playlistHeap := &PlaylistHeap{}
	heap.Init(playlistHeap)
	var playlist []Song
	for i, album := range albums {
		song := album[0]
		song.albumNum = i
		song.songNum = 0
		heap.Push(playlistHeap, song)
	}
	for playlistHeap.Len() != 0 {
		song := heap.Pop(playlistHeap).(Song)
		playlist = append(playlist, song)
		if song.songNum < len(albums[song.albumNum])-1 {
			nextSong := albums[song.albumNum][song.songNum+1]
			nextSong.songNum = song.songNum + 1
			nextSong.albumNum = song.albumNum
			heap.Push(playlistHeap, nextSong)
		}
	}

	return playlist
}

func main() {
	albums := importData()
	printTable(makePlaylist(albums))
}

// printTable prints merged playlist as a table
func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()

}

// importData reads the input data from file and creates the friends map
func importData() [][]Song {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data [][]Song
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
