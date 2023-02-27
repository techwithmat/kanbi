<script>
	import BoardMenu from '#components/BoardMenu.svelte';
	import BoardSettings from '#components/BoardSettings.svelte';
	import Logo from './Logo.svelte';
</script>

<header class="h-16 w-full fixed top-0 z-50 border-b border-zinc-700">
	<div class="h-full flex items-center justify-between gap-2 px-3 max-w-screen-xl mx-auto">
		<Logo />

		<div class="flex items-center gap-2">
			<BoardMenu />
			<BoardSettings />
		</div>
	</div>
</header>
