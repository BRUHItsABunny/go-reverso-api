package go_reverso_api

import gokhttp "github.com/BRUHItsABunny/gOkHttp"

type ReversoClient struct {
	Client *gokhttp.HttpClient
}

type TranslateOptions struct {
	Origin            string `json:"origin"`
	SentenceSplitter  bool   `json:"sentenceSplitter"`
	ContextResults    bool   `json:"contextResults,omitempty"`
	LanguageDetection bool   `json:"languageDetection,omitempty"`
}

type TranslateRequest struct {
	Input   string           `json:"input"`
	From    string           `json:"from"`
	To      string           `json:"to"`
	Format  string           `json:"format"`
	Options TranslateOptions `json:"options"`
}

type Language struct {
	Code   string `json:"code"`
	Alpha3 string `json:"alpha3"`
}

type SynonymsResponse struct {
	ID                  int64           `json:"id"`
	Search              string          `json:"search"`
	Language            string          `json:"language"`
	Fuzzyhash           string          `json:"fuzzyhash"`
	Input               string          `json:"input"`
	Pos                 Pos             `json:"pos"`
	SearchType          string          `json:"searchType"`
	Allowredirect       bool            `json:"allowredirect"`
	Rude                bool            `json:"rude"`
	Colloquial          bool            `json:"colloquial"`
	PotentialRude       int64           `json:"potentialRude"`
	PotentialColloquial int64           `json:"potentialColloquial"`
	Groupable           bool            `json:"groupable"`
	ResultsCount        int64           `json:"resultsCount"`
	Results             []SynonymResult `json:"results"`
	Suggestions         []interface{}   `json:"suggestions"`
	Antonyms            []interface{}   `json:"antonyms"`
	Related             []Related       `json:"related"`
	Stopwatch           Stopwatch       `json:"stopwatch"`
}

type Pos struct {
	Mask int64    `json:"mask"`
	Desc []string `json:"desc"`
}

type Related struct {
	ID         int64  `json:"id"`
	Word       string `json:"word"`
	Pos        Pos    `json:"pos"`
	Language   string `json:"language"`
	Rude       bool   `json:"rude"`
	Colloquial bool   `json:"colloquial"`
}

type SynonymResult struct {
	Pos               Pos           `json:"pos"`
	Weight            int64         `json:"weight"`
	Nrows             int64         `json:"nrows"`
	Relevance         int64         `json:"relevance"`
	RudeResults       int64         `json:"rudeResults"`
	ColloquialResults int64         `json:"colloquialResults"`
	Merged            interface{}   `json:"merged"`
	Relevantitems     int64         `json:"relevantitems"`
	Cluster           []Cluster     `json:"cluster"`
	Examples          []Example     `json:"examples"`
	Antonyms          []interface{} `json:"antonyms"`
}

type Cluster struct {
	ID           int64   `json:"id"`
	Word         string  `json:"word"`
	Language     string  `json:"language"`
	Cluster      int64   `json:"cluster"`
	Weight       int64   `json:"weight"`
	Nrows        int64   `json:"nrows"`
	Isentry      bool    `json:"isentry"`
	Pos          Pos     `json:"pos"`
	Rude         bool    `json:"rude"`
	Colloquial   bool    `json:"colloquial"`
	Relevance    float64 `json:"relevance"`
	MostRelevant bool    `json:"mostRelevant"`
}

type Example struct {
	ID      int64  `json:"id"`
	Cluster int64  `json:"cluster"`
	Example string `json:"example"`
	Pos     Pos    `json:"pos"`
}

type Stopwatch struct {
	Start   float64 `json:"start"`
	Ended   float64 `json:"ended"`
	Elapsed float64 `json:"elapsed"`
}

type TranslateResponse struct {
	ID                string            `json:"id"`
	From              string            `json:"from"`
	To                string            `json:"to"`
	Input             []string          `json:"input"`
	CorrectedText     interface{}       `json:"correctedText"`
	Translation       []string          `json:"translation"`
	Engines           []string          `json:"engines"`
	LanguageDetection LanguageDetection `json:"languageDetection"`
	ContextResults    ContextResults    `json:"contextResults"`
	Truncated         bool              `json:"truncated"`
	TimeTaken         int64             `json:"timeTaken"`
}

type ContextResults struct {
	RudeWords             bool              `json:"rudeWords"`
	Colloquialisms        bool              `json:"colloquialisms"`
	RiskyWords            bool              `json:"riskyWords"`
	Results               []TranslateResult `json:"results"`
	TotalContextCallsMade int64             `json:"totalContextCallsMade"`
	TimeTakenContext      int64             `json:"timeTakenContext"`
}

type TranslateResult struct {
	Translation    string   `json:"translation"`
	SourceExamples []string `json:"sourceExamples"`
	TargetExamples []string `json:"targetExamples"`
	Rude           bool     `json:"rude"`
	Colloquial     bool     `json:"colloquial"`
	PartOfSpeech   string   `json:"partOfSpeech"`
}

type LanguageDetection struct {
	DetectedLanguage                string `json:"detectedLanguage"`
	IsDirectionChanged              bool   `json:"isDirectionChanged"`
	OriginalDirection               string `json:"originalDirection"`
	OriginalDirectionContextMatches int64  `json:"originalDirectionContextMatches"`
	ChangedDirectionContextMatches  int64  `json:"changedDirectionContextMatches"`
}

type SuggestResponse struct {
	Suggestions []FuzzySuggestion `json:"suggestions"`
	Fuzzy1      []FuzzySuggestion `json:"fuzzy1"`
	Fuzzy2      []FuzzySuggestion `json:"fuzzy2"`
	Request     SuggestionRequest `json:"request"`
	TimeMS      int64             `json:"time_ms"`
}

type FuzzySuggestion struct {
	Lang       string `json:"lang"`
	Suggestion string `json:"suggestion"`
	Weight     int64  `json:"weight"`
	IsFromDict bool   `json:"isFromDict"`
}

type SuggestionRequest struct {
	Search     string `json:"search"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
	MaxResults int64  `json:"max_results"`
	Mode       int64  `json:"mode"`
}

type SearchResponse struct {
	List                     []SearchResult        `json:"list"`
	Nrows                    int64                 `json:"nrows"`
	NrowsExact               int64                 `json:"nrows_exact"`
	Pagesize                 int64                 `json:"pagesize"`
	Npages                   int64                 `json:"npages"`
	Page                     int64                 `json:"page"`
	RemovedSuperstrings      []string              `json:"removed_superstrings"`
	DictionaryEntryList      []DictionaryEntryList `json:"dictionary_entry_list"`
	DictionaryOtherFrequency int64                 `json:"dictionary_other_frequency"`
	DictionaryNrows          int64                 `json:"dictionary_nrows"`
	TimeMS                   int64                 `json:"time_ms"`
	Request                  SearchRequest         `json:"request"`
	Suggestions              []FuzzySuggestion     `json:"suggestions"`
	DymCase                  int64                 `json:"dym_case"`
	DymList                  []interface{}         `json:"dym_list"`
	DymApplied               interface{}           `json:"dym_applied"`
	DymNonadaptedSearch      interface{}           `json:"dym_nonadapted_search"`
	DymPairApplied           interface{}           `json:"dym_pair_applied"`
	DymNonadaptedSearchPair  interface{}           `json:"dym_nonadapted_search_pair"`
	DymPair                  interface{}           `json:"dym_pair"`
	ExtractedPhrases         []interface{}         `json:"extracted_phrases"`
}

type DictionaryEntryList struct {
	Frequency        int64                 `json:"frequency"`
	Term             string                `json:"term"`
	IsFromDict       bool                  `json:"isFromDict"`
	IsPrecomputed    bool                  `json:"isPrecomputed"`
	Stags            []string              `json:"stags"`
	Pos              *string               `json:"pos"`
	Sourcepos        []string              `json:"sourcepos"`
	Variant          interface{}           `json:"variant"`
	Domain           interface{}           `json:"domain"`
	Definition       *string               `json:"definition"`
	Vowels2          interface{}           `json:"vowels2"`
	Transliteration2 interface{}           `json:"transliteration2"`
	AlignFreq        int64                 `json:"alignFreq"`
	ReverseValidated bool                  `json:"reverseValidated"`
	PosGroup         int64                 `json:"pos_group"`
	IsTranslation    bool                  `json:"isTranslation"`
	IsFromLOCD       bool                  `json:"isFromLOCD"`
	InflectedForms   []DictionaryEntryList `json:"inflectedForms"`
}

type SearchResult struct {
	SText string `json:"s_text"`
	TText string `json:"t_text"`
	Ref   string `json:"ref"`
	Cname string `json:"cname"`
	URL   string `json:"url"`
	Ctags string `json:"ctags"`
	Pba   bool   `json:"pba"`
}

type SearchRequest struct {
	SourceText     string `json:"source_text"`
	TargetText     string `json:"target_text"`
	SourceLang     string `json:"source_lang"`
	TargetLang     string `json:"target_lang"`
	Npage          int64  `json:"npage"`
	Corpus         string `json:"corpus"`
	Nrows          int64  `json:"nrows"`
	Adapted        bool   `json:"adapted"`
	NonadaptedText string `json:"nonadapted_text"`
	RudeWords      bool   `json:"rude_words"`
	Colloquialisms bool   `json:"colloquialisms"`
	RiskyWords     bool   `json:"risky_words"`
	Mode           int64  `json:"mode"`
	ExprSug        int64  `json:"expr_sug"`
	DymApply       bool   `json:"dym_apply"`
	PosReorder     int64  `json:"pos_reorder"`
	Device         int64  `json:"device"`
	SplitLong      bool   `json:"split_long"`
	HasLocd        bool   `json:"has_locd"`
	SourcePos      string `json:"source_pos"`
}
