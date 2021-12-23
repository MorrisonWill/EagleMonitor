<script context="module">
	export function load({ error, status }) {
		return {
			props: { error, status },
		};
	}
</script>

<script>
	import { dev } from "$app/env";
	export let error, status;
</script>

<svelte:head>
	<title>{status}</title>
</svelte:head>

<section class="py-20"><div class="container px-4 mx-auto text-center">
	<img class="mb-4 mx-auto md:max-w-md mb-12" src="metis-assets/illustrations/error2.png" alt=""><span class="text-4xl text-blue-600 font-bold font-heading">Error {status}</span>
	<h2 class="mb-2 text-4xl font-bold font-heading">{error.message}</h2>
	<p class="mb-6 text-blueGray-400">Sorry! We encountered an error. You can try again or go back to
		the homepage with the button below. If this doesn't seem normal, please contact me at the email listed in the footer.</p>
	<div><a class="block sm:inline-block py-4 px-8 mb-4 sm:mb-0 sm:mr-3 text-xs text-white text-center font-semibold leading-none bg-blue-600 hover:bg-blue-700 rounded" href="#">Go back to Homepage</a><a class="block sm:inline-block py-4 px-8 text-xs text-blueGray-500 hover:text-blueGray-800 text-center font-semibold leading-none bg-blueGray-50 hover:bg-blueGray-100 rounded" href="#" data-removed="true">Try Again</a></div>
	</div>
</section>
{#if dev && error.stack}
	<pre>{error.stack}</pre>
{/if}