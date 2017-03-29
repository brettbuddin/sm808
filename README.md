# sm808

A drum machine inspired by the Roland TR-808.

Example tablature files can be found in the [examples directory](examples).

## Requirements

- [Go](https://golang.org)
- [PortAudio](https://portaudio.org) `brew install portaudio`

## Installation

    go get github.com/brettbuddin/sm808

## Manual

Usage:

    sm808 [-tempo bpm] [-steps steps] [-no-print] tablature.txt

### Tempo and Steps per Beat

Supplying the `-tempo` flag will establish the tempo for the sequence. 

By default, each step acts as a quarter note (1 step per beat), but this can be overridden by supplying the `-steps`
flag. For example: supplying a value of 3 will play 3 steps per beat.

### Supported Voices

- `closedhihat`
- `openhihat`
- `bassdrum`
- `snaredrum`

### Tablature Files

A tablature file is the means of programming sm808. These files consist of a list of voices, one per line that are
loaded and played simultaneously to create rythms. Here's an example file:

    closedhat [x _ x X]
    snaredrum [_ x X _ _ _ x _]
    bassdrum  [x _ _ _ x _ _ _]

Each line consists of a voice name (e.g. "hihat") and a pattern to play for that voice. Patterns can be any length
and consist of three possible values separated by spaces. The three valid options are:

- `x` (lowercase X): low-velocity strike
- `X` (uppercase X): high-velocity strike
- `_` (underscore):  rest

Note that the length of the patterns do not need to be equal. Providing patterns of different lengths will provide
you with more interesting interleaving of the voices.

The output printed while the program is running is a limited visual representation of what the composition of each
step is in the timeline. Each voice that's actively playing something in that step will be represented. Steps are
separated by `-----`. This output can be disabled by providing the `-no-print` flag.

### Example

	$ sm808 -help
	Usage of sm808:
  	-tempo int
        	beats per minute (default 60)
  	-no-print
        	disable step printing
  	-steps int
        	steps per beat (default 2)
    $ sm808 -tempo 120 -steps 4 examples/metal.txt
	...
