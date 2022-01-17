<script>
  import {
    Tile,
    Tabs,
    Tab,
    TabContent,
    Content,
    Grid,
    Row,
    Column,
    Header,
    FluidForm,
    TextInput,
  } from "carbon-components-svelte";

  import { getStats } from "$lib/stats.js";

  import UserAvatar32 from "carbon-icons-svelte/lib/UserAvatar32";
  import Catalog32 from "carbon-icons-svelte/lib/Catalog32";
  import ThumbsUp32 from "carbon-icons-svelte/lib/ThumbsUp32";

  let stats = getStats();
</script>

<Header platformName="Eagle Monitor" />

<Content>
  <Grid>
    <Tile>
      <Row>
        <Column>
          <h1>Eagle Monitor - BC Course Search & Notifications</h1>
        </Column>
      </Row>
    </Tile>

    <Row>
      <Column>
        <h2>What is Eagle Monitor?</h2>
        <p>
          Eagle Monitor helps you keep track of courses that are full that you
          want to join. You can select courses that you want to be notified
          about when an opening arises. Additionally, Eagle Monitor provides an
          easy-to-use search page to learn more about BC courses.
        </p>
      </Column>

      <Column>
        <FluidForm>
          <TextInput
            labelText="@bc.edu email address"
            placeholder="Enter your Boston College email..."
            required
          />
        </FluidForm>
      </Column>
    </Row>

    <Row>
      {#await stats}
        <Column><UserAvatar32 /> loading...</Column>
        <Column><Catalog32 /> loading...</Column>
        <Column><ThumbsUp32 /> loading...</Column>
      {:then statsObj}
        <Column><UserAvatar32 /> {statsObj.profiles} Users</Column>
        <Column><Catalog32 /> {statsObj.courses} Courses Listed</Column>
        <Column
          ><ThumbsUp32 /> {statsObj.notifications} Notifications Sent</Column
        >
      {/await}
    </Row>

    <div>
      <div>
        <p>
          EagleMonitor has been helping Boston College students find, monitor,
          and register for courses since 2022.
        </p>
      </div>
      <div><strong>Creator</strong>: Will Morrison</div>
      <div><strong>Contact</strong>: morriswk@bc.edu</div>
      <div>
        <strong>Contributors</strong>: Andrew Clark, Jonathan Zarnstorff
      </div>
    </div>
  </Grid>
</Content>
