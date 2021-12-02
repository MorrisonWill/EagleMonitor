<script>
  import { supabase } from './supabaseClient'
  import { user } from './sessionStore'

  let loading = true
  let courseLoading = true
  let username = null
  let course = null

  async function getProfile() {
    try {
      loading = true
      const user = supabase.auth.user()

      let { data, error, status } = await supabase
        .from('profiles')
        .select(`username, website, avatar_url`)
        .eq('id', user.id)
        .single()

      if (error && status !== 406) throw error

      if (data) {
        username = data.username
      }
    } catch (error) {
      alert(error.message)
    } finally {
      loading = false
    }
  }

  async function randomCourse() {
    let number = Math.floor((Math.random() * 500) + 1); 
    courseLoading = true
    const response = await fetch(`http://localhost:8080/courses/1-${number}`, {
      headers: {
        Authorization: `Bearer ${supabase.auth.session().access_token}`,
      },
    });
    const data = await response.json();

    data.forEach(obj => course = obj.name)

    courseLoading = false;

  }

  async function updateProfile() {
    try {
      loading = true
      const user = supabase.auth.user()

      const updates = {
        id: user.id,
        username,
        website,
        avatar_url,
        updated_at: new Date(),
      }

      let { error } = await supabase.from('profiles').upsert(updates, {
        returning: 'minimal', // Don't return the value after inserting
      })

      if (error) throw error
    } catch (error) {
      alert(error.message)
    } finally {
      loading = false
    }
  }

  async function signOut() {
    try {
      loading = true
      let { error } = await supabase.auth.signOut()
      if (error) throw error
    } catch (error) {
      alert(error.message)
    } finally {
      loading = false
    }
    
  }

  randomCourse();
</script>

<form use:getProfile class="form-widget" on:submit|preventDefault={updateProfile}>
  <div>
    <label for="email">Email</label>
    <input id="email" type="text" value={$user.email} disabled />
  </div>
  <div>
    <label for="username">Name</label>
    <input
      id="username"
      type="text"
      bind:value={username}
    />
  </div>
  <div>
    <input type="submit" class="button block primary" value={loading ? 'Loading ...' : 'Update'} disabled={loading}/>
  </div>

  <div>
    <button class="button block" on:click={signOut} disabled={loading}>
      Sign Out
    </button>
  </div>
</form>

<br/>
<br/>

<button on:click={randomCourse}>
  Random Course
</button>

<p>{courseLoading ? 'Loading ...' : course}</p>
