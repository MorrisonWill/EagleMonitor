<script>
  import { onMount } from 'svelte';
  import { scale } from 'svelte/transition'; 
  import { supabase } from "$lib/supabaseClient";

  export let user;

  let show = false; // menu state
  let menu = null; // menu wrapper DOM reference
    
  async function SignOut(){
      try {
        let { error } = await supabase.auth.signOut();       
        if (error) throw error;
      } catch (error) {
        alert(error.message);
      } finally {
        window.location.href = "http://localhost:3000";  
      }
  }

  onMount(() => {
    const handleOutsideClick = (event) => {
      if (show && !menu.contains(event.target)) {
        show = false;
      }
    };

    const handleEscape = (event) => {
      if (show && event.key === 'Escape') {
        show = false;
      }
    };

    // add events when element is added to the DOM
    document.addEventListener('click', handleOutsideClick, false);
    document.addEventListener('keyup', handleEscape, false);

    // remove events when element is removed from the DOM
    return () => {
      document.removeEventListener('click', handleOutsideClick, false);
      document.removeEventListener('keyup', handleEscape, false);
    };
  });
</script>

<div class="relative" bind:this={menu}>
  <div>    
    <button 
      on:click={() => (show = !show)}
      class="navbar-burger flex items-center py-2 px-3 text-blue-600 hover:text-blue-700 rounded border border-blue-200 hover:border-blue-300"
      >
        <svg
          class="fill-current h-4 w-4"
          viewBox="0 0 20 20"
          xmlns="http://www.w3.org/2000/svg"
          ><title>Mobile menu</title><path
             d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z"
          /></svg
     ></button>
    {#if show}
      <div
        in:scale={{ duration: 100, start: 0.95 }}
        out:scale={{ duration: 75, start: 0.95 }}
        class="origin-top-right absolute right-0 w-48 py-2 mt-1 bg-blue-600
          rounded shadow-md"
      >
        <a
          href="/allcourses"
          class="block px-4 py-2 hover:bg-black-500 hover:text-green-100"
        >All Courses</a>
      {#if user}
        <a 
          href="/mycourses"
          class="block px-4 py-2 hover:bg-black-500 hover:text-green-100"
        >My Courses</a>
      {/if}
      {#if !user}
        <a
          href=""
          class="block px-4 py-2 hover:bg-black-500 hover:text-green-100"
        >Sign In</a>
      {/if}
      {#if user}
        <a
          on:click={SignOut}
          class="block px-4 py-2 hover:bg-black-500 hover:text-green-100"
        >Logout</a> 
      {/if}
      </div>
    {/if}
  </div>
</div>
