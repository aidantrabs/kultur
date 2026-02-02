<script lang="ts">
    import { Button } from '$lib/components/ui/button';
    import { config } from '$lib/config';
    import { subscriptionsApi } from '$lib/api';

    interface Props {
        variant?: 'default' | 'compact' | 'hero';
    }

    const { variant = 'default' }: Props = $props();

    let email = $state('');

    let isSubmitting = $state(false);
    let showSuccess = $state(false);
    let error = $state<string | null>(null);

    async function handleSubmit(e: Event) {
        e.preventDefault();
        console.log('handleSubmit called!', { email });
        error = null;

        if (!email.trim()) {
            error = 'Please enter your email address.';
            console.log('Validation failed: empty email');
            return;
        }

        if (!email.includes('@') || !email.includes('.')) {
            error = 'Please enter a valid email address.';
            console.log('Validation failed: invalid email');
            return;
        }

        console.log('Validation passed, submitting...');
        isSubmitting = true;

        try {
            console.log('Subscribing with config:', {
                useApi: config.useApi,
                apiUrl: config.apiUrl,
            });

            if (config.useApi) {
                const result = await subscriptionsApi.subscribe({
                    email: email.trim(),
                });
                console.log('Subscription result:', result);
            } else {
                console.log('Using mock mode');
                await new Promise((resolve) => setTimeout(resolve, 1000));
            }

            showSuccess = true;
            email = '';

            setTimeout(() => {
                showSuccess = false;
            }, 3000);
        } catch (err: unknown) {
            console.error('Subscription error:', err);
            error = err instanceof Error ? err.message : 'Something went wrong. Please try again.';
        } finally {
            isSubmitting = false;
        }
    }
</script>

{#if variant === 'hero'}
	<section class="bg-tt-black text-white py-16">
		<div class="container max-w-6xl mx-auto px-4">
			<div class="grid md:grid-cols-2 gap-8 items-center">
				<div>
					<h2 class="text-3xl lg:text-4xl font-bold mb-4">
						Never Miss a <span class="text-tt-gold">Festival</span>
					</h2>
					<p class="text-gray-300 text-lg mb-4">
						Get notified when festivals are coming up, with first-timer guides and practical tips.
					</p>
					<div class="flex flex-wrap gap-4 text-sm text-gray-400">
						<span class="flex items-center gap-1.5">
							<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-gold">
								<path d="M6 8a6 6 0 0 1 12 0c0 7 3 9 3 9H3s3-2 3-9"/>
								<path d="M10.3 21a1.94 1.94 0 0 0 3.4 0"/>
							</svg>
							Festival alerts
						</span>
						<span class="flex items-center gap-1.5">
							<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-gold">
								<path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/>
							</svg>
							First-timer guides
						</span>
						<span class="flex items-center gap-1.5">
							<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-gold">
								<rect width="18" height="18" x="3" y="4" rx="2" ry="2"/>
								<line x1="16" x2="16" y1="2" y2="6"/>
								<line x1="8" x2="8" y1="2" y2="6"/>
								<line x1="3" x2="21" y1="10" y2="10"/>
							</svg>
							Weekly digest
						</span>
					</div>
				</div>
				<div>
					<form onsubmit={handleSubmit} class="space-y-4">
						{#if error}
							<div class="bg-red-500/20 text-red-300 text-sm p-3 rounded-md">
								{error}
							</div>
						{/if}

						<div class="flex flex-col sm:flex-row gap-3">
							<input
								type="email"
								bind:value={email}
								placeholder={showSuccess ? 'Subscribed!' : 'Enter your email'}
								class="flex-1 rounded-md border px-4 py-3 placeholder:text-gray-400 focus:outline-none focus:ring-2 focus:ring-tt-red focus:border-transparent transition-colors duration-300
									{showSuccess ? 'border-green-500/50 bg-green-500/10 text-green-300' : 'border-white/20 bg-white/10 text-white'}"
								disabled={isSubmitting}
							/>
							<Button
								type="submit"
								size="lg"
								class="shrink-0 transition-all duration-300 {showSuccess ? 'bg-green-600 hover:bg-green-700' : 'bg-tt-red hover:bg-tt-red-dark'} text-white"
								disabled={isSubmitting || showSuccess}
							>
								{#if isSubmitting}
									<svg class="animate-spin -ml-1 mr-2 h-4 w-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
										<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
										<path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
									</svg>
									Subscribing...
								{:else if showSuccess}
									<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-1">
										<path d="M20 6 9 17l-5-5"/>
									</svg>
									Subscribed
								{:else}
									Subscribe
								{/if}
							</Button>
						</div>

						<p class="text-xs text-gray-400">
							We respect your privacy. Unsubscribe anytime.
						</p>
					</form>
				</div>
			</div>
		</div>
	</section>

{:else if variant === 'compact'}
	<div class="bg-card border rounded-lg p-6">
		<h3 class="font-semibold mb-2 flex items-center gap-2">
			<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-tt-red">
				<rect width="20" height="16" x="2" y="4" rx="2"/>
				<path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/>
			</svg>
			Stay Updated
		</h3>
		<p class="text-sm text-muted-foreground mb-4">Get festival reminders and updates.</p>

		<form onsubmit={handleSubmit} class="space-y-3">
			{#if error}
				<p class="text-sm text-destructive">{error}</p>
			{/if}

			<input
				type="email"
				bind:value={email}
				placeholder={showSuccess ? 'Subscribed!' : 'Your email'}
				class="w-full rounded-md border bg-background px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-ring transition-colors duration-300
					{showSuccess ? 'border-green-500 bg-green-50' : 'border-input'}"
				disabled={isSubmitting}
			/>
			<Button
				type="submit"
				class="w-full transition-all duration-300 {showSuccess ? 'bg-green-600 hover:bg-green-700' : 'bg-tt-red hover:bg-tt-red-dark'}"
				disabled={isSubmitting || showSuccess}
			>
				{#if isSubmitting}
					Subscribing...
				{:else if showSuccess}
					<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-1">
						<path d="M20 6 9 17l-5-5"/>
					</svg>
					Subscribed
				{:else}
					Subscribe
				{/if}
			</Button>
		</form>
	</div>

{:else}
	<div class="bg-muted/50 rounded-lg p-6 md:p-8">
		<div class="flex flex-col md:flex-row md:items-center gap-6">
			<div class="flex-1">
				<h3 class="font-semibold text-lg mb-1">Get Festival Updates</h3>
				<p class="text-muted-foreground text-sm">
					Weekly digest of upcoming festivals, tips, and cultural insights.
				</p>
			</div>
			<form onsubmit={handleSubmit} class="flex flex-col sm:flex-row gap-2 md:w-auto w-full">
				{#if error}
					<p class="text-sm text-destructive sm:hidden">{error}</p>
				{/if}
				<input
					type="email"
					bind:value={email}
					placeholder={showSuccess ? 'Subscribed!' : 'Enter your email'}
					class="rounded-md border bg-background px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-ring min-w-[240px] transition-colors duration-300
						{showSuccess ? 'border-green-500 bg-green-50' : 'border-input'}"
					disabled={isSubmitting}
				/>
				<Button
					type="submit"
					class="shrink-0 transition-all duration-300 {showSuccess ? 'bg-green-600 hover:bg-green-700' : 'bg-tt-red hover:bg-tt-red-dark'}"
					disabled={isSubmitting || showSuccess}
				>
					{#if isSubmitting}
						Subscribing...
					{:else if showSuccess}
						<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-1">
							<path d="M20 6 9 17l-5-5"/>
						</svg>
						Subscribed
					{:else}
						Subscribe
					{/if}
				</Button>
			</form>
		</div>
		{#if error}
			<p class="text-sm text-destructive mt-2 hidden sm:block">{error}</p>
		{/if}
	</div>
{/if}
