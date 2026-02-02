import { festivals, getFestivalBySlug } from '$lib/data/festivals';
import { getMemoriesByFestivalSlug } from '$lib/data/memories';
import { config } from '$lib/config';
import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import memoriesData from '$lib/data/memories.json';

export const load: PageLoad = async ({ params }) => {
	// Try API first if enabled, otherwise use mock data
	let festival;
	
	if (config.useApi) {
		try {
			festival = await getFestivalBySlug(params.slug);
		} catch (err) {
			// Fallback to mock data
			festival = festivals.find((f) => f.slug === params.slug);
		}
	} else {
		festival = festivals.find((f) => f.slug === params.slug);
	}

	if (!festival) {
		error(404, {
			message: 'Festival not found'
		});
	}

	// Get memories for this festival (API-aware)
	const memories = await getMemoriesByFestivalSlug(params.slug);

	console.log('Loaded memories for festival:', memories);
	console.log(params.slug)

	// Get related festivals (same heritage type, excluding current)
	const relatedFestivals = festivals
		.filter((f) => f.heritageType === festival.heritageType && f.id !== festival.id)
		.slice(0, 3);

	return {
		festival,
		memories,
		relatedFestivals
	};
};
