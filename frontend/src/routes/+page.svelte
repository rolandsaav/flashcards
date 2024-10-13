<script>
	/** @type {import('./$types').PageData} */
	export let data;
</script>

<h1>Welcome to Flashcards</h1>
<p>This is all of the flashcards</p>

<ul>
	{#each data.flashcards as item}
		<li>{item.term} x {item.definition}</li>
	{/each}
</ul>
