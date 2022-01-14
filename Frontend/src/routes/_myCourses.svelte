<script>
  import { supabase } from "$lib/supabaseClient";
  import { onMount } from "svelte";

  let courses = [];

  function removeItemFromArray(array, n) {
    const newArray = [];

    for (let i = 0; i < array.length; i++) {
      if (array[i] !== n) {
        newArray.push(array[i]);
      }
    }
    return newArray;
  }

  async function remove(id) {
    let { data: response } = await supabase.from("profiles").select("courses");

    let pruned = removeItemFromArray(response[0].courses, id);

    await supabase.from("profiles").update({ courses: pruned });

    courses = pruned;
  }

  async function fetchCourses() {
    let { data: response } = await supabase.from("profiles").select("courses");
    courses = response[0].courses;
  }

  onMount(() => {
    fetchCourses();
  });
</script>

<div class="flex flex-row justify-center items-center">
  <table>
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
              ><button on:click={() => remove(row)} class="bg-red-100"
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
</div>
