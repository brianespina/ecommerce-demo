<script lang="ts">
	import Session from 'supertokens-web-js/recipe/session';

	let isLogged = $state(false);
	let isLoading = $state(true);

	async function doesSessionExist() {
		if (typeof window === 'undefined') return {};

		if (await Session.doesSessionExist()) {
			// user is logged in
			isLogged = true;
		} else {
			// user has not logged in yet
			isLogged = false;
		}
		isLoading = false;
	}
	doesSessionExist();
</script>

{#if isLoading}
	...Loading
{:else if isLogged}
	Welcome user
{:else}
	Welcome guest
{/if}
