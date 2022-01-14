<script>
  import { supabase } from "$lib/supabaseClient";
  import { onMount } from "svelte";

  let courses = [];

  async function remove(id) {
    let { data: response } = await supabase.from("profiles").select("courses");

    let pruned = response[0].courses;

    await supabase.from("profiles").update({ courses: pruned });
  }

  async function fetchCourses() {
    let { data: response } = await supabase.from("profiles").select("courses");
    courses = response[0].courses;
    console.log(response);
  }

  onMount(() => {
    fetchCourses();
  });
</script>

<table class="blueTable">
  <thead>
    <tr>
      <th>ID</th>
    </tr>
  </thead>
  <tbody>
    {#await courses}
      <p>loading...</p>
    {:then rows}
      {#each rows as row}
        <tr>
          <td>{row}</td>
          <td
            ><button on:click={() => remove(row.id)} class="bg-red-100"
              >remove</button
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
