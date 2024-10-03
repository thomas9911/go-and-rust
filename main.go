package main

// NOTE: There should be NO space between the comments and the `import "C"` line.

/*
#cgo LDFLAGS: -L${SRCDIR}\out -l go_and_rust.dll
#include "./out/go_and_rust.h"
#include <stdio.h>
#include <windows.h>
*/

/*
#cgo LDFLAGS: ./out/libgo_and_rust.a -ldl
#include "./out/go_and_rust.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unsafe"
)

type Song struct {
	Title       string
	Artist      string
	ReleaseYear int32
}

func SongintoGo(song C.Song_t) Song {
	return Song{
		Title:       C.GoString(song.title),
		Artist:      C.GoString(song.artist),
		ReleaseYear: int32(song.release_year),
	}
}

func SongResultIntoGo(songResult C.GoResult_Song_SongError_t) (Song, error) {
	if songResult.err._0 != false {
		var error error
		switch C.SongError_t(songResult.err._1) {
		case C.SONG_ERROR_INVALID_TITLE:
			error = errors.New("Invalid title")
			break
		case C.SONG_ERROR_INVALID_ARTIST:
			error = errors.New("Invalid artist")
			break
		case C.SONG_ERROR_INVALID_RELEASE_YEAR:
			error = errors.New("Invalid release year")
			break
		default:
			log.Fatal("Unhandled error code")
		}

		return Song{}, error
	}

	// assume if not error then song is valid
	return SongintoGo(songResult.val._1), nil
}

func (song Song) intoC() (song_c C.Song_t) {
	song_c.title = C.CString(song.Title)
	song_c.artist = C.CString(song.Artist)
	song_c.release_year = C.uint(song.ReleaseYear)

	return song_c
}

func main() {
	argsWithoutProg := os.Args[1:]
	word := strings.Join(argsWithoutProg, " ")

	str1 := C.CString(word)
	defer C.free(unsafe.Pointer(str1))

	out := C.count_numbers(str1)

	title := C.CString("Hello")
	defer C.free(unsafe.Pointer(title))

	artist := C.CString("Adele")
	defer C.free(unsafe.Pointer(artist))

	songC := C.new_song(title, artist, 2015)
	song := SongintoGo(songC)

	fmt.Printf("%v\n", song)

	song2, err := SongResultIntoGo(C.try_new_song(title, artist, 2015))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("=> %v\n", song2)

	fmt.Printf("Counted %d numbers\n", out)
}
