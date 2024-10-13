import type { Flashcard } from "../types/Flashcard";

export async function load() {
    const res = await fetch("http://localhost:8080/flashcards");

    const flashcards: Flashcard[] = await res.json()

    return {
        flashcards: flashcards
    };
}
