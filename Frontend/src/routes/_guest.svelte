<script>
  import { supabase } from "$lib/supabaseClient";

  import {
    Form,
    InlineNotification,
    Tile,
    Tabs,
    Tab,
    TabContent,
    Content,
    Button,
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

  let loading = false;
  let emailSent = false;
  let error = "";
  let email;

  const handleLogin = async () => {
    try {
      if (!email.includes("bc.edu")) {
        error = "only bc.edu email addresses may be used";
      } else {
        loading = true;
        const { err } = await supabase.auth.signIn({ email });
        if (err) throw err;
        error = "";
        emailSent = true;
      }
    } catch (e) {
      error = e.error_description || e.message;
    } finally {
      loading = false;
    }
  };
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

    <br />

    {#if emailSent}
      <InlineNotification
        lowContrast
        kind="success"
        title="Success:"
        subtitle="Please check your email for the login link"
      />
    {:else if error}
      <InlineNotification
        lowContrast
        kind="error"
        title="Error:"
        subtitle={error}
      />
    {/if}
    <Row margin: $spacing-05>
      <Column>
        <h2>What is Eagle Monitor?</h2>
        <br />
        <p>
          Eagle Monitor helps you keep track of courses that are full that you
          want to join. You can select courses that you want to be notified
          about when an opening arises. Additionally, Eagle Monitor provides an
          easy-to-use search page to learn more about BC courses.
        </p>
      </Column>

      <div class="emailcontainer">
        <Column>
          <FluidForm on:submit={handleLogin}>
            <TextInput
              labelText="@bc.edu email address"
              invalidText="Please use a valid @bc.edu email address"
              placeholder="Enter your Boston College email..."
              bind:value={email}
              required
            />
            <Button kind="secondary" tabIndex={0} type="submit">Sign In</Button>
          </FluidForm>
        </Column>
      </div>
    </Row>

    <div class="statscontainer" style="position:relative;">
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
    </div>

    <div class="infocontainer">
      <div class="summarycontainer">
        <p style="color:#ffffff">
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

<style>
  .titlecontainer {
    height: 5em;
    width: 100%;
  }

  .emailcontainer {
    height: 5em;
    width: 66%;
    padding-top: 1em;
    padding-bottom: 6em;
  }

  .statscontainer {
    padding-bottom: 4em;
    margin: 40px 5px 0 0px;
  }

  :global(.infocontainer > *) {
    height: 1em;
    width: 100%;
    display: inline-block;
    padding: 1em;
    background: #f4f4f4;
  }

  .summarycontainer {
    width: 100%;
    height: 100%;
    background: #393939;
  }
</style>
