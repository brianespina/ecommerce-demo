import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	const id = params.id;
	const res = await fetch(`http://localhost:8080/api/products/${id}`);

	if (!res.ok) {
		throw new Error('failed to fetch products');
	}
	const product = await res.json();
	return {
		product
	};
};
