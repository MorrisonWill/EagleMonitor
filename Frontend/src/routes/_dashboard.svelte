<script>
	import { supabase } from "$lib/supabaseClient";
	import { user } from "$lib/sessionStore";

	let loading = true;
	let courseLoading = true;
	let username = null;
	let course = null;

	async function getProfile() {
		const user = supabase.auth.user();

		const { data, count } = supabase
			.from('profiles')
			.select('id, email')
			.eq('id', user.id)
			.single()
		console.log(data)
	}

	async function randomCourse() {
		let number = Math.floor(Math.random() * 500 + 1);
		courseLoading = true;
		const response = await fetch(
			`http://localhost:8080/courses/1-${number}`,
			{
				headers: {
					Authorization: `Bearer ${
						supabase.auth.session().access_token
					}`,
				},
			}
		);
		const data = await response.json();

		data.forEach((obj) => (course = obj.name));

		courseLoading = false;
	}

	async function updateProfile() {
		try {
			loading = true;
			const user = supabase.auth.user();

			const updates = {
				id: user.id,
				username,
				website,
				avatar_url,
				updated_at: new Date(),
			};

			let { error } = await supabase.from("profiles").upsert(updates, {
				returning: "minimal", // Don't return the value after inserting
			});

			if (error) throw error;
		} catch (error) {
			alert(error.message);
		} finally {
			loading = false;
		}
	}

	async function signOut() {
		try {
			loading = true;
			let { error } = await supabase.auth.signOut();
			if (error) throw error;
		} catch (error) {
			alert(error.message);
		} finally {
			loading = false;
		}
	}

	randomCourse();
	getProfile();
</script>

<button on:click={randomCourse}> Random Course </button>

<p>{courseLoading ? "Loading ..." : course}</p>
