<script lang="ts">
import type { PageProps } from './$types';
let {data}: Pageprops = $props();
</script>

<h1>Products</h1>
{#each data.products as product}
	<p>
		{product.name}
	</p>
{/each}

