<script lang="ts">
    import * as Card from '$lib/components/ui/card';
    import { Badge } from '$lib/components/ui/badge';
    import type { Festival, HeritageType } from '$lib/types/festival';
    import { heritageLabels, regionLabels } from '$lib/types/festival';
    import { formatDateRange, getRelativeTime } from '$lib/utils/calendar';

    interface Props {
        festival: Festival;
        variant?: 'default' | 'compact';
    }

    const { festival, variant = 'default' }: Props = $props();

    function getHeritageBadgeClass(heritage: HeritageType): string {
        const classes: Record<HeritageType, string> = {
            african: 'bg-heritage-african text-white hover:bg-heritage-african/80',
            indian: 'bg-heritage-indian text-black hover:bg-heritage-indian/80',
            indigenous: 'bg-heritage-indigenous text-white hover:bg-heritage-indigenous/80',
            mixed: 'bg-heritage-mixed text-white hover:bg-heritage-mixed/80',
            christian: 'bg-heritage-christian text-white hover:bg-heritage-christian/80',
        };
        return classes[heritage];
    }
</script>

<a href="/festivals/{festival.slug}" class="block group h-full w-full">
	<Card.Root class="h-full overflow-hidden transition-all duration-300 hover:shadow-lg hover:border-tt-red/50 hover:-translate-y-1">
		{#if festival.coverImageUrl && variant === 'default'}
			<div class="relative h-40 overflow-hidden">
				<img
					src={festival.coverImageUrl}
					alt={festival.name}
					class="w-full h-full object-cover transition-transform duration-300 group-hover:scale-105"
				/>
				<div class="absolute inset-0 bg-linear-to-t from-black/60 to-transparent"></div>
				<div class="absolute bottom-2 left-2 right-2">
					<Badge class="{getHeritageBadgeClass(festival.heritageType)} text-xs">
						{heritageLabels[festival.heritageType]}
					</Badge>
				</div>
				{#if festival.date2026Start}
					<div class="absolute top-2 right-2">
						<Badge variant="secondary" class="bg-tt-red text-white text-xs font-semibold">
							{getRelativeTime(festival.date2026Start)}
						</Badge>
					</div>
				{/if}
			</div>
		{/if}

		<Card.Header class="{variant === 'compact' ? 'p-3' : 'p-4'} pb-2">
			<div class="flex items-start justify-between gap-2 {variant === 'compact' ? 'min-h-[2.5rem]' : 'min-h-[3rem]'}">
				<Card.Title class="{variant === 'compact' ? 'text-base' : 'text-lg'} font-semibold group-hover:text-tt-red transition-colors line-clamp-2">
					{festival.name}
				</Card.Title>
				{#if variant === 'compact' || !festival.coverImageUrl}
					<Badge class="{getHeritageBadgeClass(festival.heritageType)} text-xs shrink-0">
						{heritageLabels[festival.heritageType].replace(' Heritage', '').replace('/First Peoples', '')}
					</Badge>
				{/if}
			</div>
		</Card.Header>

		<Card.Content class="{variant === 'compact' ? 'p-3' : 'p-4'} pt-0 flex-grow flex flex-col">
			<p class="text-sm text-muted-foreground line-clamp-2 mb-3 {variant === 'compact' ? 'min-h-[2.5rem]' : 'min-h-[2.75rem]'}">
				{festival.summary}
			</p>

			<div class="flex flex-wrap items-center gap-3 text-xs text-muted-foreground mt-auto">
				<span class="flex items-center gap-1">
					<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
						<rect width="18" height="18" x="3" y="4" rx="2" ry="2"/>
						<line x1="16" x2="16" y1="2" y2="6"/>
						<line x1="8" x2="8" y1="2" y2="6"/>
						<line x1="3" x2="21" y1="10" y2="10"/>
					</svg>
					{formatDateRange(festival.date2026Start, festival.date2026End)}
				</span>
				<span class="flex items-center gap-1">
					<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
						<path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"/>
						<circle cx="12" cy="10" r="3"/>
					</svg>
					{regionLabels[festival.region]}
				</span>
			</div>
		</Card.Content>
	</Card.Root>
</a>
