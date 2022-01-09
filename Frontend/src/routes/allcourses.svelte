<script>
  import { supabase } from "$lib/supabaseClient";

  let from = 0;
  let to = 9;

  let rows = [];

  async function monitor(id) {
    let { data: courses } = await supabase
      .from("profiles")
      .select("courses")
      .eq("id", supabase.auth.user().id);

    let existing = courses[0].courses;

    if (existing == null) {
      existing = [];
    }

    existing.push(id);

    let cleaned = [...new Set(existing)];

    await supabase
      .from("profiles")
      .update({ courses: cleaned })
      .eq("id", supabase.auth.user().id);
  }

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

  export let style;
  let blueStyle = "blueTable";
</script>

<table class="blueTable">
  <thead>
    <tr>
      <th>ID</th>
      <th>Name</th>
      <th>Status</th>
      <th>Monitor</th>
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
          <td>
            <button on:click={() => monitor(row.id)} class="bg-red-100">
               <img src="../src/assets/plus.png" width=10px height=10px>
            </button>
          </td>
        </tr>
      {:else}
        <tr>
          <h5 class="text-center">There are no courses to show.</h5>
        </tr>
      {/each}
    {/await}
  </tbody>
</table>

<div class="container">
  {#if from != 0}
    <button class="primary" on:click={previousPage}>Previous</button>
  {/if}
    <button class="primary" on:click={nextPage}>Next</button>
</div>

<style>
  /* Table Style */
  table,
  th,
  td {
    border: 2px solid;
    border-collapse: collapse;
    margin: 10px;
  }

  table.blueTable {
    border: 1px solid #1c6ea4;
    background-color: #eeeeee;
    width: 100%;
    text-align: left;
    border-collapse: collapse;
  }

  table.blueTable td,
  table.blueTable th {
    border: 1px solid #aaaaaa;
    padding: 3px 2px;
  }
  table.blueTable tbody td {
    font-size: 13px;
  }
  table.blueTable tr:nth-child(even) {
    background: #d0e4f5;
  }
  table.blueTable thead {
    background: #1c6ea4;
    background: -moz-linear-gradient(
      top,
      #5592bb 0%,
      #327cad 66%,
      #1c6ea4 100%
    );
    background: -webkit-linear-gradient(
      top,
      #5592bb 0%,
      #327cad 66%,
      #1c6ea4 100%
    );
    background: linear-gradient(
      to bottom,
      #5592bb 0%,
      #327cad 66%,
      #1c6ea4 100%
    );
    border-bottom: 2px solid #444444;
  }
  table.blueTable thead th {
    font-size: 15px;
    font-weight: bold;
    color: #ffffff;
    border-left: 2px solid #d0e4f5;
  }
  table.blueTable thead th:first-child {
    border-left: none;
  }

  /* Button Styles */
  .primary {
    background-color:#2563eb;
    border-radius: 10px;
    border-color: black;
    color: white;
    padding: 15px 32px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;  
  }
  .secondary {
    color:black;
  }

  /* Button Container */
  .container {
    margin-left:auto;
    margin-right: auto;
    width: 15em;
    padding-top: 1vw;

  }

</style>
