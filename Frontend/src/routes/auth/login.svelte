<script>
  import { supabase } from "$lib/supabaseClient";

  let loading = false;
  let emailSent = false;
  let error = "";
  let email;

  const handleLogin = async () => {
    try {
      if (!email.includes("bc.edu")) {
        error = "only bc.edu email addresses may be used";
      } else {
        loading = true;
        const { err } = await supabase.auth.signIn({ email });
        if (err) throw err;
        error = "";
        emailSent = true;
      }
    } catch (e) {
      error = e.error_description || e.message;
    } finally {
      loading = false;
    }
  };
</script>

<section class="py-12">
  <div class="container px-4 mx-auto">
    <form on:submit|preventDefault={handleLogin}>
      <div class="flex flex-col max-w-sm mx-auto text-center">
        <div class="my-12">
          <h4 class="text-red-700">
            {#if error}
              {error}
            {/if}
          </h4>
          <h4 class="mb-6 text-3xl">
            {#if loading}
              Loading... please wait
            {:else if emailSent}
              Please check your email for the login link
            {:else}
              Make your life easier
            {/if}
          </h4>
          <div class="flex mb-4 px-4 bg-blueGray-50 rounded">
            <input
              class="w-full py-4 text-xs placeholder-blueGray-400 font-semibold leading-none bg-blueGray-50 outline-none"
              type="email"
              placeholder="email@bc.edu"
              bind:value={email}
            /><svg
              class="h-6 w-6 ml-4 my-auto text-blueGray-300"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              ><path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207"
              /></svg
            >
          </div>
          <button
            class="block w-full p-4 text-center text-xs text-white font-semibold leading-none bg-blue-600 hover:bg-blue-700 rounded"
            type="submit"
            disabled={loading}
            >{loading ? "Loading" : "Log In / Sign Up"}</button
          >
        </div>
      </div>
    </form>
  </div>
</section>
