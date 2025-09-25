import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	const res = await fetch('http://localhost:8080/api/products');

	if (!res.ok) {
		throw new Error('failed to fetch products');
	}
	const products = await res.json();

	return {
		products
	};
};
