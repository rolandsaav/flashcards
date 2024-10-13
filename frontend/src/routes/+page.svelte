<script lang="ts">
	import type { Flashcard } from '../types/Flashcard';
	export let data: { flashcards: Flashcard[] };

	let term = '';
	let definition = '';
	let owner = 0;

	async function createFlashcard(event: SubmitEvent) {
		event.preventDefault();

		try {
			const response = await fetch('http://localhost:8080/flashcards', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({
					owner_id: owner,
					term: term,
					definition: definition
				})
			});

			if (response.ok) {
				console.log('Created flashcard');
				const body: Flashcard = await response.json();
				data.flashcards = data.flashcards.concat(body);
			} else {
				console.log('There was a problem');
			}
		} catch (error) {
			console.error(error);
		}
	}
</script>

<h1>Welcome to Flashcards</h1>
<p>This is all of the flashcards</p>

<ul>
	{#each data.flashcards as item}
		<li>{item.term} : {item.definition} &slarr; {item.owner_id}</li>
	{/each}
</ul>

<h2>Create New Flashcard</h2>
<form class="inputs" on:submit={createFlashcard}>
	<div class="input">
		<label for="term">Term</label>
		<input bind:value={term} name="term" id="term" type="text" />
	</div>
	<div class="input">
		<label for="definition">Definition</label>
		<textarea bind:value={definition} name="definition" id="definition" rows="2" />
	</div>
	<div class="input">
		<label for="owner">Owner ID</label>
		<input bind:value={owner} name="owner" id="owner" type="number" />
	</div>

	<button type="submit">Create</button>
</form>

<style>
	.inputs {
		width: 20rem;
		display: flex;
		flex-direction: column;
	}

	.input {
		display: flex;
		margin-bottom: 1rem;
		align-items: baseline;
	}

	.input label {
		margin-right: 1rem;
	}

	.input input {
		flex-grow: 1;
	}

	.input textarea {
		flex-grow: 1;
		resize: vertical;
	}
</style>
