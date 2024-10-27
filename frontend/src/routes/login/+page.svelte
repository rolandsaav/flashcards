<script lang="ts">
	import { register } from '$lib/register';
	import { setContext } from 'svelte';
	import { writable } from 'svelte/store';
	import { goto } from "$app/navigation"
	let username = '';
	let password = '';

	const login = async (username: string, password: string) => {
		console.log('Login');
		const response = await fetch('http://localhost:8080/login', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({
				username: username,
				password: password
			}),
			credentials: 'include'
		});

		if (response.ok) {
			console.log('Logged in');
			const data = await response.json();
			const user = writable();
			user.set(data);
			setContext('user', user);
			console.log(data);
			goto('/');
		}
	};
</script>

<h1>Login Page</h1>
<form class="inputs">
	<div class="input">
		<label for="term">Username</label>
		<input bind:value={username} name="username" id="username" type="text" />
	</div>
	<div class="input">
		<label for="definition">Password</label>
		<input bind:value={password} name="password" id="password" type="password" />
	</div>

	<div class="buttons">
		<button on:click={() => login(username, password)}>Login</button>
		<button on:click={() => register(username, password)}>Register</button>
	</div>
</form>

<style>
	.buttons {
		display: flex;
	}
</style>
