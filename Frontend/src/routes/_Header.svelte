<script>
  import { user } from "$lib/sessionStore";
  import { supabase } from "$lib/supabaseClient";
  import Menu from "./_Menu.svelte";

  async function signOut() {
    try {
      let { error } = await supabase.auth.signOut();
      if (error) throw error;
    } catch (error) {
      alert(error.message);
    } finally {
      window.location.href = "http://localhost:3000";
    }
  }
</script>

<section>
  <div class="container px-4 mx-auto">
    <nav class="flex justify-between items-center py-6">
      <a class="text-3xl font-semibold leading-none" href="#">
        <img
          class="h-10"
          src="https://shuffle.dev/metis-assets/logos/metis/metis.svg"
          alt=""
          width="auto"
        /></a
      >
      <div class="lg:hidden">
        <Menu />
      </div>
      {#if $user}

        <button
          class="inline-block px-4 py-3 text-xs font-semibold leading-none bg-blue-600 hover:bg-blue-700 text-white rounded"
          on:click={signOut}>Sign Out</button
        >
      {:else}
        <div class="hidden lg:block">
          <a
            class="mr-2 inline-block px-4 py-3 text-xs text-blue-600 hover:text-blue-700 font-semibold leading-none border border-blue-200 hover:border-blue-300 rounded"
            href="auth/login">Log In</a
          ><a
            class="inline-block px-4 py-3 text-xs font-semibold leading-none bg-blue-600 hover:bg-blue-700 text-white rounded"
            href="auth/login">Sign Up</a
          >
        </div>
      {/if}
    </nav>
  </div>
  <div
    class="navbar-menu fixed top-0 left-0 bottom-0 w-5/6 max-w-sm z-50 hidden"
  >
    <div class="navbar-backdrop fixed inset-0 bg-blueGray-800 opacity-25" />
    <nav
      class="relative flex flex-col py-6 px-6 w-full h-full bg-white border-r overflow-y-auto"
    >
      <div class="flex items-center mb-8">
        <a class="mr-auto text-3xl font-semibold leading-none" href="#">
          <img
            class="h-10"
            src="https://shuffle.dev/metis-assets/logos/metis/metis.svg"
            alt=""
            width="auto"
          /></a
        >
        <button class="navbar-close">
          <svg
            class="h-6 w-6 text-blueGray-400 cursor-pointer hover:text-blueGray-500"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            ><path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            /></svg
          ></button
        >
      </div>
      <div>
        {#if $user}
          <button
            class="block px-4 py-3 mb-2 text-xs text-center text-blue-600 hover:text-blue-700 font-semibold leading-none border border-blue-200 hover:border-blue-300 rounded"
            on:click={signOut}>Sign Out</button
          >
        {:else}
          <div class="mt-4 pt-6 border-t border-blueGray-100">
            <a
              class="block px-4 py-3 mb-3 text-xs text-center font-semibold leading-none bg-blue-600 hover:bg-blue-700 text-white rounded"
              href="auth/login">Sign Up</a
            ><a
              class="block px-4 py-3 mb-2 text-xs text-center text-blue-600 hover:text-blue-700 font-semibold leading-none border border-blue-200 hover:border-blue-300 rounded"
              href="auth/login">Log In</a
            >
          </div>
        {/if}
      </div>
      <div class="mt-auto" data-removed="true">
        <p class="my-4 text-xs text-blueGray-400">
          <span>Get in Touch</span>
          <a
            class="text-blue-600 hover:text-blue-600 underline"
            href="#"
            data-removed="true">info@example.com</a
          >
        </p>
        <a class="inline-block px-1" href="#">
          <img
            src="https://shuffle.dev/metis-assets/icons/facebook-blue.svg"
            alt=""
          /></a
        >
        <a class="inline-block px-1" href="#">
          <img
            src="https://shuffle.dev/metis-assets/icons/twitter-blue.svg"
            alt=""
          /></a
        >
        <a class="inline-block px-1" href="#">
          <img
            src="https://shuffle.dev/metis-assets/icons/instagram-blue.svg"
            alt=""
          /></a
        >
      </div>
    </nav>
  </div>
</section>
