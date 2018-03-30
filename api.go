package main

type Result struct {
	status string
}

type Equalizer struct {
	Bass              int
	Treble            int
	Loudness          bool
	SpeechEnhancement bool
	NightMode         bool
}

type Track struct {
	Artist              string
	Title               string
	Album               string
	AlbumArtUri         string
	Duration            int
	Uri                 string
	Type                string
	StationName         string
	AbsoluteAlbumArtUri string
}

type PlayMode struct {
	Repeat    string
	Shuffle   bool
	Crossfade bool
}

type Sub struct {
	Gain      int
	Crossover int
	Polarity  int
	Enabled   bool
}

type State struct {
	Volume               int
	Mute                 bool
	Equalizer            Equalizer
	CurrentTrack         Track
	NextTrack            Track
	TrackNo              int
	ElapsedTime          int
	ElapsedTimeFormatted string
	PlaybackState        string
	PlayMode             PlayMode
	Sub                  Sub
}

type GroupState struct {
	Volume int
	Mute   bool
}

type Member struct {
	Uuid        string
	State       State
	RoomName    string
	Coordinator string
	GroupState  GroupState
}

type Zones []Zone

type Zone struct {
	Uuid        string
	Coordinator Member
	Members     []Member
}
