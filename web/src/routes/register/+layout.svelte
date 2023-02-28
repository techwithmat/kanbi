<script lang="ts">
	import Logo from '#components/Logo.svelte';
</script>

<header class="h-16 w-full border-b border-zinc-700 bg-zinc-900">
	<div class="h-full flex items-center px-3 max-w-screen-xl mx-auto">
		<Logo />
	</div>
</header>
<main>
	<slot />
</main>
