<script>
  import CourseCard from "./_courseCard.svelte";
  import { supabase } from "$lib/supabaseClient";
  import SearchBox from "./_searchBox.svelte";

  let from = 0;
  let to = 9;

  let rows = [];

  let onlyOpen = false;

  async function getCourses(start, end) {
    if (onlyOpen) {
      let { data: courses } = await supabase
        .from("courses")
        .select("*")
        .order("id", { ascending: true })
        .filter("status", "in", '("open")')
        .range(start, end);
      rows = courses;
      return;
    }
    let { data: courses } = await supabase
      .from("courses")
      .select("*")
      .order("id", { ascending: true })
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

  async function searchTitles(search) {
    if (onlyOpen) {
      const { data: courses } = await supabase
        .from("courses")
        .select()
        .filter("status", "in", '("open")')
        .order("id", { ascending: true })
        .textSearch("name", search.value);
      rows = courses;
      return;
    }
    const { data: courses } = await supabase
      .from("courses")
      .select()
      .order("id", { ascending: true })
      .textSearch("name", search.value);
    rows = courses;
  }

  function clearSearch() {
    rows = [];
    from = 0;
    to = 9;
    getCourses(from, to);
  }

  function toggleOnlyOpen() {
    onlyOpen = !onlyOpen;
    console.log("aaa");
    console.log(onlyOpen);

    clearSearch();
  }

  getCourses(from, to);
</script>

<SearchBox
  searchTitlesProp={(search) => searchTitles(search)}
  clearSearchProp={() => clearSearch()}
/>

<div class="flex justify-left">
  <div>
    <div class="form-check">
      <input
        on:change={() => toggleOnlyOpen()}
        class="form-check-input appearance-none h-4 w-4 border border-gray-300 rounded-sm bg-white checked:bg-blue-600 checked:border-blue-600 focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-2 cursor-pointer"
        type="checkbox"
        value=""
        id="flexCheckDefault"
      />
      <label
        class="form-check-label inline-block text-gray-800"
        for="flexCheckDefault"
      >
        Only display open courses
      </label>
    </div>
  </div>
</div>

{#if from != 0}
  <button on:click={previousPage}>previous</button>
{/if}
<button on:click={nextPage}>next</button>

<body>
  <div class="container mt-4 mx-auto">
    <div class="grid grid-cols-1">
      {#await rows}
        <p>loading course info...</p>
      {:then rows}
        {#each rows as row}
          <CourseCard course={row} />
        {:else}
          <p>loading course info...</p>
        {/each}
      {/await}
    </div>
  </div>
</body>

<!--
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

{#if from != 0}
  <button on:click={previousPage}>previous</button>
{/if}
<button on:click={nextPage}>next</button>

<style>
  table,
  th,
  td {
    border: 1px solid;
    border-collapse: collapse;
    margin: 10px;
  }

  table.redTable {
    border: 2px solid #a40808;
    background-color: #eee7db;
    width: 100%;
    text-align: center;
    border-collapse: collapse;
  }
  table.redTable td,
  table.redTable th {
    border: 1px solid #aaaaaa;
    padding: 3px 2px;
  }
  table.redTable tbody td {
    font-size: 13px;
  }
  table.redTable tr:nth-child(even) {
    background: #f5c8bf;
  }
  table.redTable thead {
    background: #a40808;
    background: -moz-linear-gradient(
      top,
      #bb4646 0%,
      #ad2020 66%,
      #a40808 100%
    );
    background: -webkit-linear-gradient(
      top,
      #bb4646 0%,
      #ad2020 66%,
      #a40808 100%
    );
    background: linear-gradient(
      to bottom,
      #bb4646 0%,
      #ad2020 66%,
      #a40808 100%
    );
  }
  table.redTable thead th {
    font-size: 19px;
    font-weight: bold;
    color: #ffffff;
    text-align: center;
    border-left: 2px solid #a40808;
  }
  table.redTable thead th:first-child {
    border-left: none;
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
</style>
-->
