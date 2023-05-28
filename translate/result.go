package translate

type Result struct {
	ReturnPhrase  []string      `json:"returnPhrase"`
	Query         string        `json:"query"`
	ErrorCode     string        `json:"errorCode"`
	L             string        `json:"l"`
	TSpeakURL     string        `json:"tSpeakUrl"`
	Web           []Web         `json:"web"`
	RequestID     string        `json:"requestId"`
	Translation   []string      `json:"translation"`
	MTerminalDict MTerminalDict `json:"mTerminalDict"`
	Dict          Dict          `json:"dict"`
	Webdict       Webdict       `json:"webdict"`
	Basic         Basic         `json:"basic"`
	IsWord        bool          `json:"isWord"`
	SpeakURL      string        `json:"speakUrl"`
}
type Web struct {
	Value []string `json:"value"`
	Key   string   `json:"key"`
}
type MTerminalDict struct {
	URL string `json:"url"`
}
type Dict struct {
	URL string `json:"url"`
}
type Webdict struct {
	URL string `json:"url"`
}
type Basic struct {
	UkSpeech string   `json:"uk-speech"`
	Explains []string `json:"explains"`
	UsSpeech string   `json:"us-speech"`
}
