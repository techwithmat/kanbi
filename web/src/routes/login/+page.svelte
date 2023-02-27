<script lang="ts">
	import Button from '#components/Button.svelte';
	import Input from '#components/Input.svelte';
	import { type LoginFormData, LoginUser } from '#services/auth';

	async function handleOnSubmit(e: SubmitEvent) {
		const formData = Object.fromEntries(new FormData(e.target as HTMLFormElement)) as LoginFormData;

		await LoginUser(formData);
	}
</script>

<div class="w-full px-5 my-32 sm:my-40">
	<section class="max-w-xs mx-auto">
		<h1 class="text-3xl text-center font-bold tracking-wide	 mb-8 pb-4 border-b border-zinc-600">
			Login
		</h1>

		<form class="flex flex-col gap-3" on:submit|preventDefault={handleOnSubmit}>
			<Input id="email" placeholder="Enter your email" labelText="Email" type="email" />
			<Input id="password" placeholder="Enter your password" labelText="Password" />
			<Button type="submit">Submit</Button>
		</form>
	</section>
</div>
