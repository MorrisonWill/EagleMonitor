<script>
  export let course;

  export let monitorProp;

  let expanded = false;
  let monitored = false;

  function toggleExpansion() {
    expanded = !expanded;
  }
</script>

<div
  on:click={toggleExpansion}
  class="card m-2 cursor-pointer border border-gray-400 rounded-lg hover:shadow-md hover:border-opacity-0 transform hover:-translate-y-1 transition-all duration-200"
>
  <div class="m-3">
    {#if course.status == "open"}
      <span
        class="text-sm text-teal-800 font-mono bg-green-100 inline rounded-full px-2 align-top float-right animate-pulse"
        >open</span
      >
    {/if}
    {#if monitored}
      <span
        class="text-sm text-black-800 font-mono bg-red-100 inline rounded-full px-2 align-top float-right animate-pulse"
        >monitored</span
      >
    {/if}
    <h2 class="text-lg mb-2">
      {course.name}
      <button
        on:click={() => {
          monitored = true;
          monitorProp(course.id);
        }}
        class="bg-blue-300 hover:bg-blue-400 rounded">monitor</button
      >
    </h2>
    <p
      class="font-light font-mono text-sm text-gray-700 hover:text-gray-900 transition-all duration-200"
    >
      {course.description}
    </p>
  </div>

  <table>
    <thead>
      <tr class="active">
        <th class="text-left">School</th>
        <th>Credits</th>
        <th>Course Level</th>
      </tr>
    </thead>
    <tbody>
      <tr class="active">
        <td>{course.school}</td>
        <td class="text-center">{course.creditCount}</td>
        <td>{course.courseLevel}</td>
      </tr>
    </tbody>
  </table>
  {#if expanded}
    <table>
      <thead>
        <tr class="active">
          <th>Section</th>
          <th>Total</th>
          <th>Used</th>
          <th>Room</th>
          <th>Time</th>
          <th>Instructor</th>
        </tr>
      </thead>

      {#each course.seatData as seatData, index}
        <tbody>
          <tr class="active">
            <td>{seatData.split(":")[0].split(" ")[1]}</td>
            <td class="text-center">{seatData.split(":")[1]}</td>
            <td>{seatData.split(":")[2]}</td>
            <td>{course.rooms[index]}</td>
            <td class="px-5">{course.times[index]}</td>
            <td>{course.instructors[index]}</td>
          </tr>
        </tbody>
      {/each}
    </table>
  {/if}
</div>

<!--
    <button
      on:click={() => {
        monitored = true;
        monitorProp(course.id);
      }}
      class="bg-gray-300 hover:bg-gray-400 rounded absolute bottom-0 right-0 h-16 w-16"
      >monitor</button
    >

-->
