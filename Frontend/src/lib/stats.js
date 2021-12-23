import { supabase } from "$lib/supabaseClient";

export async function getNumProfiles() {
    let { data, error } = await supabase
        .rpc('get_num_profiles');

    if (error) console.error(error);
    return data;
}

export async function getNumListedCourses() {
    let { data, error } = await supabase
        .from('stats')
        .select("*")
        .eq('id', 'listed')
    return data[0].stat
}

export async function getNumMonitoredCourses() {
    let { data, error } = await supabase
        .from('stats')
        .select("*")
        .eq('id', 'monitored')
    return data[0].stat
}

export async function getNumNotifications() {
    let { data, error } = await supabase
        .from('stats')
        .select("*")
        .eq('id', 'notified')
    return data[0].stat
}