<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import SuperTokens from 'supertokens-web-js';
	import Session from 'supertokens-web-js/recipe/session';
	import EmailPassword from 'supertokens-web-js/recipe/emailpassword';

	if (typeof window !== 'undefined') {
		SuperTokens.init({
			appInfo: {
				apiDomain: 'http://localhost:8080/',
				apiBasePath: '/auth',
				appName: '...'
			},
			recipeList: [Session.init(), EmailPassword.init()]
		});
	}

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children?.()}
