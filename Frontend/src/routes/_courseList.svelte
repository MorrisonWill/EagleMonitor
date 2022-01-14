<script>
  import { onMount } from "svelte";
  import CourseCard from "./_courseCard.svelte";
  import { supabase } from "$lib/supabaseClient";
  import SearchBox from "./_searchBox.svelte";
  import MyCourses from "./_myCourses.svelte";

  let onlyOpen = false;
  let inSearch = false;
  let myCourses = false;

  let start = 0;
  let end = 9;

  let data = [];
  let newBatch = [];

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

  async function fetchData() {
    if (onlyOpen) {
      let { data: courses } = await supabase
        .from("courses")
        .select("*")
        .order("id", { ascending: true })
        .filter("status", "in", '("open")')
        .range(start, end);
      newBatch = courses;
      return;
    }
    let { data: courses } = await supabase
      .from("courses")
      .select("*")
      .order("id", { ascending: true })
      .range(start, end);

    newBatch = courses;
  }

  async function searchTitles(search) {
    data = [];
    newBatch = [];

    inSearch = true;
    if (onlyOpen) {
      const { data: courses } = await supabase
        .from("courses")
        .select()
        .filter("status", "in", '("open")')
        .order("id", { ascending: true })
        .textSearch("name", search.value);
      data = courses;
      return;
    }
    const { data: courses } = await supabase
      .from("courses")
      .select()
      .order("id", { ascending: true })
      .textSearch("name", search.value);
    console.log(courses);
    data = courses;
  }

  function clearSearch() {
    inSearch = false;
    data = [];
    newBatch = [];

    start = 0;
    end = 9;
    fetchData();
  }

  function toggleOnlyOpen() {
    onlyOpen = !onlyOpen;

    clearSearch();
  }

  function toggleMyCourses() {
    myCourses = !myCourses;
  }

  onMount(() => {
    // load first batch onMount
    fetchData();
  });

  $: data = [...data, ...newBatch];
</script>

<button
  on:click={toggleMyCourses}
  class="block p-4 text-sm font-semibold hover:bg-indigo-50 hover:text-indigo-500 rounded"
  >View My Courses</button
>

{#if myCourses}
  <MyCourses />
{/if}

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

<body>
  <div class="container mt-4 mx-auto">
    <div class="grid grid-cols-1">
      {#await data}
        <p>loading course info...</p>
      {:then rows}
        {#each rows as row}
          <CourseCard monitorProp={(id) => monitor(id)} course={row} />
        {:else}
          <p>loading course info...</p>
        {/each}
      {/await}
    </div>
  </div>

  {#if !inSearch}
    <button
      on:click={() => {
        start += 10;
        end += 10;
        fetchData();
      }}>Load more</button
    >
  {/if}
</body>
