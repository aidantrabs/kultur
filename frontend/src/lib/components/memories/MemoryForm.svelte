<script lang="ts">
	import { submitMemory } from '$lib/data/memories';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';

	interface Props {
		festivalId: string;
		festivalName: string;
		onSubmit?: (memory: MemorySubmission) => void;
	}

	interface MemorySubmission {
		festivalId: string;
		authorName: string;
		authorEmail: string;
		content: string;
		yearOfMemory: string;
	}

	let { festivalId, festivalName, onSubmit }: Props = $props();

	// Form state
	let authorName = $state('');
	let authorEmail = $state('');
	let content = $state('');
	let yearOfMemory = $state('');

	// UI state
	let isSubmitting = $state(false);
	let isSubmitted = $state(false);
	let error = $state<string | null>(null);

	// Generate year options (last 30 years)
	const currentYear = new Date().getFullYear();
	const yearOptions = Array.from({ length: 30 }, (_, i) => currentYear - i);

	// Character count
	const maxLength = 1000;
	const charCount = $derived(content.length);
	const isOverLimit = $derived(charCount > maxLength);

	async function handleSubmit(e: Event) {
		e.preventDefault();
		error = null;

		// Validation
		if (!content.trim()) {
			error = 'Please share your memory before submitting.';
			return;
		}

		if (isOverLimit) {
			error = `Your memory is too long. Please keep it under ${maxLength} characters.`;
			return;
		}

		isSubmitting = true;

		try {
			const submission: MemorySubmission = {
				festivalId,
				authorName: authorName.trim(),
				authorEmail: authorEmail.trim(),
				content: content.trim(),
				yearOfMemory,
			};

			// Submit using API-aware function (handles both mock and real API)
			await submitMemory(submission);

			// Call optional callback
			onSubmit?.(submission);

			// Show success state
			isSubmitted = true;

			// Reset form
			authorName = '';
			authorEmail = '';
			content = '';
			yearOfMemory = '';
		} catch (err) {
			error = 'Something went wrong. Please try again.';
			console.error('Memory submission error:', err);
		} finally {
			isSubmitting = false;
		}
	}

	function resetForm() {
		isSubmitted = false;
		error = null;
	}
</script>

<Card.Root>
	<Card.Header>
		<Card.Title class="flex items-center gap-2">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				width="20"
				height="20"
				viewBox="0 0 24 24"
				fill="none"
				stroke="currentColor"
				stroke-width="2"
				stroke-linecap="round"
				stroke-linejoin="round"
				class="text-tt-red"
			>
				<path d="M12 20h9" />
				<path d="M16.5 3.5a2.12 2.12 0 0 1 3 3L7 19l-4 1 1-4Z" />
			</svg>
			Share Your Memory
		</Card.Title>
		<Card.Description>
			Been to {festivalName}? Share your experience to help first-timers know what to expect.
		</Card.Description>
	</Card.Header>

	<Card.Content>
		{#if isSubmitted}
			<!-- Success State -->
			<div class="text-center py-6">
				<div
					class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						width="32"
						height="32"
						viewBox="0 0 24 24"
						fill="none"
						stroke="currentColor"
						stroke-width="2"
						stroke-linecap="round"
						stroke-linejoin="round"
						class="text-green-600"
					>
						<path d="M22 11.08V12a10 10 0 1 1-5.93-9.14" />
						<path d="m9 11 3 3L22 4" />
					</svg>
				</div>
				<h3 class="font-semibold text-lg mb-2">Thank You!</h3>
				<p class="text-muted-foreground text-sm mb-4">
					Your memory has been submitted and will appear after review. We appreciate you sharing
					your experience!
				</p>
				<Button variant="outline" onclick={resetForm}>Share Another Memory</Button>
			</div>
		{:else}
			<!-- Form -->
			<form onsubmit={handleSubmit} class="space-y-4">
				{#if error}
					<div class="bg-destructive/10 text-destructive text-sm p-3 rounded-md">
						{error}
					</div>
				{/if}

				<!-- Memory Content -->
				<div class="space-y-2">
					<label for="content" class="text-sm font-medium">
						Your Memory <span class="text-destructive">*</span>
					</label>
					<textarea
						id="content"
						bind:value={content}
						placeholder="What was your experience like? What surprised you? What advice would you give first-timers?"
						rows="5"
						class="w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 resize-none"
						disabled={isSubmitting}
					></textarea>
					<div class="flex justify-end">
						<span
							class="text-xs {isOverLimit
								? 'text-destructive'
								: charCount > maxLength * 0.9
									? 'text-yellow-600'
									: 'text-muted-foreground'}"
						>
							{charCount}/{maxLength}
						</span>
					</div>
				</div>

				<div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
					<!-- Name -->
					<div class="space-y-2">
						<label for="authorName" class="text-sm font-medium">Your Name</label>
						<input
							type="text"
							id="authorName"
							bind:value={authorName}
							placeholder="Optional (or stay anonymous)"
							class="w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
							disabled={isSubmitting}
						/>
					</div>

					<!-- Year Attended -->
					<div class="space-y-2">
						<label for="yearOfMemory" class="text-sm font-medium">Year Attended</label>
						<select
							id="yearOfMemory"
							bind:value={yearOfMemory}
							class="w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
							disabled={isSubmitting}
						>
							<option value="">Select year (optional)</option>
							{#each yearOptions as year}
								<option value={year.toString()}>{year}</option>
							{/each}
						</select>
					</div>
				</div>

				<!-- Email (for moderation contact if needed) -->
				<div class="space-y-2">
					<label for="authorEmail" class="text-sm font-medium">Email</label>
					<input
						type="email"
						id="authorEmail"
						bind:value={authorEmail}
						placeholder="Optional — only used if we need to contact you"
						class="w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
						disabled={isSubmitting}
					/>
					<p class="text-xs text-muted-foreground">Never displayed publicly</p>
				</div>

				<!-- Submit -->
				<Button type="submit" class="w-full bg-tt-red hover:bg-tt-red-dark" disabled={isSubmitting}>
					{#if isSubmitting}
						<svg
							class="animate-spin -ml-1 mr-2 h-4 w-4"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
						>
							<circle
								class="opacity-25"
								cx="12"
								cy="12"
								r="10"
								stroke="currentColor"
								stroke-width="4"
							></circle>
							<path
								class="opacity-75"
								fill="currentColor"
								d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
							></path>
						</svg>
						Submitting...
					{:else}
						Submit Memory
					{/if}
				</Button>

				<p class="text-xs text-muted-foreground text-center">
					Memories are reviewed before appearing. Please be respectful and authentic.
				</p>
			</form>
		{/if}
	</Card.Content>
</Card.Root>
