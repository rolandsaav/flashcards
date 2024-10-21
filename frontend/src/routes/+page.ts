export const ssr = false;
import { redirect } from "@sveltejs/kit";
import type { Flashcard, GetFlashcardsResponse } from "../types/Flashcard";
import type { Load } from "@sveltejs/kit"


export const load: Load = async ({ fetch }) => {
    const res = await fetch("http://localhost:8080/flashcards", {
        credentials: "include"
    });

    const response: GetFlashcardsResponse = await res.json()
    let flashcards: Flashcard[] = []
    if (res.status == 401) {
        redirect(302, "/login")
    } else {
        flashcards = response.data
    }


    return {
        flashcards: flashcards,
    };
}
