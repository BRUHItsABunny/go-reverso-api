package api

type TranslateOptions struct {
	Origin            string `json:"origin"`
	SentenceSplitter  bool   `json:"sentenceSplitter"`
	ContextResults    bool   `json:"contextResults,omitempty"`
	LanguageDetection bool   `json:"languageDetection,omitempty"`
}

type TranslateRequestBody struct {
	Input   string            `json:"input"`
	From    string            `json:"from"`
	To      string            `json:"to"`
	Format  string            `json:"format"`
	Options *TranslateOptions `json:"options"`
}

func NewTranslateRequestBody(text string, srcLang, dstLang *Language) *TranslateRequestBody {
	result := &TranslateRequestBody{
		Input:  text,
		From:   srcLang.Alpha3,
		To:     dstLang.Alpha3,
		Format: "text",
		Options: &TranslateOptions{
			Origin:           "reverso.app.android",
			SentenceSplitter: false,
		},
	}
	return result
}

type Language struct {
	Code   string `json:"code"`
	Alpha3 string `json:"alpha3"`
}

type ErrorResponse struct {
	Message string `json:"Message"`
}

func (err *ErrorResponse) Error() string {
	return err.Message
}

type AutoCompleteResponse []string

type ConjugatorResponse struct {
	IsTransliterated     bool                `json:"IsTransliterated"`
	IsFuzzyResult        bool                `json:"IsFuzzyResult"`
	IsApproximateResult  bool                `json:"IsApproximateResult"`
	Colloquial           int64               `json:"Colloquial"`
	IsUnknownVerb        bool                `json:"IsUnknownVerb"`
	Axis                 []*Axi              `json:"Axis"`
	Forms                []interface{}       `json:"Forms"`
	CanonicalForm        string              `json:"CanonicalForm"`
	CanonicalLink        string              `json:"CanonicalLink"`
	ConjugationModel     string              `json:"ConjugationModel"`
	Auxiliary            string              `json:"Auxiliary"`
	Auxiliaries          []string            `json:"Auxiliaries"`
	Comment              string              `json:"Comment"`
	Comments             *Comments           `json:"Comments"`
	ReflexiveForm        string              `json:"ReflexiveForm"`
	NegativeForm         string              `json:"NegativeForm"`
	OtherForms           []string            `json:"OtherForms"`
	Translations         map[string][]string `json:"Translations"`
	IsReflexive          bool                `json:"IsReflexive"`
	IsNegative           bool                `json:"IsNegative"`
	IsRare               bool                `json:"IsRare"`
	IsInflectedFormFound bool                `json:"IsInflectedFormFound"`
}

type Axi struct {
	IsUnknownVerb        bool    `json:"IsUnknownVerb"`
	Name                 string  `json:"Name"`
	Axis                 []*Axi  `json:"Axis"`
	Forms                []*Form `json:"Forms"`
	IsReflexive          bool    `json:"IsReflexive"`
	IsNegative           bool    `json:"IsNegative"`
	IsRare               bool    `json:"IsRare"`
	IsInflectedFormFound bool    `json:"IsInflectedFormFound"`
}

type Form struct {
	Index     int64   `json:"Index"`
	IsAltForm bool    `json:"IsAltForm"`
	Atoms     []*Atom `json:"Atoms"`
}

type Atom struct {
	AtomType   int64    `json:"AtomType"`
	Value      string   `json:"Value"`
	ValueParts []string `json:"ValueParts"`
	SpaceAfter bool     `json:"SpaceAfter"`
	Highlite   bool     `json:"Highlite"`
	Searched   bool     `json:"Searched"`
}

type Comments struct {
}

type Translations struct {
	Ru []string `json:"ru"`
	Fr []string `json:"fr"`
	Es []string `json:"es"`
}

type RephraseResponse struct {
	TextToRephrase    string       `json:"text_to_rephrase"`
	ProcessingTimeSEC float64      `json:"processing_time_sec"`
	Classification    string       `json:"classification"`
	NumberOfWords     int64        `json:"number_of_words"`
	Candidates        []*Candidate `json:"candidates"`
}

type Candidate struct {
	Candidate            string  `json:"candidate"`
	Classification       string  `json:"classification"`
	NumberWords          int64   `json:"number_words"`
	NumberWordsDifferent int64   `json:"number_words_different"`
	DiversityScore       float64 `json:"diversity_score"`
	Correction           bool    `json:"correction"`
}

type SynonymsResponse struct {
	ID                  int64                 `json:"id"`
	Search              string                `json:"search"`
	Language            string                `json:"language"`
	Fuzzyhash           string                `json:"fuzzyhash"`
	Input               string                `json:"input"`
	Pos                 *Pos                  `json:"pos"`
	SearchType          string                `json:"searchType"`
	Allowredirect       bool                  `json:"allowredirect"`
	Rude                bool                  `json:"rude"`
	Colloquial          bool                  `json:"colloquial"`
	PotentialRude       int64                 `json:"potentialRude"`
	PotentialColloquial int64                 `json:"potentialColloquial"`
	Groupable           bool                  `json:"groupable"`
	ResultsCount        int64                 `json:"resultsCount"`
	Results             []*Result             `json:"results"`
	Suggestions         []interface{}         `json:"suggestions"`
	Antonyms            []interface{}         `json:"antonyms"`
	Related             []interface{}         `json:"related"`
	ExtendedSuggestions []*ExtendedSuggestion `json:"extendedSuggestions"`
	Stopwatch           *Stopwatch            `json:"stopwatch"`
}

type ExtendedSuggestion struct {
	ID         int64  `json:"id"`
	Word       string `json:"word"`
	Pos        *Pos   `json:"pos"`
	Language   string `json:"language"`
	Rude       bool   `json:"rude"`
	Colloquial bool   `json:"colloquial"`
}

type Pos struct {
	Mask int64    `json:"mask"`
	Desc []string `json:"desc"`
}

type Result struct {
	Pos               *Pos          `json:"pos"`
	Weight            int64         `json:"weight"`
	Nrows             int64         `json:"nrows"`
	Relevance         int64         `json:"relevance"`
	RudeResults       int64         `json:"rudeResults"`
	ColloquialResults int64         `json:"colloquialResults"`
	Merged            interface{}   `json:"merged"`
	Relevantitems     int64         `json:"relevantitems"`
	Cluster           []*Cluster    `json:"cluster"`
	Examples          []*Example    `json:"examples"`
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
	Pos          *Pos    `json:"pos"`
	Rude         bool    `json:"rude"`
	Colloquial   bool    `json:"colloquial"`
	Relevance    float64 `json:"relevance"`
	MostRelevant bool    `json:"mostRelevant"`
}

type Example struct {
	ID      int64  `json:"id"`
	Cluster int64  `json:"cluster"`
	Example string `json:"example"`
	Pos     *Pos   `json:"pos"`
}

type Stopwatch struct {
	Start   float64 `json:"start"`
	Ended   float64 `json:"ended"`
	Elapsed float64 `json:"elapsed"`
}

type TranslateResponse struct {
	ID                string      `json:"id"`
	From              string      `json:"from"`
	To                string      `json:"to"`
	Input             []string    `json:"input"`
	CorrectedText     interface{} `json:"correctedText"`
	Translation       []string    `json:"translation"`
	Engines           []string    `json:"engines"`
	LanguageDetection interface{} `json:"languageDetection"`
	ContextResults    interface{} `json:"contextResults"`
	Truncated         bool        `json:"truncated"`
	TimeTaken         int64       `json:"timeTaken"`
}

type SuggestResponse struct {
	Suggestions []*Suggestion           `json:"suggestions"`
	Fuzzy1      []*Suggestion           `json:"fuzzy1"`
	Fuzzy2      []*Suggestion           `json:"fuzzy2"`
	Request     *SuggestResponseRequest `json:"request"`
	TimeMS      int64                   `json:"time_ms"`
}

type Suggestion struct {
	Lang       string `json:"lang"`
	Suggestion string `json:"suggestion"`
	Weight     int64  `json:"weight"`
	IsFromDict bool   `json:"isFromDict"`
}

type SuggestResponseRequest struct {
	Search     string `json:"search"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
	MaxResults int64  `json:"max_results"`
	Mode       int64  `json:"mode"`
}

type QueryResponse struct {
	List                     []*QueryResult         `json:"list"`
	Nrows                    int64                  `json:"nrows"`
	NrowsExact               int64                  `json:"nrows_exact"`
	Pagesize                 int64                  `json:"pagesize"`
	Npages                   int64                  `json:"npages"`
	Page                     int64                  `json:"page"`
	RemovedSuperstrings      []string               `json:"removed_superstrings"`
	DictionaryEntryList      []*DictionaryEntryList `json:"dictionary_entry_list"`
	DictionaryOtherFrequency int64                  `json:"dictionary_other_frequency"`
	DictionaryNrows          int64                  `json:"dictionary_nrows"`
	TimeMS                   int64                  `json:"time_ms"`
	Request                  *QueryResponseRequest  `json:"request"`
	Suggestions              []*Suggestion          `json:"suggestions"`
	DymCase                  int64                  `json:"dym_case"`
	DymList                  []interface{}          `json:"dym_list"`
	DymApplied               interface{}            `json:"dym_applied"`
	DymNonadaptedSearch      interface{}            `json:"dym_nonadapted_search"`
	DymPairApplied           interface{}            `json:"dym_pair_applied"`
	DymNonadaptedSearchPair  interface{}            `json:"dym_nonadapted_search_pair"`
	DymPair                  interface{}            `json:"dym_pair"`
	ExtractedPhrases         []interface{}          `json:"extracted_phrases"`
	SourceTransliterations   []interface{}          `json:"sourceTransliterations"`
}

type DictionaryEntryList struct {
	Frequency           int64                 `json:"frequency"`
	Term                string                `json:"term"`
	IsFromDict          bool                  `json:"isFromDict"`
	IsPrecomputed       bool                  `json:"isPrecomputed"`
	Stags               []string              `json:"stags"`
	Article             *string               `json:"article"`
	Pos                 *string               `json:"pos"`
	Sourcepos           []string              `json:"sourcepos"`
	Variant             interface{}           `json:"variant"`
	Domain              interface{}           `json:"domain"`
	Definition          *string               `json:"definition"`
	Vowels1             interface{}           `json:"vowels1"`
	Transliteration1    interface{}           `json:"transliteration1"`
	Vowels2             interface{}           `json:"vowels2"`
	Transliteration2    interface{}           `json:"transliteration2"`
	AlignFreq           int64                 `json:"alignFreq"`
	ReverseValidated    bool                  `json:"reverseValidated"`
	PosGroup            int64                 `json:"pos_group"`
	IsTranslation       bool                  `json:"isTranslation"`
	IsFromLOCD          bool                  `json:"isFromLOCD"`
	InflectedForms      []DictionaryEntryList `json:"inflectedForms"`
	IsHiddenInFirstView bool                  `json:"isHiddenInFirstView"`
}

type QueryResult struct {
	SText string `json:"s_text"`
	TText string `json:"t_text"`
	Ref   string `json:"ref"`
	Cname string `json:"cname"`
	URL   string `json:"url"`
	Ctags string `json:"ctags"`
	Pba   bool   `json:"pba"`
}

type QueryResponseRequest struct {
	SourceText      string `json:"source_text"`
	TargetText      string `json:"target_text"`
	SourceLang      string `json:"source_lang"`
	TargetLang      string `json:"target_lang"`
	Npage           int64  `json:"npage"`
	Corpus          string `json:"corpus"`
	Nrows           int64  `json:"nrows"`
	Adapted         bool   `json:"adapted"`
	NonadaptedText  string `json:"nonadapted_text"`
	RudeWords       bool   `json:"rude_words"`
	Colloquialisms  bool   `json:"colloquialisms"`
	RiskyWords      bool   `json:"risky_words"`
	Mode            int64  `json:"mode"`
	ExprSug         int64  `json:"expr_sug"`
	DymApply        bool   `json:"dym_apply"`
	PosReorder      int64  `json:"pos_reorder"`
	Device          int64  `json:"device"`
	SplitLong       bool   `json:"split_long"`
	HasLocd         bool   `json:"has_locd"`
	AvoidLOCD       bool   `json:"avoidLOCD"`
	SourcePos       string `json:"source_pos"`
	ToolwordRequest bool   `json:"toolwordRequest"`
}
