<script>
  import {user} from "$lib/sessionStore"
  import {supabase} from "$lib/supabaseClient"
  import Guest from "./_Guest.svelte"
  import Dashboard from "./_dashboard.svelte"

  user.set(supabase.auth.user())

  supabase.auth.onAuthStateChange((_, session) => {
      user.set(session.user)
  })
</script>

{#if $user}
  <Dashboard />
{:else}
  <Guest />
{/if}