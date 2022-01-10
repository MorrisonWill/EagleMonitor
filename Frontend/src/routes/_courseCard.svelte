<script>
  export let course;

  let expanded = false;

  function toggleExpansion() {
    expanded = !expanded;
  }

  let seatDataArray = [];

  function arrangeSeatData(input) {
    let split = input.split(":");
    var obj = {
      title: split[0],
      total: split[1],
      used: split[2],
    };

    seatDataArray.push(obj);
  }

  course.seatData.forEach(arrangeSeatData);
</script>

<div
  on:click={toggleExpansion}
  class="card m-2 cursor-pointer border border-gray-400 rounded-lg hover:shadow-md hover:border-opacity-0 transform hover:-translate-y-1 transition-all duration-200"
>
  <div class="m-3">
    <h2 class="text-lg mb-2">
      {course.name}
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
    {#each course.seatData as seatData, index}
      <table>
        <thead>
          <tr class="active">
            <th class="text-left">Section</th>
            <th>Total</th>
            <th>Used</th>
            <th>Room</th>
            <th>Time</th>
            <th>Instructor</th>
          </tr>
        </thead>
        <tbody>
          <tr class="active">
            <td>{seatData.split(":")[0]}</td>
            <td class="text-center">{seatData.split(":")[1]}</td>
            <td>{seatData.split(":")[2]}</td>
            <td>{course.rooms[index]}</td>
            <td>{course.times[index]}</td>
            <td>{course.instructors[index]}</td>
          </tr>
        </tbody>
      </table>
    {/each}
  {/if}
</div>
