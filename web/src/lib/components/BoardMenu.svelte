<script lang="ts">
	/**
	 * TODO:
	 * 1. Add a function to read the userName
	 * 2. Add a fetch function that will be triggered when the user click the menu
	 * 3. When the user Enters a Board we need to store the ID of the board so that when the user opens the application we will be on the same Board.
	 */
	import ArrowDown from '#icons/ArrowDown.svelte';
	import Board from '#icons/Board.svelte';

	// This will be replaced with a cookie or some local storage logic to store the ID
	export let selectedOption = 'Platform Launch';

	// The state to open and close the Menu
	export let isOpen = false;

	// The user name will be obtained from the User Session
	const userName: string = '';
	// The boards will be obtained dynamically when the user clicks to see all the boards
	const BOARDS: any = [];
</script>

<div class="relative max-w-xl">
	<button
		class="inline-flex gap-3 px-3 py-1.5 border border-transparent rounded bg-purple-600 hover:bg-purple-500 transition-colors font-medium"
		on:click={() => (isOpen = !isOpen)}
	>
		{selectedOption}
		<ArrowDown />
	</button>

	{#if isOpen}
		<ul
			class="absolute z-50 w-full mt-1 overflow-auto text-base bg-zinc-800 border border-zinc-700 rounded-md shadow-lg max-h-60 divide-y-2 divide-zinc-700 py-1"
		>
			{#each BOARDS as board}
				<li class="select-none">
					<a
						href="/{userName}/{board.id}"
						class="flex items-center gap-2 px-3 py-2.5 hover:bg-zinc-700"
					>
						<Board />
						{board.name}
					</a>
				</li>
			{/each}
		</ul>
	{/if}
</div>
