<script>
  import { supabase } from "$lib/supabaseClient";
  import { user } from "$lib/sessionStore";
  import MyCourses from "./_myCourses.svelte";
  import CourseList from "./_courseList.svelte";
</script>

<CourseList />
