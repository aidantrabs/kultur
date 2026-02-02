<script lang="ts">
    import * as Select from '$lib/components/ui/select';
    import { Button } from '$lib/components/ui/button';
    import type { HeritageType, Region } from '$lib/types/festival';
    import { heritageLabels, regionLabels } from '$lib/types/festival';
    import { MONTH_NAMES } from '$lib/utils/calendar';

    interface Props {
        selectedMonth: number | null;
        selectedRegion: Region | null;
        selectedHeritage: HeritageType | null;
        onMonthChange: (month: number | null) => void;
        onRegionChange: (region: Region | null) => void;
        onHeritageChange: (heritage: HeritageType | null) => void;
        onClearFilters: () => void;
    }

    const {
        selectedMonth,
        selectedRegion,
        selectedHeritage,
        onMonthChange,
        onRegionChange,
        onHeritageChange,
        onClearFilters,
    }: Props = $props();

    const hasActiveFilters = $derived(
        selectedMonth !== null || selectedRegion !== null || selectedHeritage !== null
    );

    const regions: Region[] = ['north', 'south', 'central', 'east', 'west', 'tobago', 'nationwide'];
    const heritageTypes: HeritageType[] = ['african', 'indian', 'indigenous', 'mixed', 'christian'];
</script>

<div class="flex flex-wrap items-center gap-3">
	<!-- Month Filter -->
	<Select.Root
		type="single"
		value={selectedMonth !== null ? String(selectedMonth) : undefined}
		onValueChange={(v) => onMonthChange(v ? parseInt(v) : null)}
	>
		<Select.Trigger class="w-40">
			{selectedMonth !== null ? MONTH_NAMES[selectedMonth] : 'All Months'}
		</Select.Trigger>
		<Select.Content>
			<Select.Item value="">All Months</Select.Item>
			{#each MONTH_NAMES as monthName, index}
				<Select.Item value={String(index)}>{monthName}</Select.Item>
			{/each}
		</Select.Content>
	</Select.Root>

	<!-- Region Filter -->
	<Select.Root
		type="single"
		value={selectedRegion ?? undefined}
		onValueChange={(v) => onRegionChange(v as Region | null)}
	>
		<Select.Trigger class="w-40">
			{selectedRegion ? regionLabels[selectedRegion] : 'All Regions'}
		</Select.Trigger>
		<Select.Content>
			<Select.Item value="">All Regions</Select.Item>
			{#each regions as region}
				<Select.Item value={region}>{regionLabels[region]}</Select.Item>
			{/each}
		</Select.Content>
	</Select.Root>

	<!-- Heritage Filter -->
	<Select.Root
		type="single"
		value={selectedHeritage ?? undefined}
		onValueChange={(v) => onHeritageChange(v as HeritageType | null)}
	>
		<Select.Trigger class="w-45">
			{selectedHeritage ? heritageLabels[selectedHeritage] : 'All Heritage Types'}
		</Select.Trigger>
		<Select.Content>
			<Select.Item value="">All Heritage Types</Select.Item>
			{#each heritageTypes as heritage}
				<Select.Item value={heritage}>{heritageLabels[heritage]}</Select.Item>
			{/each}
		</Select.Content>
	</Select.Root>

	<!-- Clear Filters -->
	{#if hasActiveFilters}
		<Button variant="ghost" size="sm" onclick={onClearFilters}>
			<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-1">
				<path d="M18 6 6 18"/>
				<path d="m6 6 12 12"/>
			</svg>
			Clear Filters
		</Button>
	{/if}
</div>
