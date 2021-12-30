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
    console.log(courses[0].courses);

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

  async function test() {
    let { data: courses } = await supabase
      .from("courses")
      .select("*")
      .contains("instructors", ["McElwaine, Michelle L"]);
    console.log(courses);
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

  test();

  getCourses(from, to);

  export let style;
  let blueStyle="blueTable";
</script>

<table class="blueTable"> 
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
          <td
            ><button on:click={() => monitor(row.id)} class="bg-red-100"
              >monitor</button
            ></td
          >
        </tr>
      {:else}
        <tr>
          <h5 class="text-center">There are no courses to show.</h5>
        </tr>
      {/each}
    {/await}
  </tbody>
</table>


<style>
 table, th, td {
  border: 2px solid;
  border-collapse: collapse;
  margin: 10px; 
 }
 
 table.blueTable {
  border: 1px solid #1C6EA4;
  background-color: #EEEEEE;
  width: 100%;
  text-align: left;
  border-collapse: collapse;
 }

 table.blueTable td, table.blueTable th {
  border: 1px solid #AAAAAA;
  padding: 3px 2px;
 }
 table.blueTable tbody td {
   font-size: 13px
 }
 table.blueTable tr:nth-child(even) {
  background: #D0E4F5;
 }
 table.blueTable thead {
  background: #1C6EA4;
  background: -moz-linear-gradient(top, #5592bb 0%, #327cad 66%, #1C6EA4 100%);
  background: -webkit-linear-gradient(top, #5592bb 0%, #327cad 66%, #1C6EA4 100%);
  background: linear-gradient(to bottom, #5592bb 0%, #327cad 66%, #1C6EA4 100%);
  border-bottom: 2px solid #444444;
 }
 table.blueTable thead th {
  font-size: 15px;
  font-weight: bold;
  color: #FFFFFF;
  border-left: 2px solid #D0E4F5;
 }
 table.blueTable thead th:first-child {
  border-left: none;
 } 
</style>

{#if from != 0}
  <button on:click={previousPage}>previous</button>
{/if}
<button on:click={nextPage}>next</button>

