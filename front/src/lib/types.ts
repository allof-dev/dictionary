export interface Lemma {
    id: string
    written_form: string
    part_of_speech: string
    synsets: Synset[]
}

interface Synset {
    id: string
    words: string[]
    definitions: string[]
}