import { user } from "../auth"

export const logout = async () => {
    console.log("Logout")
    const response = await fetch('http://localhost:8080/logout', {
        method: "POST",
        headers: { 'Content-Type': 'application/json' },
        credentials: "include"
    })

    if (response.ok) {
        console.log("Logged out")
        const data = await response.json()
        user.set(null);
        console.log(user)
        return true
    }

    return false
}
