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
