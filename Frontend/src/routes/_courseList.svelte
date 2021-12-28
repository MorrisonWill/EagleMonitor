<script>
  import { supabase } from "$lib/supabaseClient";

  let from = 0;
  let to = 9;

  let rows = [];

  async function getCourses(start, end) {
    let { data: courses } = await supabase
      .from("courses")
      .select("*")
      .range(start, end);
    rows = courses;
  }

  function nextPage() {
    from += 10;
    to += 10;
    getCourses(from, to);
  }

  function previousPage() {
    from -= 10;
    to -= 10;
    getCourses(from, to);
  }

  getCourses(from, to);
</script>

<table class="table">
  <thead>
    <tr>
      <th>ID</th>
      <th>Name</th>
      <th>Status</th>
    </tr>
  </thead>
  <tbody>
    {#await rows}
      <p>loading...</p>
    {:then rows}
      {#each rows as row}
        <tr>
          <td>{row.id}</td>
          <td>{row.name}</td>
          <td>{row.status}</td>
        </tr>
      {:else}
        <tr>
          <h5 class="text-center">There are no courses to show.</h5>
        </tr>
      {/each}
    {/await}
  </tbody>
</table>

{#if from != 0}
  <button on:click={previousPage}>previous</button>
{/if}
<button on:click={nextPage}>next</button>
