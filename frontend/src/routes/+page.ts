export async function load() {
    const res = await fetch("http://localhost:8080/flashcards");

    const flashcards = await res.json()

    return {
        flashcards: flashcards
    };
}
