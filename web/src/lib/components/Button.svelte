<script lang="ts">
	export let type: 'submit' | 'button' = 'button';
</script>

<button
	class="flex items-center justify-center gap-3 px-3 py-1.5 rounded bg-purple-600 hover:bg-purple-500 transition-colors font-medium"
	{type}
	on:click
>
	<slot />
</button>
