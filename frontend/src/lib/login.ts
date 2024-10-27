import { goto } from "$app/navigation"
import { setContext } from 'svelte';
import { writable } from 'svelte/store';


export const login = async (username: string, password: string) => {
    console.log("Login")
    const response = await fetch('http://localhost:8080/login', {
        method: "POST",
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            username: username,
            password: password
        }),
        credentials: "include"
    })

    if (response.ok) {
        console.log("Logged in")
        const data = await response.json()
        const user = writable();
        $: user.set(data)
        setContext("user", user)
        console.log(data)
        goto("/")
    }

    
}
