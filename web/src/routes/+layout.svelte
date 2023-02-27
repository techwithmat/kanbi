<script>
	import '../app.css';
	import Header from '#components/Header.svelte';
</script>

<Header />
<main>
	<slot />
</main>

<style lang="postcss">
	:global(html) {
		@apply bg-zinc-900 text-zinc-200;
	}
</style>
