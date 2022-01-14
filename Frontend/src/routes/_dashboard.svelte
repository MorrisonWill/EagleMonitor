<script>
  import { supabase } from "$lib/supabaseClient";
  import { user } from "$lib/sessionStore";
  import CourseList from "./_courseList.svelte";
</script>

<CourseList />
