<script lang="ts">
    import { marked } from 'marked';
    import { Badge } from '$lib/components/ui/badge';
    import { Button } from '$lib/components/ui/button';
    import * as Card from '$lib/components/ui/card';
    import { Separator } from '$lib/components/ui/separator';
    import { FestivalCard } from '$lib/components/calendar';
    import { MemoryCard, MemoryForm } from '$lib/components/memories';
    import { NewsletterSignup } from '$lib/components/newsletter';
    import { heritageLabels, regionLabels } from '$lib/types/festival';
    import type { HeritageType } from '$lib/types/festival';
    import { formatDateRange, getRelativeTime } from '$lib/utils/calendar';

    const { data } = $props();
    const festival = $derived(data.festival);
    const memories = $derived(data.memories);
    const relatedFestivals = $derived(data.relatedFestivals);

    // Configure marked for safety
    marked.setOptions({
        breaks: true,
        gfm: true,
    });

    let linkCopied = $state(false);

    function getHeritageBadgeClass(heritage: HeritageType): string {
        const classes: Record<HeritageType, string> = {
            african: 'bg-heritage-african text-white',
            indian: 'bg-heritage-indian text-black',
            indigenous: 'bg-heritage-indigenous text-white',
            mixed: 'bg-heritage-mixed text-white',
            christian: 'bg-heritage-christian text-white',
        };
        return classes[heritage];
    }

    function renderMarkdown(content: string | null): string {
        if (!content) return '';
        return marked.parse(content) as string;
    }

    async function copyLink() {
        try {
            await navigator.clipboard.writeText(window.location.href);
            linkCopied = true;
            setTimeout(() => {
                linkCopied = false;
            }, 2000);
        } catch (err) {
            console.error('Failed to copy link:', err);
        }
    }
</script>

<svelte:head>
	<title>{festival.name} | KULTUR</title>
	<meta name="description" content={festival.summary} />
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
					<a href="/" class="text-sm hover:text-tt-red transition-colors">Home</a>
					<a href="/festivals" class="text-sm hover:text-tt-red transition-colors">Calendar</a>
				</nav>
			</div>
		</div>
	</header>

	<!-- Breadcrumb -->
	<div class="container max-w-6xl mx-auto px-4 py-4">
		<nav class="flex items-center gap-2 text-sm text-muted-foreground">
			<a href="/" class="hover:text-foreground transition-colors">Home</a>
			<span>/</span>
			<a href="/festivals" class="hover:text-foreground transition-colors">Festivals</a>
			<span>/</span>
			<span class="text-foreground">{festival.name}</span>
		</nav>
	</div>

	<!-- Hero Section -->
	<section class="relative">
		{#if festival.coverImageUrl}
			<div class="h-64 md:h-80 lg:h-96 overflow-hidden">
				<img
					src={festival.coverImageUrl}
					alt={festival.name}
					class="w-full h-full object-cover"
				/>
				<div class="absolute inset-0 bg-linear-to-t from-black/80 via-black/40 to-transparent"></div>
			</div>
			<div class="absolute bottom-0 left-0 right-0 p-6 md:p-8">
				<div class="container max-w-6xl mx-auto">
					<div class="flex flex-wrap gap-2 mb-3">
						<Badge class={getHeritageBadgeClass(festival.heritageType)}>
							{heritageLabels[festival.heritageType]}
						</Badge>
						<Badge variant="secondary">{festival.festivalType}</Badge>
						{#if festival.date2026Start}
							<Badge variant="outline" class="bg-white/10 text-white border-white/30">
								{getRelativeTime(festival.date2026Start)}
							</Badge>
						{/if}
					</div>
					<h1 class="text-3xl md:text-4xl lg:text-5xl font-bold text-white mb-2">
						{festival.name}
					</h1>
					<p class="text-lg text-gray-200 max-w-2xl">
						{festival.summary}
					</p>
				</div>
			</div>
		{:else}
			<div class="bg-muted py-12 md:py-16">
				<div class="container max-w-6xl mx-auto px-4">
					<div class="flex flex-wrap gap-2 mb-3">
						<Badge class={getHeritageBadgeClass(festival.heritageType)}>
							{heritageLabels[festival.heritageType]}
						</Badge>
						<Badge variant="secondary">{festival.festivalType}</Badge>
					</div>
					<h1 class="text-3xl md:text-4xl lg:text-5xl font-bold mb-2">
						{festival.name}
					</h1>
					<p class="text-lg text-muted-foreground max-w-2xl">
						{festival.summary}
					</p>
				</div>
			</div>
		{/if}
	</section>

	<main class="container max-w-6xl mx-auto px-4 py-8">
		<div class="grid grid-cols-1 lg:grid-cols-[1fr_320px] gap-8">
			<!-- Main Content -->
			<div class="space-y-8">
				<!-- Quick Info Card -->
				<Card.Root>
					<Card.Header>
						<Card.Title class="flex items-center gap-2">
							<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
								<circle cx="12" cy="12" r="10"/>
								<path d="M12 16v-4"/>
								<path d="M12 8h.01"/>
							</svg>
							At a Glance
						</Card.Title>
					</Card.Header>
					<Card.Content>
						<div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
							<div class="flex items-start gap-3">
								<div class="w-10 h-10 rounded-lg bg-tt-red/10 flex items-center justify-center shrink-0">
									<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-red">
										<rect width="18" height="18" x="3" y="4" rx="2" ry="2"/>
										<line x1="16" x2="16" y1="2" y2="6"/>
										<line x1="8" x2="8" y1="2" y2="6"/>
										<line x1="3" x2="21" y1="10" y2="10"/>
									</svg>
								</div>
								<div>
									<p class="text-sm font-medium">When</p>
									<p class="text-sm text-muted-foreground">
										{formatDateRange(festival.date2026Start, festival.date2026End)}
									</p>
									{#if festival.usualMonth}
										<p class="text-xs text-muted-foreground">Usually in {festival.usualMonth}</p>
									{/if}
								</div>
							</div>

							<div class="flex items-start gap-3">
								<div class="w-10 h-10 rounded-lg bg-tt-red/10 flex items-center justify-center shrink-0">
									<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-red">
										<path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"/>
										<circle cx="12" cy="10" r="3"/>
									</svg>
								</div>
								<div>
									<p class="text-sm font-medium">Where</p>
									<p class="text-sm text-muted-foreground">{regionLabels[festival.region]}</p>
								</div>
							</div>

							<div class="flex items-start gap-3">
								<div class="w-10 h-10 rounded-lg bg-tt-red/10 flex items-center justify-center shrink-0">
									<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-red">
										<path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2 2 6.477 2 12s4.477 10 10 10z"/>
										<path d="m9 12 2 2 4-4"/>
									</svg>
								</div>
								<div>
									<p class="text-sm font-medium">Type</p>
									<p class="text-sm text-muted-foreground capitalize">{festival.festivalType} · {festival.dateType} date</p>
								</div>
							</div>

							<div class="flex items-start gap-3">
								<div class="w-10 h-10 rounded-lg bg-tt-red/10 flex items-center justify-center shrink-0">
									<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-red">
										<path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/>
										<circle cx="9" cy="7" r="4"/>
										<path d="M22 21v-2a4 4 0 0 0-3-3.87"/>
										<path d="M16 3.13a4 4 0 0 1 0 7.75"/>
									</svg>
								</div>
								<div>
									<p class="text-sm font-medium">Heritage</p>
									<p class="text-sm text-muted-foreground">{heritageLabels[festival.heritageType]}</p>
								</div>
							</div>
						</div>
					</Card.Content>
				</Card.Root>

				<!-- The Story -->
				{#if festival.story}
					<section>
						<h2 class="text-2xl font-bold mb-4 flex items-center gap-2">
							<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-red">
								<path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/>
							</svg>
							The Story
						</h2>
						<Card.Root>
							<Card.Content class="pt-6 prose prose-sm max-w-none">
								{@html renderMarkdown(festival.story)}
							</Card.Content>
						</Card.Root>
					</section>
				{/if}

				<!-- What to Expect -->
				{#if festival.whatToExpect}
					<section>
						<h2 class="text-2xl font-bold mb-4 flex items-center gap-2">
							<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-red">
								<path d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z"/>
								<circle cx="12" cy="12" r="3"/>
							</svg>
							What to Expect
						</h2>
						<Card.Root>
							<Card.Content class="pt-6 prose prose-sm max-w-none">
								{@html renderMarkdown(festival.whatToExpect)}
							</Card.Content>
						</Card.Root>
					</section>
				{/if}

				<!-- How to Participate -->
				{#if festival.howToParticipate}
					<section>
						<h2 class="text-2xl font-bold mb-4 flex items-center gap-2">
							<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-red">
								<path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/>
								<circle cx="9" cy="7" r="4"/>
								<line x1="19" x2="19" y1="8" y2="14"/>
								<line x1="22" x2="16" y1="11" y2="11"/>
							</svg>
							How to Participate
						</h2>
						<Card.Root class="border-tt-red/20 bg-tt-red/5">
							<Card.Content class="pt-6 prose prose-sm max-w-none">
								{@html renderMarkdown(festival.howToParticipate)}
							</Card.Content>
						</Card.Root>
					</section>
				{/if}

				<!-- Practical Info -->
				{#if festival.practicalInfo}
					<section>
						<h2 class="text-2xl font-bold mb-4 flex items-center gap-2">
							<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-red">
								<path d="M14.5 4h-5L7 7H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-3l-2.5-3z"/>
								<circle cx="12" cy="13" r="3"/>
							</svg>
							Practical Information
						</h2>
						<Card.Root>
							<Card.Content class="pt-6 prose prose-sm max-w-none">
								{@html renderMarkdown(festival.practicalInfo)}
							</Card.Content>
						</Card.Root>
					</section>
				{/if}

				<!-- Community Memories Section -->
				<section id="memories">
					<h2 class="text-2xl font-bold mb-4 flex items-center gap-2">
						<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-red">
							<path d="M17 6.1H3"/>
							<path d="M21 12.1H3"/>
							<path d="M15.1 18H3"/>
						</svg>
						Community Memories
					</h2>
					
					{#if memories.length > 0}
						<div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
							{#each memories as memory (memory.id)}
								<MemoryCard {memory} />
							{/each}
						</div>
					{:else}
						<Card.Root class="mb-6">
							<Card.Content class="py-8 text-center">
								<svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="mx-auto text-muted-foreground/50 mb-4">
									<path d="M17 6.1H3"/>
									<path d="M21 12.1H3"/>
									<path d="M15.1 18H3"/>
								</svg>
								<p class="text-muted-foreground">No memories shared yet.</p>
								<p class="text-sm text-muted-foreground mt-1">Be the first to share your experience!</p>
							</Card.Content>
						</Card.Root>
					{/if}

					<!-- Memory Submission Form -->
					<MemoryForm festivalId={festival.id} festivalName={festival.name} />
				</section>
			</div>

			<!-- Sidebar -->
			<aside class="space-y-6">
				<!-- Back to Calendar -->
				<Button href="/festivals" variant="outline" class="w-full">
					<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-2">
						<path d="m15 18-6-6 6-6"/>
					</svg>
					Back to Calendar
				</Button>

				<!-- Date Reminder Card -->
				{#if festival.date2026Start}
					<Card.Root class="bg-tt-red text-white">
						<Card.Header class="pb-2">
							<Card.Title class="text-white text-lg">Mark Your Calendar</Card.Title>
						</Card.Header>
						<Card.Content>
							<p class="text-2xl font-bold mb-1">
								{formatDateRange(festival.date2026Start, festival.date2026End)}
							</p>
							<p class="text-white/80 text-sm">
								{getRelativeTime(festival.date2026Start)}
							</p>
						</Card.Content>
					</Card.Root>
				{/if}

				<!-- Share Section -->
				<Card.Root>
					<Card.Header class="pb-2">
						<Card.Title class="text-base">Share This Festival</Card.Title>
					</Card.Header>
					<Card.Content>
						<div class="flex gap-2">
							<Button variant="outline" size="sm" class="flex-1" onclick={copyLink}>
								{#if linkCopied}
									<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-1 text-green-600">
										<path d="M20 6 9 17l-5-5"/>
									</svg>
									Copied!
								{:else}
									<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-1">
										<path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/>
										<path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/>
									</svg>
									Copy Link
								{/if}
							</Button>
						</div>
					</Card.Content>
				</Card.Root>

				<!-- Quick Tips -->
				<Card.Root>
					<Card.Header class="pb-2">
						<Card.Title class="text-base flex items-center gap-2">
							<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-gold">
								<path d="M15 14c.2-1 .7-1.7 1.5-2.5 1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"/>
								<path d="M9 18h6"/>
								<path d="M10 22h4"/>
							</svg>
							First-Timer Tips
						</Card.Title>
					</Card.Header>
					<Card.Content class="text-sm text-muted-foreground space-y-2">
						<p>✓ Arrive early for the best experience</p>
						<p>✓ Bring cash for food and vendors</p>
						<p>✓ Dress comfortably and respectfully</p>
						<p>✓ Stay hydrated - T&T sun is strong!</p>
						<p>✓ Ask before photographing people</p>
					</Card.Content>
				</Card.Root>

				<!-- Newsletter Signup -->
				<NewsletterSignup variant="compact" />
			</aside>
		</div>

		<!-- Related Festivals -->
		{#if relatedFestivals.length > 0}
			<Separator class="my-12" />
			<section>
				<h2 class="text-2xl font-bold mb-6">
					More {heritageLabels[festival.heritageType]} Festivals
				</h2>
				<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
					{#each relatedFestivals as related (related.id)}
						<FestivalCard festival={related} />
					{/each}
				</div>
			</section>
		{/if}
	</main>

	<!-- Footer -->
	<footer class="border-t bg-card mt-12">
		<div class="container max-w-6xl mx-auto px-4 py-6 text-center text-sm text-muted-foreground">
			<p>Trini2DBone 2026  for Trinidad & Tobago · TrinbagoTech Hackathon 2026</p>
		</div>
	</footer>
</div>

<style>
	/* Prose styling for markdown content */
	:global(.prose) {
		color: var(--foreground);
	}
	:global(.prose h1),
	:global(.prose h2),
	:global(.prose h3),
	:global(.prose h4) {
		font-weight: 600;
		color: var(--foreground);
		margin-top: 1rem;
		margin-bottom: 0.5rem;
	}
	:global(.prose p) {
		margin-bottom: 0.75rem;
		line-height: 1.625;
	}
	:global(.prose ul) {
		list-style-type: disc;
		padding-left: 1.25rem;
		margin-bottom: 0.75rem;
	}
	:global(.prose ol) {
		list-style-type: decimal;
		padding-left: 1.25rem;
		margin-bottom: 0.75rem;
	}
	:global(.prose li) {
		line-height: 1.625;
		margin-bottom: 0.25rem;
	}
	:global(.prose strong) {
		font-weight: 600;
		color: var(--foreground);
	}
	:global(.prose a) {
		color: var(--tt-red);
	}
	:global(.prose a:hover) {
		text-decoration: underline;
	}
</style>
