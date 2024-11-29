package wordnet

type LexicalResource struct {
	Lexicon []Lexicon `xml:",any"`
}

type Lexicon struct {
	ID                  string               `xml:"id,attr"`
	Label               string               `xml:"label,attr"`
	Language            string               `xml:"language,attr"`
	Email               string               `xml:"email,attr"`
	License             string               `xml:"license,attr"`
	Version             string               `xml:"version,attr"`
	URL                 string               `xml:"url,attr,omitempty"`
	LexicalEntries      []LexicalEntry       `xml:"LexicalEntry"`
	Synsets             []Synset             `xml:"Synset,omitempty"`
	SyntacticBehaviours []SyntacticBehaviour `xml:"SyntacticBehaviour,omitempty"`
}

type Requires struct {
	ID      string `xml:"id,attr"`
	Version string `xml:"version,attr"`
	URL     string `xml:"url,attr,omitempty"`
}

type LexicalEntry struct {
	ID                  string               `xml:"id,attr"`
	Lemma               Lemma                `xml:"Lemma"`
	Forms               []Form               `xml:"Form,omitempty"`
	Senses              []Sense              `xml:"Sense,omitempty"`
	SyntacticBehaviours []SyntacticBehaviour `xml:"SyntacticBehaviour,omitempty"`
}

type Lemma struct {
	WrittenForm    string          `xml:"writtenForm,attr"`
	PartOfSpeech   string          `xml:"partOfSpeech,attr"`
	Pronunciations []Pronunciation `xml:"Pronunciation"`
}

type Pronunciation struct {
	Variety string `xml:"variety,attr"`
	Text    string `xml:",chardata"`
}

type Form struct {
	ID          string `xml:"id,attr,omitempty"`
	WrittenForm string `xml:"writtenForm,attr"`
	Script      string `xml:"script,attr,omitempty"`
}

type Synset struct {
	ID              string           `xml:"id,attr"`
	Ili             string           `xml:"ili,attr"`
	PartOfSpeech    string           `xml:"partOfSpeech,attr,omitempty"`
	Contributor     string           `xml:"dc:contributor,attr,omitempty"`
	Definitions     []string         `xml:"Definition,omitempty"`
	ILIDefinition   string           `xml:"ILIDefinition,omitempty"`
	SynsetRelations []SynsetRelation `xml:"SynsetRelation,omitempty"`
	Examples        []string         `xml:"Example,omitempty"`
}

type SyntacticBehaviour struct {
	ID                     string `xml:"id,attr,omitempty"`
	SubcategorizationFrame string `xml:"subcategorizationFrame,attr"`
	Senses                 string `xml:"senses,attr,omitempty"`
}

type Sense struct {
	ID             string          `xml:"id,attr"`
	Synset         string          `xml:"synset,attr"`
	SenseRelations []SenseRelation `xml:"SenseRelation"`
}

type SenseRelation struct {
	Target  string `xml:"target,attr"`
	RelType string `xml:"relType,attr"`
}
type SynsetRelation struct {
	Target  string `xml:"target,attr"`
	RelType string `xml:"relType,attr"`
}
