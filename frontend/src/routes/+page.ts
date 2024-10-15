import type { GetFlashcardsResponse } from "../types/Flashcard";

export async function load() {
    const res = await fetch("http://localhost:8080/flashcards");

    const response: GetFlashcardsResponse = await res.json()

    const flashcards = response.data

    return {
        flashcards: flashcards,
    };
}
