<script lang="ts">
	import type { PageProps } from './$types';
	let { data }: PageProps = $props();
</script>

<h1>Products</h1>

{#each data.products as product (product.id)}
	<p>
		<a href={`/products/${product.id}`}> {product.name}</a>
	</p>
{/each}
