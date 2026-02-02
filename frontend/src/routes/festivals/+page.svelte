<script lang="ts">
    import { Button } from '$lib/components/ui/button';
    import { Separator } from '$lib/components/ui/separator';
    import {
        CalendarGrid,
        FestivalSidebar,
        FilterBar,
        FestivalCard,
    } from '$lib/components/calendar';
    import { festivals, getThisWeekFestivals } from '$lib/data/festivals';
    import type { Festival, FestivalFilters, HeritageType, Region } from '$lib/types/festival';
    import { MONTH_NAMES } from '$lib/utils/calendar';

    // Current calendar view state
    let currentYear = $state(2026);
    let currentMonth = $state(1); // February 2026 to show Carnival

    // Filter state
    let filters = $state<FestivalFilters>({
        month: null,
        region: null,
        heritageType: null,
    });

    // Derived: check if any filters are active
    const hasActiveFilters = $derived(
        filters.month !== null || filters.region !== null || filters.heritageType !== null
    );

    // Derived: filtered festivals based on current filters
    const filteredFestivals = $derived.by(() => {
        return festivals.filter((f) => {
            // Filter by month (check if festival occurs in selected month)
            if (filters.month !== null && f.date2026Start) {
                const startMonth = new Date(f.date2026Start).getMonth();
                const endMonth = f.date2026End ? new Date(f.date2026End).getMonth() : startMonth;
                if (!(filters.month >= startMonth && filters.month <= endMonth)) {
                    return false;
                }
            }

            // Filter by region
            if (filters.region !== null && f.region !== filters.region) {
                return false;
            }

            // Filter by heritage type
            if (filters.heritageType !== null && f.heritageType !== filters.heritageType) {
                return false;
            }

            return true;
        });
    });

    // Set of filtered festival IDs for quick lookup
    const filteredFestivalIds = $derived(new Set(filteredFestivals.map((f) => f.id)));

    // Sidebar festivals (sorted by date)
    const sidebarFestivals = $derived(
        [...filteredFestivals].sort((a, b) => {
            if (!a.date2026Start) return 1;
            if (!b.date2026Start) return -1;
            return new Date(a.date2026Start).getTime() - new Date(b.date2026Start).getTime();
        })
    );

    // "This Week" festivals - using a demo date in February 2026 to show Carnival
    const thisWeekFestivals = $derived.by(() => {
        // For demo purposes, use Feb 10, 2026 as "today" to show Carnival in "upcoming"
        const demoToday = new Date('2026-02-10');
        return getThisWeekFestivals(demoToday);
    });

    // Navigation
    function prevMonth() {
        if (currentMonth === 0) {
            currentMonth = 11;
            currentYear--;
        } else {
            currentMonth--;
        }
    }

    function nextMonth() {
        if (currentMonth === 11) {
            currentMonth = 0;
            currentYear++;
        } else {
            currentMonth++;
        }
    }

    function goToToday() {
        // For demo, go to February 2026
        currentYear = 2026;
        currentMonth = 1;
    }

    // Filter handlers
    function handleMonthChange(month: number | null) {
        filters.month = month;
        // If a month is selected, navigate calendar to that month
        if (month !== null) {
            currentMonth = month;
        }
    }

    function handleRegionChange(region: Region | null) {
        filters.region = region;
    }

    function handleHeritageChange(heritage: HeritageType | null) {
        filters.heritageType = heritage;
    }

    function clearFilters() {
        filters = { month: null, region: null, heritageType: null };
    }

    // Day click handler
    function handleDayClick(date: Date, dayFestivals: Festival[]) {
        // Could open a modal or navigate to first festival
        if (dayFestivals.length === 1) {
            window.location.href = `/festivals/${dayFestivals[0].slug}`;
        }
        // For multiple festivals, could show a popover (future enhancement)
    }
</script>

<svelte:head>
	<title>Festival Calendar | KULTUR</title>
	<meta name="description" content="Discover Trinidad & Tobago's cultural festivals. Browse by month, region, or heritage type." />
</svelte:head>

<div class="min-h-screen bg-background">
	<!-- Header -->
	<header class="border-b bg-card sticky top-0 z-50">
		<div class="container max-w-6xl mx-auto px-4 py-4">
			<div class="flex items-center justify-between">
				<div>
					<a href="/" class="flex items-center gap-2 hover:opacity-80 transition-opacity">
						<img src="/logo.png" alt="" class="h-10 w-10 logo-img" loading="eager" width="40" height="40" />
						<span class="text-xl font-bold text-tt-red">KULTUR</span>
					</a>
					<p class="text-sm text-muted-foreground">Your guide to Trinidad & Tobago's cultural celebrations</p>
				</div>
				<nav class="flex items-center gap-4">
					<a href="/" class="text-sm hover:text-tt-red transition-colors">Home</a>
					<a href="/festivals" class="text-sm font-medium text-tt-red">Calendar</a>
				</nav>
			</div>
		</div>
	</header>

	<main class="container max-w-6xl mx-auto px-4 py-6">
		<!-- Page Title -->
		<div class="mb-6">
			<h1 class="text-3xl font-bold mb-2">Festival Calendar</h1>
			<p class="text-muted-foreground">
				Explore 20+ cultural festivals across Trinidad & Tobago. Filter by month, region, or heritage to find your next experience.
			</p>
		</div>

		<!-- Filters -->
		<div class="mb-6">
			<FilterBar
				selectedMonth={filters.month}
				selectedRegion={filters.region}
				selectedHeritage={filters.heritageType}
				onMonthChange={handleMonthChange}
				onRegionChange={handleRegionChange}
				onHeritageChange={handleHeritageChange}
				onClearFilters={clearFilters}
			/>
		</div>

		<!-- Main Content: Calendar + Sidebar -->
		<div class="grid grid-cols-1 lg:grid-cols-[1fr_320px] gap-6 mb-8">
			<!-- Calendar Section -->
			<div class="bg-card rounded-lg border p-4">
				<!-- Calendar Navigation -->
				<div class="flex items-center justify-between mb-4">
					<h2 class="text-xl font-semibold">
						{MONTH_NAMES[currentMonth]} {currentYear}
					</h2>
					<div class="flex items-center gap-2">
						<Button variant="outline" size="sm" onclick={goToToday}>
							Today
						</Button>
						<Button variant="outline" size="icon" onclick={prevMonth}>
							<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
								<path d="m15 18-6-6 6-6"/>
							</svg>
						</Button>
						<Button variant="outline" size="icon" onclick={nextMonth}>
							<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
								<path d="m9 18 6-6-6-6"/>
							</svg>
						</Button>
					</div>
				</div>

				<!-- Calendar Grid -->
				<CalendarGrid
					year={currentYear}
					month={currentMonth}
					{festivals}
					{filteredFestivalIds}
					{hasActiveFilters}
					onDayClick={handleDayClick}
				/>

				<!-- Legend -->
				<div class="mt-4 pt-4 border-t">
					<p class="text-xs text-muted-foreground mb-2">Heritage Types:</p>
					<div class="flex flex-wrap gap-2">
						<span class="flex items-center gap-1 text-xs">
							<span class="w-3 h-3 rounded bg-heritage-african"></span>
							African
						</span>
						<span class="flex items-center gap-1 text-xs">
							<span class="w-3 h-3 rounded bg-heritage-indian"></span>
							Indian
						</span>
						<span class="flex items-center gap-1 text-xs">
							<span class="w-3 h-3 rounded bg-heritage-indigenous"></span>
							Indigenous
						</span>
						<span class="flex items-center gap-1 text-xs">
							<span class="w-3 h-3 rounded bg-heritage-mixed"></span>
							Mixed
						</span>
						<span class="flex items-center gap-1 text-xs">
							<span class="w-3 h-3 rounded bg-heritage-christian"></span>
							Christian
						</span>
					</div>
				</div>
			</div>

			<!-- Sidebar -->
			<div class="bg-card rounded-lg border h-150 overflow-hidden">
				<FestivalSidebar
					festivals={sidebarFestivals}
					title={hasActiveFilters ? 'Matching Festivals' : 'All Festivals'}
					emptyMessage="No festivals match your current filters. Try adjusting your selection."
				/>
			</div>
		</div>

		<Separator class="my-8" />

		<!-- This Week Section -->
		<section class="mb-8">
			<div class="flex items-center justify-between mb-4">
				<div>
					<h2 class="text-2xl font-bold">This Week in T&T</h2>
					<p class="text-muted-foreground text-sm">Festivals happening in the next 2 weeks</p>
				</div>
			</div>

			{#if thisWeekFestivals.length > 0}
				<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
					{#each thisWeekFestivals as festival (festival.id)}
						<FestivalCard {festival} />
					{/each}
				</div>
			{:else}
				<div class="bg-muted/50 rounded-lg p-8 text-center">
					<p class="text-muted-foreground">No festivals scheduled for the coming weeks.</p>
					<p class="text-sm text-muted-foreground mt-1">Browse the calendar above to explore upcoming events.</p>
				</div>
			{/if}
		</section>

		<!-- All Festivals Grid (for when no filters are active) -->
		{#if !hasActiveFilters}
			<section>
				<h2 class="text-2xl font-bold mb-4">All Festivals</h2>
				<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
					{#each festivals as festival (festival.id)}
						<FestivalCard {festival} variant="compact" />
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
