<script>
  import {
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
</script>

<Header platformName="Eagle Monitor Home Page" />

<Content>
  <h1>Eagle Monitor - BC Course Search & Notifications</h1>

  <h2>What is Eagle Monitor?</h2>
  <p>
    Eagle Monitor helps you keep track of courses that are full that you want to
    join. You can select courses that you want to be notified about when an
    opening arises. Additionally, Eagle Monitor provides an easy-to-use search
    page to learn more about BC courses.
  </p>

  <FluidForm>
    <TextInput
      labelText="@bc.edu email address"
      placeholder="Enter your Boston College email..."
      required
    />
  </FluidForm>

  <p>statistics: here</p>
  <p>faq: here</p>

  <div class="landing-footer">
    <div class="landing-bio">
      EagleMonitor has been helping Boston College students monitor and register
      for courses since 2022.
    </div>
    <div class="landing-footer-creator">Creator: Will Morrison</div>
    <div class="landing-footer-contact">Contact: morriswk@bc.edu</div>
  </div>
</Content>
