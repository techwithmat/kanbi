<script>
	import BoardMenu from '#components/BoardMenu.svelte';
	import BoardSettings from '#components/BoardSettings.svelte';
</script>

<header class="h-16 w-full fixed top-0 z-50 border-b border-zinc-700">
	<div class="h-full flex items-center justify-between px-3 max-w-screen-xl mx-auto">
		<a href="/" class="text-xl font-medium">Kanbi</a>

		<div class="flex items-center gap-2">
			<BoardMenu />
			<BoardSettings />
		</div>
	</div>
</header>
