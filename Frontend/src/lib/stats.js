import { supabase } from "$lib/supabaseClient";

export async function getStats() {
  let profiles = await getNumProfiles();
  let courses = await getNumCourses();
  let notifications = await getNumNotifications();

  return {
    profiles,
    courses,
    notifications,
  };
}

export async function getNumProfiles() {
  let { data, error } = await supabase.rpc("get_num_profiles");

  if (error) console.error(error);
  return data;
}

export async function getNumCourses() {
  let { data, error } = await supabase.rpc("get_num_courses");
  if (error) console.error(error);
  return data;
}

export async function getNumNotifications() {
  let { data, error } = await supabase
    .from("stats")
    .select("*")
    .eq("id", "notified");
  return data[0].stat;
}
