<script lang="ts">
    import { Button } from '$lib/components/ui/button';
    import { FestivalCard } from '$lib/components/calendar';
    import { NewsletterSignup } from '$lib/components/newsletter';
    import { festivals, getThisWeekFestivals } from '$lib/data/festivals';
    import { heritageLabels } from '$lib/types/festival';

    // Featured festivals for hero section
    const featuredFestivals = festivals.slice(0, 3);

    // "This Week" using demo date
    const thisWeekFestivals = getThisWeekFestivals(new Date('2026-02-10'));

    // Stats for impact section
    const stats = {
        festivals: festivals.length,
        heritageTypes: new Set(festivals.map((f) => f.heritageType)).size,
        regions: new Set(festivals.map((f) => f.region)).size,
    };
</script>

<svelte:head>
	<title>KULTUR | Your Guide to Trinidad & Tobago's Cultural Celebrations</title>
	<meta name="description" content="Discover 20+ cultural festivals across Trinidad & Tobago. From Carnival to Hosay, find your next authentic cultural experience." />
	<meta property="og:title" content="KULTUR | Trinidad & Tobago Cultural Festivals" />
	<meta property="og:description" content="Discover 20+ cultural festivals across Trinidad & Tobago. From Carnival to Hosay, find your next authentic cultural experience." />
	<meta property="og:url" content="https://kultur-tt.vercel.app" />
	<meta name="twitter:title" content="KULTUR | Trinidad & Tobago Cultural Festivals" />
	<meta name="twitter:description" content="Discover 20+ cultural festivals across Trinidad & Tobago." />
</svelte:head>

<div class="min-h-screen bg-background">
	<!-- Header -->
	<header class="border-b bg-card/80 backdrop-blur-sm sticky top-0 z-50">
		<div class="container max-w-6xl mx-auto px-4 py-4">
			<div class="flex items-center justify-between">
				<a href="/" class="flex items-center gap-2 hover:opacity-80 transition-opacity">
					<img src="/logo.png" alt="" class="h-10 w-10 logo-img" loading="eager" width="40" height="40" />
					<span class="text-xl font-bold text-tt-red">KULTUR</span>
				</a>
				<nav class="flex items-center gap-4">
					<a href="/" class="text-sm font-medium text-tt-red">Home</a>
					<a href="/festivals" class="text-sm hover:text-tt-red transition-colors">Calendar</a>
				</nav>
			</div>
		</div>
	</header>

	<div class="min-h-[calc(100dvh-72px)] flex flex-col">
		<!-- Hero Section -->
		<section class="relative bg-linear-to-br from-tt-black via-tt-black to-tt-red/20 text-white flex-1 flex items-center">
			<div class="absolute inset-0 bg-[url('https://www.cultursmag.com/wp-content/uploads/2024/05/Carnival-in-Trinidad-Photo-credit-Hayden-Greene.jpg')] bg-cover bg-center opacity-20"></div>
			<div class="relative container max-w-6xl mx-auto px-4 py-12 lg:py-20">
				<div class="max-w-3xl">
					<h1 class="text-4xl lg:text-6xl font-bold mb-6">
						Discover Trinidad & Tobago's
						<span class="text-tt-gold">Cultural Festivals</span>
					</h1>
					<p class="text-lg lg:text-xl text-gray-300 mb-8 leading-relaxed">
						Your first-timer's guide to 20+ festivals spanning African, Indian, Indigenous, and mixed heritage.
						Know what to expect, how to participate, and experience authentic culture.
					</p>
					<div class="flex flex-wrap gap-4">
						<Button href="/festivals" size="lg" class="bg-tt-red hover:bg-tt-red-dark text-white">
							Explore the Calendar
							<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="ml-2">
								<path d="m9 18 6-6-6-6"/>
							</svg>
						</Button>
						<Button href="#this-week" variant="outline" size="lg" class="border-white text-black hover:bg-gray-1000">
							What's Coming Up
						</Button>
					</div>
				</div>
			</div>
		</section>

		<!-- Stats Section -->
		<section class="bg-card border-b shrink-0">
			<div class="container max-w-6xl mx-auto px-4 py-8">
				<div class="grid grid-cols-3 gap-4 text-center">
					<div>
						<p class="text-3xl font-bold text-tt-red">{stats.festivals}+</p>
						<p class="text-sm text-muted-foreground">Cultural Festivals</p>
					</div>
					<div>
						<p class="text-3xl font-bold text-tt-red">{stats.heritageTypes}</p>
						<p class="text-sm text-muted-foreground">Heritage Types</p>
					</div>
					<div>
						<p class="text-3xl font-bold text-tt-red">{stats.regions}</p>
						<p class="text-sm text-muted-foreground">Regions Covered</p>
					</div>
				</div>
			</div>
		</section>
	</div>

	<!-- This Week Section -->
	<section id="this-week" class="container max-w-6xl mx-auto px-4 py-12">
		<div class="flex items-center justify-between mb-6">
			<div>
				<h2 class="text-2xl lg:text-3xl font-bold">This Week in T&T</h2>
				<p class="text-muted-foreground">Festivals happening in the next 2 weeks</p>
			</div>
			<Button href="/festivals" variant="outline">
				View All
				<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="ml-1">
					<path d="m9 18 6-6-6-6"/>
				</svg>
			</Button>
		</div>

		{#if thisWeekFestivals.length > 0}
			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
				{#each thisWeekFestivals as festival (festival.id)}
					<FestivalCard {festival} />
				{/each}
			</div>
		{:else}
			<div class="bg-muted/50 rounded-lg p-12 text-center">
				<p class="text-muted-foreground text-lg">No festivals in the immediate upcoming period.</p>
				<p class="text-sm text-muted-foreground mt-2">
					<a href="/festivals" class="text-tt-red hover:underline">Browse the full calendar</a> to find upcoming events.
				</p>
			</div>
		{/if}
	</section>

	<!-- Why This Matters Section -->
	<section class="bg-muted/30 py-12">
		<div class="container max-w-6xl mx-auto px-4">
			<div class="max-w-3xl mx-auto text-center mb-10">
				<h2 class="text-2xl lg:text-3xl font-bold mb-4">More Than Just Carnival</h2>
				<p class="text-muted-foreground">
					Trinidad & Tobago has a rich tapestry of cultural celebrations. Most people only know Carnival, 
					but there's so much more to discover.
				</p>
			</div>

			<div class="grid md:grid-cols-3 gap-6">
				<div class="bg-card rounded-lg p-6 border">
					<div class="w-12 h-12 bg-heritage-indigenous/20 rounded-lg flex items-center justify-center mb-4">
						<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-heritage-indigenous">
							<circle cx="12" cy="12" r="10"/>
							<path d="M12 2a14.5 14.5 0 0 0 0 20 14.5 14.5 0 0 0 0-20"/>
							<path d="M2 12h20"/>
						</svg>
					</div>
					<h3 class="font-semibold mb-2">8,000 Years of History</h3>
					<p class="text-sm text-muted-foreground">
						The Santa Rosa First Peoples Community preserves traditions from millennia of continuous Amerindian presence.
					</p>
				</div>

				<div class="bg-card rounded-lg p-6 border">
					<div class="w-12 h-12 bg-heritage-indian/20 rounded-lg flex items-center justify-center mb-4">
						<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-heritage-indian">
							<path d="M12 2v8"/>
							<path d="m4.93 10.93 1.41 1.41"/>
							<path d="M2 18h2"/>
							<path d="M20 18h2"/>
							<path d="m19.07 10.93-1.41 1.41"/>
							<path d="M22 22H2"/>
							<path d="m8 6 4-4 4 4"/>
							<path d="M16 18a4 4 0 0 0-8 0"/>
						</svg>
					</div>
					<h3 class="font-semibold mb-2">Hidden Stories</h3>
					<p class="text-sm text-muted-foreground">
						In 1884, British authorities massacred 22 people at a Hosay procession. Most Trinis have never heard of it.
					</p>
				</div>

				<div class="bg-card rounded-lg p-6 border">
					<div class="w-12 h-12 bg-heritage-african/20 rounded-lg flex items-center justify-center mb-4">
						<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-heritage-african">
							<path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/>
							<circle cx="9" cy="7" r="4"/>
							<path d="M22 21v-2a4 4 0 0 0-3-3.87"/>
							<path d="M16 3.13a4 4 0 0 1 0 7.75"/>
						</svg>
					</div>
					<h3 class="font-semibold mb-2">Community Connection</h3>
					<p class="text-sm text-muted-foreground">
						374,500 Trinis in diaspora looking to reconnect. These festivals are bridges to heritage.
					</p>
				</div>
			</div>
		</div>
	</section>

	<!-- CTA Section -->
	<section class="bg-tt-red text-white py-12">
		<div class="container max-w-6xl mx-auto px-4 text-center">
			<h2 class="text-2xl lg:text-3xl font-bold mb-4">Ready to Explore?</h2>
			<p class="text-lg opacity-90 mb-6 max-w-2xl mx-auto">
				Whether you're a first-timer, returning local, or curious visitor,
				find your next authentic cultural experience.
			</p>
			<Button href="/festivals" size="lg" class="bg-white text-tt-red hover:bg-gray-100">
				Browse the Festival Calendar
			</Button>
		</div>
	</section>

	<!-- Featured Festivals -->
	<section class="container max-w-6xl mx-auto px-4 py-12">
		<h2 class="text-2xl lg:text-3xl font-bold mb-6">Featured Festivals</h2>
		<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
			{#each featuredFestivals as festival (festival.id)}
				<FestivalCard {festival} />
			{/each}
		</div>
		<div class="text-center mt-8">
			<Button href="/festivals" size="lg" variant="outline">
				View All {festivals.length} Festivals
			</Button>
		</div>
	</section>

	<!-- Newsletter Signup Section -->
	<NewsletterSignup variant="hero" />

	<!-- Footer -->
	<footer class="border-t bg-card">
		<div class="container max-w-6xl mx-auto px-4 py-8">
			<div class="flex flex-col md:flex-row items-center justify-between gap-4">
				<div class="text-center md:text-left">
					<div class="flex items-center gap-2 justify-center md:justify-start">
						<img src="/logo.png" alt="" class="h-8 w-8 logo-img" loading="eager" width="32" height="32" />
						<p class="font-semibold text-tt-red">KULTUR</p>
					</div>
					<p class="text-sm text-muted-foreground">Your guide to Trinidad & Tobago's cultural celebrations</p>
				</div>
				<p class="text-sm text-muted-foreground">
					Trini2DBone 2026  · TrinbagoTech Hackathon 2026
				</p>
			</div>
		</div>
	</footer>
</div>
