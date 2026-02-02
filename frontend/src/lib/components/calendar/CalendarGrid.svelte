<script lang="ts">
    import { Badge } from '$lib/components/ui/badge';
    import type { Festival, HeritageType } from '$lib/types/festival';
    import { heritageLabels } from '$lib/types/festival';
    import { getCalendarDays, DAY_NAMES, isDateInRange } from '$lib/utils/calendar';

    interface Props {
        year: number;
        month: number;
        festivals: Festival[];
        filteredFestivalIds: Set<string>;
        hasActiveFilters: boolean;
        onDayClick?: (date: Date, festivals: Festival[]) => void;
    }

    const { year, month, festivals, filteredFestivalIds, hasActiveFilters, onDayClick }: Props =
        $props();

    const days = $derived(getCalendarDays(year, month));

    // Get festivals for a specific day
    function getFestivalsForDay(date: Date): Festival[] {
        return festivals.filter((f) => isDateInRange(date, f.date2026Start, f.date2026End));
    }

    // Check if a festival matches current filters
    function isHighlighted(festival: Festival): boolean {
        if (!hasActiveFilters) return true;
        return filteredFestivalIds.has(festival.id);
    }

    // Get heritage color class
    function getHeritageColorClass(heritage: HeritageType): string {
        const colors: Record<HeritageType, string> = {
            african: 'bg-heritage-african',
            indian: 'bg-heritage-indian',
            indigenous: 'bg-heritage-indigenous',
            mixed: 'bg-heritage-mixed',
            christian: 'bg-heritage-christian',
        };
        return colors[heritage];
    }

    function handleDayClick(date: Date) {
        const dayFestivals = getFestivalsForDay(date);
        if (dayFestivals.length > 0 && onDayClick) {
            onDayClick(date, dayFestivals);
        }
    }
</script>

<div class="w-full">
	<!-- Day headers -->
	<div class="grid grid-cols-7 gap-1 mb-2">
		{#each DAY_NAMES as dayName}
			<div class="text-center text-sm font-medium text-muted-foreground py-2">
				{dayName}
			</div>
		{/each}
	</div>

	<!-- Calendar grid -->
	<div class="grid grid-cols-7 gap-1">
		{#each days as day}
			{@const dayFestivals = getFestivalsForDay(day.date)}
			{@const hasFestivals = dayFestivals.length > 0}
			{@const hasHighlightedFestivals = dayFestivals.some((f) => isHighlighted(f))}
			
			<button
				type="button"
				onclick={() => handleDayClick(day.date)}
				class="
					relative min-h-20 p-1 rounded-md border transition-all duration-200
					{day.isCurrentMonth ? 'bg-card' : 'bg-muted/30'}
					{day.isToday ? 'ring-2 ring-tt-red ring-offset-1' : ''}
					{hasFestivals ? 'cursor-pointer hover:border-tt-red/50' : 'cursor-default'}
					{hasFestivals && hasActiveFilters && !hasHighlightedFestivals ? 'opacity-40' : ''}
				"
				disabled={!hasFestivals}
			>
				<!-- Day number -->
				<span
					class="
						absolute top-1 left-2 text-sm font-medium
						{day.isCurrentMonth ? 'text-foreground' : 'text-muted-foreground'}
						{day.isToday ? 'text-tt-red font-bold' : ''}
					"
				>
					{day.day}
				</span>

				<!-- Festival indicators -->
				{#if hasFestivals}
					<div class="mt-5 space-y-1">
						{#each dayFestivals.slice(0, 2) as festival}
							{@const highlighted = isHighlighted(festival)}
							<div
								class="
									text-[10px] leading-tight px-1 py-0.5 rounded truncate
									transition-all duration-200
									{getHeritageColorClass(festival.heritageType)}
									{highlighted ? 'opacity-100 text-white' : 'opacity-30'}
								"
								title={festival.name}
							>
								{festival.name}
							</div>
						{/each}
						{#if dayFestivals.length > 2}
							<div class="text-[10px] text-muted-foreground text-center">
								+{dayFestivals.length - 2} more
							</div>
						{/if}
					</div>
				{/if}
			</button>
		{/each}
	</div>
</div>
