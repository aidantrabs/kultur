<script lang="ts">
    import * as Card from '$lib/components/ui/card';
    import type { Memory } from '$lib/types/festival';

    const { memory }: { memory: Memory } = $props();

    function formatDate(dateString: string): string {
        return new Date(dateString).toLocaleDateString('en-TT', {
            month: 'short',
            year: 'numeric',
        });
    }

    function getInitials(name: string | null): string {
        if (!name || name === 'Anonymous') return '?';
        return name
            .split(' ')
            .map((n) => n[0])
            .join('')
            .toUpperCase()
            .slice(0, 2);
    }
</script>

<Card.Root class="h-full">
	<Card.Content class="pt-6">
		<!-- Quote icon -->
		<svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="currentColor"
			class="text-tt-red/20 mb-3"
		>
			<path
				d="M11.192 15.757c0-.88-.23-1.618-.69-2.217-.326-.412-.768-.683-1.327-.812-.55-.128-1.07-.137-1.54-.028-.16-.95.1-1.956.76-3.022.66-1.065 1.515-1.867 2.558-2.403L9.373 5c-.8.396-1.56.898-2.26 1.505-.71.607-1.34 1.305-1.9 2.094s-.98 1.68-1.25 2.69-.346 2.04-.217 3.1c.168 1.4.62 2.52 1.356 3.35.735.84 1.652 1.26 2.748 1.26.965 0 1.766-.29 2.4-.878.628-.576.94-1.365.94-2.368l.002.004zm9.124 0c0-.88-.23-1.618-.69-2.217-.326-.42-.768-.695-1.327-.825-.55-.13-1.07-.14-1.54-.03-.16-.94.09-1.95.75-3.02.66-1.06 1.514-1.86 2.557-2.4L18.49 5c-.8.396-1.555.898-2.26 1.505-.708.607-1.34 1.305-1.894 2.094-.556.79-.97 1.68-1.24 2.69-.273 1-.345 2.04-.217 3.1.165 1.4.615 2.52 1.35 3.35.732.833 1.646 1.25 2.742 1.25.967 0 1.768-.29 2.402-.876.627-.576.942-1.365.942-2.368v.012z"
			/>
		</svg>

		<!-- Content -->
		<p class="text-sm leading-relaxed text-foreground/90 mb-4">
			{memory.content}
		</p>

		<!-- Author & Meta -->
		<div class="flex items-center gap-3 pt-3 border-t">
			<!-- Avatar -->
			<div
				class="w-10 h-10 rounded-full bg-tt-red/10 flex items-center justify-center text-tt-red font-semibold text-sm"
			>
				{getInitials(memory.authorName)}
			</div>

			<div class="flex-1 min-w-0">
				<p class="font-medium text-sm truncate">
					{memory.authorName || 'Anonymous'}
				</p>
				<div class="flex items-center gap-2 text-xs text-muted-foreground">
					{#if memory.yearOfMemory}
						<span>Attended {memory.yearOfMemory}</span>
						<span>·</span>
					{/if}
					<span>Shared {formatDate(memory.submittedAt)}</span>
				</div>
			</div>
		</div>
	</Card.Content>
</Card.Root>
