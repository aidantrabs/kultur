<script lang="ts">
    import * as Card from '$lib/components/ui/card';
    import { Badge } from '$lib/components/ui/badge';
    import { ScrollArea } from '$lib/components/ui/scroll-area';
    import { Separator } from '$lib/components/ui/separator';
    import type { Festival, HeritageType } from '$lib/types/festival';
    import { heritageLabels, regionLabels } from '$lib/types/festival';
    import { formatDateRange, getRelativeTime } from '$lib/utils/calendar';

    interface Props {
        festivals: Festival[];
        title?: string;
        emptyMessage?: string;
    }

    const {
        festivals,
        title = 'Festivals',
        emptyMessage = 'No festivals match your filters',
    }: Props = $props();

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

<div class="flex flex-col h-full overflow-hidden">
	<h3 class="font-semibold text-lg px-4 py-3 border-b shrink-0">
		{title}
		{#if festivals.length > 0}
			<span class="text-muted-foreground font-normal">({festivals.length})</span>
		{/if}
	</h3>

	<ScrollArea class="flex-1 min-h-0">
		<div class="p-4 space-y-3">
			{#if festivals.length === 0}
				<p class="text-muted-foreground text-sm text-center py-8">
					{emptyMessage}
				</p>
			{:else}
				{#each festivals as festival (festival.id)}
					<a
						href="/festivals/{festival.slug}"
						class="block group"
					>
						<Card.Root class="transition-all duration-200 hover:border-tt-red/50 hover:shadow-md">
							<Card.Header class="p-3 pb-2">
								<div class="flex items-start justify-between gap-2">
									<Card.Title class="text-sm font-medium leading-tight group-hover:text-tt-red transition-colors">
										{festival.name}
									</Card.Title>
									<Badge 
										variant="secondary" 
										class="text-[10px] shrink-0 {getHeritageBadgeClass(festival.heritageType)}"
									>
										{heritageLabels[festival.heritageType].replace(' Heritage', '').replace('/First Peoples', '')}
									</Badge>
								</div>
							</Card.Header>
							<Card.Content class="p-3 pt-0">
								<div class="space-y-1.5">
									<p class="text-xs text-muted-foreground flex items-center gap-1.5">
										<svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
											<rect width="18" height="18" x="3" y="4" rx="2" ry="2"/>
											<line x1="16" x2="16" y1="2" y2="6"/>
											<line x1="8" x2="8" y1="2" y2="6"/>
											<line x1="3" x2="21" y1="10" y2="10"/>
										</svg>
										{formatDateRange(festival.date2026Start, festival.date2026End)}
									</p>
									<p class="text-xs text-muted-foreground flex items-center gap-1.5">
										<svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
											<path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"/>
											<circle cx="12" cy="10" r="3"/>
										</svg>
										{regionLabels[festival.region]}
									</p>
									{#if festival.date2026Start}
										<p class="text-xs font-medium text-tt-red">
											{getRelativeTime(festival.date2026Start)}
										</p>
									{/if}
								</div>
							</Card.Content>
						</Card.Root>
					</a>
				{/each}
			{/if}
		</div>
	</ScrollArea>
</div>
