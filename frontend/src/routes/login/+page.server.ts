import { redirect } from "@sveltejs/kit";
import type { Actions } from "@sveltejs/kit"


export const actions: Actions = {
    login: async ({ request }) => {
        console.log("Login action")
        const data = await request.formData();
        const username = data.get("username")
        const password = data.get("password")
        let shouldRedirect = false
        try {
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    username: username,
                    password: password
                }),
                credentials: 'include'
            });

            shouldRedirect = response.ok
        } catch (error) {
            console.error(error);
        }
        if (shouldRedirect) {
            console.log("Login successful")
            throw redirect(303, '/');
        }

        return { success: true };
    },
}
