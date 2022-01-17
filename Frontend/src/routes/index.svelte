<script>
  import { user } from "$lib/sessionStore";
  import { supabase } from "$lib/supabaseClient";
  import Guest from "./_guest.svelte";
  import CourseList from "./_courseList.svelte";

  user.set(supabase.auth.user());

  supabase.auth.onAuthStateChange((_, session) => {
    user.set(session.user);
  });
</script>

{#if $user}
  <CourseList />
{:else}
  <Guest />
{/if}
