<script lang="ts">
    import { logout } from "$lib/logout"
    import { goto } from "$app/navigation"
	import { user } from "../auth";

    const handleLogout = async () => {
        const loggedOut = await logout()

        if (loggedOut) {
            console.log("Unset user state")
            goto('/login');
        } else {
            console.log("Could not log out?")
        }
    }
</script>

<h1 class="navbar">
    <a href="/">Home</a>
    <a href="/login">Login</a>
    <button on:click={handleLogout}>Logout</button>
    <h3>
        {#if $user == null}
        Not logged in
        {:else}
        Logged in as {$user.username}
        {/if}
    </h3>
    
</h1>

<slot></slot>

<style>
    .navbar {
        display: flex;
        gap: 1rem;
    }
</style>