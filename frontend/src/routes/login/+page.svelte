<script lang="ts">
	let status = 'Not logged in';
	let username = '';
	let password = '';
</script>

<h1>Login Page</h1>
<p>Status {status}</p>
<form class="inputs" method="POST" action="?/login">
	<div class="input">
		<label for="term">Username</label>
		<input bind:value={username} name="username" id="username" type="text" />
	</div>
	<div class="input">
		<label for="definition">Password</label>
		<input bind:value={password} name="password" id="password" type="password" />
	</div>

	<div class="buttons">
		<button>Login</button>
	</div>
</form>
