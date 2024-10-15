export type Flashcard = {
    id: number,
    owner_id: number,
    term: string,
    definition: string,
}

export type GetFlashcardsResponse = {
    data: Flashcard[],
    error: string,
}
