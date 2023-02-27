<script lang="ts">
	/**
	 * ID prop is used as a Name and ID
	 */
	export let id: string;
	export let labelText: string;
	export let type: 'email' | 'text' = 'text';
	export let placeholder: string;
</script>

<div class="flex flex-col">
	<label for={id} class="text-sm text-zinc-300 mb-1">{labelText}</label>
	<input
		{type}
		{placeholder}
		{id}
		name={id}
		class="rounded border border-zinc-600 px-3 py-1.5 bg-zinc-700 outline-none focus:ring focus:ring-purple-600 caret-purple-600"
	/>
</div>
