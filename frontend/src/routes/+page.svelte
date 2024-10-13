<script lang="ts">
	import type { Flashcard } from '../types/Flashcard';
	export let data: { flashcards: Flashcard[] };

	let flashcards = data.flashcards;

	let term = '';
	let definition = '';
	let owner = 0;
	let i = 0;

	$: selected = flashcards[i];
	$: resetInputs(selected);

	function resetInputs(card: Flashcard) {
		term = card ? card.term : '';
		definition = card ? card.definition : '';
		owner = card ? card.owner_id : NaN;
	}

	async function createFlashcard(event: MouseEvent) {
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
				flashcards = flashcards.concat(body);
			} else {
				console.log('There was a problem');
			}
		} catch (error) {
			console.error(error);
		}
	}

	async function updateFlashcard(event: MouseEvent) {
		event.preventDefault();

		try {
			const response = await fetch('http://localhost:8080/flashcards', {
				method: 'PATCH',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({
					id: selected.id,
					owner_id: owner,
					term: term,
					definition: definition
				})
			});

			if (response.ok) {
				console.log('Updated flashcard');
				const body: Flashcard = await response.json();
				selected.term = body.term;
				selected.definition = body.definition;
				selected.owner_id = body.owner_id;
				flashcards = flashcards;
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

<select size={5} bind:value={i}>
	{#each flashcards as item, index}
		<option value={index}>{item.term} : {item.definition} &slarr; {item.owner_id} : {index}</option>
	{/each}
</select>

<h2>Create New Flashcard</h2>
<form class="inputs">
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

	<div class="buttons">
		<button on:click={createFlashcard}>Create</button>
		<button on:click={updateFlashcard}>Update</button>
		<button>Delete</button>
	</div>
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

	.buttons {
		display: flex;
		gap: 5px;
	}

	.buttons button {
		flex-grow: 1;
	}
</style>
