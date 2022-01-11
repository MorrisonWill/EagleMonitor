<script>
  import Search from "svelte-search";
  let value = "";  
</script>

<Search label="searchBox" hideLabel bind:value on:submit={() => console.log(value)} placeholder="Search..."/>

{#if value} 
  <button on:click={() => (value = "")}>Clear "{value}"</button>
{/if}


<style>
  :global([data-svelte-search] input) {
    width: 95%;
    font-size: 1rem;
    padding: 0.5rem;
    margin-right: 1em;
    margin-left: 1em;
    border: 1px solid #e0e0e0;
    border-radius: 0.25rem;
  }
</style>


