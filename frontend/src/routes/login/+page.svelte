<script lang="ts">
	import { register } from '$lib/register';
	import { login } from '$lib/login';
	let username = '';
	let password = '';
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
