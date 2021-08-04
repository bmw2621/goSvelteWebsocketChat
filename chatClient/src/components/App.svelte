<script>
	import Chat from "./Chat.svelte"
	import Login from "./Login.svelte"
	import {server} from "../store"
</script>

<main>
	{#if $server}
	<Chat server={$server}/>
	{:else}
	<Login />
	{/if}
</main>

<style>
	main {
		text-align: center;
		padding: 0;
		max-width: 240px;
		margin: 0;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>