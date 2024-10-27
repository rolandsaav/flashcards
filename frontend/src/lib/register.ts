export const register = async (username: string, password: string) => {
    console.log("Register")
    const response = await fetch('http://localhost:8080/register', {
        method: "POST",
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            username: username,
            password: password
        }),
        credentials: "include"
    })

    if (response.ok) {
        console.log("created new user")
        const data = await response.json()
        console.log(data)
    }
}
