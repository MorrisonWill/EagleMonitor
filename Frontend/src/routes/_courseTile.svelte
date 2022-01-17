<script>
  import SectionTable from "./_sectionTable.svelte";
  import { ClickableTile, Tag } from "carbon-components-svelte";

  export let course;

  export let monitorProp;

  export let monitoredArray;

  let expanded = false;
  let monitored = false;

  if (monitoredArray.includes(course.id)) {
    monitored = true;
  }

  let open = course.status == "open";
</script>

<ClickableTile on:click={() => (expanded = !expanded)}>
  <strong>{course.name}</strong>
  <Tag type={open ? "green" : "red"}>{open ? "open" : "closed"}</Tag>
  {#if course.status == "closed"}
    <Tag
      style="float: right;"
      on:click={() => {
        monitored = !monitored;
        monitorProp(course.id);
      }}
      type={monitored ? "high-contrast" : "outline"}
      interactive>{monitored ? "monitored" : "monitor"}</Tag
    >
  {/if}

  <p>{course.description}</p>

  <br />

  <table>
    <thead>
      <tr>
        <th>School</th>
        <th>Credits</th>
        <th>Course Level</th>
      </tr>
    </thead>
    <tbody>
      <tr class="active">
        <td>{course.school}</td>
        <td style="text-align: center;">{course.creditCount}</td>
        <td>{course.courseLevel}</td>
      </tr>
    </tbody>
  </table>
  <small style="float: right;">Last updated {course.lastUpdated}</small>
  {#if expanded}
    <SectionTable {course} />
  {/if}
</ClickableTile>
