<script>
  import { user } from "./sessionStore";
  import { supabase } from "./supabaseClient";
  import Auth from "./Auth.svelte";
  import Profile from "./Profile.svelte";

  user.set(supabase.auth.user());

  supabase.auth.onAuthStateChange((_, session) => {
    user.set(session.user);
  });

  console.log(supabase.auth.session().access_token);

  async function getData() {
    const response = await fetch("http://localhost:8080/test", {
      headers: {
        Authorization: `Bearer ${supabase.auth.session().access_token}`,
      },
    });
    const data = await response.text();

    console.log(data);
  }

  getData();
</script>

<div class="container" style="padding: 50px 0 100px 0;">
  {#if $user}
    <Profile />
  {:else}
    <Auth />
  {/if}
</div>
