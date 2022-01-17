<script>
  import { DataTable } from "carbon-components-svelte";
  export let course;

  let rows = [];

  for (let i = 0; i < course.seatData.length; i++) {
    rows.push({
      id: course.seatData[i].split(" ")[1],
      section: course.seatData[i].split(" ")[1],
      total: course.seatData[i].split(":")[1],
      used: course.seatData[i].split(":")[2],
      room: course.rooms[i],
      time: course.times[i],
      instructor: course.instructors[i],
    });
  }
</script>

<DataTable
  headers={[
    { key: "section", value: "Section" },
    { key: "total", value: "Total" },
    { key: "used", value: "Used" },
    { key: "room", value: "Room" },
    { key: "time", value: "Time" },
    { key: "instructor", value: "Instructor" },
  ]}
  {rows}
/>
