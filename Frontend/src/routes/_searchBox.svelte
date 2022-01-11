<script>
  import Search from "svelte-search";
  let value = "";
  export let searchTitlesProp = () => {};
  export let clearSearchProp = () => {};
</script>

<Search
  label="searchBox"
  hideLabel
  bind:value
  on:submit={searchTitlesProp({ value })}
  placeholder="Search..."
/>

<button on:click={() => ((value = ""), clearSearchProp())}>Reset search</button>

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
