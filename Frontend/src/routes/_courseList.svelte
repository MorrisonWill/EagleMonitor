<script>
  import { onMount } from "svelte";

  import { supabase } from "$lib/supabaseClient";

  import CourseTile from "./_courseTile.svelte";

  import {
    Header,
    InlineNotification,
    Content,
    Grid,
    Row,
    Column,
    Search,
    Toggle,
    Button,
    Theme,
    SkeletonText,
    Tile,
  } from "carbon-components-svelte";

  let inSearch = false;
  let loading = false;

  let noResults = false;

  let search = "";

  let onlyOpen = false;

  let start = 0;
  let end = 9;

  let data = [];
  let newBatch = [];

  let theme = "white";

  let monitored = [];

  async function monitor(id) {
    let { data: courses } = await supabase
      .from("profiles")
      .select("courses")
      .single()
      .eq("id", supabase.auth.user().id);

    let existing = courses.courses;

    if (existing == null) {
      existing = [];
    }

    if (existing.includes(id)) {
      existing.splice(existing.indexOf(id), 1);
    } else {
      existing.push(id);
    }

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
      loading = false;

      newBatch = courses;
      return;
    }
    let { data: courses } = await supabase
      .from("courses")
      .select("*")
      .order("id", { ascending: true })
      .range(start, end);
    loading = false;

    newBatch = courses;
  }

  function clearSearch() {
    getMonitored();

    noResults = false;

    inSearch = false;
    data = [];
    newBatch = [];

    start = 0;
    end = 9;
    fetchData();
  }

  async function searchTitles() {
    getMonitored();
    noResults = false;

    data = [];
    newBatch = [];

    inSearch = true;
    if (onlyOpen) {
      const { data: courses } = await supabase
        .from("courses")
        .select()
        .filter("status", "in", '("open")')
        .order("id", { ascending: true })
        .textSearch("name", search, {
          type: "websearch",
          config: "english",
        });
      if (courses.length == 0) {
        clearSearch();
        noResults = true;
        return;
      }
      data = courses;

      return;
    }
    const { data: courses } = await supabase
      .from("courses")
      .select()
      .order("id", { ascending: true })
      .textSearch("name", search, {
        type: "websearch",
        config: "english",
      });
    if (courses.length == 0) {
      clearSearch();
      noResults = true;
      return;
    }

    data = courses;
  }

  async function getMonitored() {
    let { data: courses } = await supabase
      .from("profiles")
      .select("courses")
      .single();

    monitored = courses.courses;

    if (monitored == null) {
      monitored = [];
    }
  }

  onMount(() => {
    getMonitored();
    fetchData();
    console.log(data);
  });

  $: data = [...data, ...newBatch];
</script>

<Theme bind:theme />

<Header platformName="Eagle Monitor Course List" />

<Content>
  <Grid>
    <Row>
      <Column>
        <p>
          You can search for courses below or scroll to browse BC's course list.
          To view more information about a course, click on the tile. To monitor
          a course (receive an email when it opens), click the monitor button at
          the top right of the tile.
        </p>
        <br />
      </Column>
    </Row>
    <Row>
      <Column>
        <Search
          placeholder="Search by title..."
          bind:value={search}
          on:clear={clearSearch}
          on:change={searchTitles}
          size="sm"
        />
      </Column>
    </Row>
    {#if noResults}
      <Row>
        <Column>
          <InlineNotification
            lowContrast
            kind="error"
            title="Error:"
            subtitle="No results found. Please try a different search."
            timeout="10000"
          />
        </Column>
      </Row>
    {/if}

    <br />
    <Row>
      <Column>
        <Toggle
          size="sm"
          labelText="Only show open courses"
          on:toggle={() => {
            onlyOpen = !onlyOpen;
            console.log(onlyOpen);
            clearSearch();
          }}
        />
      </Column>
    </Row>
  </Grid>

  <Grid padding>
    {#await data then rows}
      {#each rows as row}
        <Row>
          <Column>
            <CourseTile
              monitoredArray={monitored}
              monitorProp={(id) => monitor(id)}
              course={row}
            />
          </Column>
        </Row>
      {:else}
        <Grid padding>
          {#each Array(10) as _}
            <Row>
              <Column>
                <Tile>
                  <SkeletonText />
                  <SkeletonText />
                  <SkeletonText />
                </Tile>
              </Column>
            </Row>
          {/each}
        </Grid>
      {/each}
    {/await}

    {#if loading}
      <Row>
        <Column>
          <Tile>
            <SkeletonText />
            <SkeletonText />
            <SkeletonText />
          </Tile>
        </Column>
      </Row>
    {/if}

    {#if !inSearch}
      <Row>
        <Column>
          <Button
            on:click={() => {
              loading = true;
              start += 10;
              end += 10;
              fetchData();
            }}>Load more</Button
          >
        </Column>
      </Row>
    {/if}
  </Grid>
</Content>
