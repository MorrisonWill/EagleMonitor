<script>
  import { getNumProfiles, getNumListedCourses, getNumMonitoredCourses, getNumNotifications } from "$lib/stats.js"

  let numProfiles = getNumProfiles();
  let numListedCourses = getNumListedCourses();
  let numMonitoredCourses = getNumMonitoredCourses();
  let numNotifications = getNumNotifications();
</script>

<section class="xl:bg-contain bg-top bg-no-repeat" style="background-image: url('https://shuffle.dev/metis-assets/backgrounds/intersect.svg');"><div class="container px-4 mx-auto">
    <div class="py-12 text-center">
    <div class="max-w-lg mx-auto mb-8">
      <h2 class="text-3xl md:text-4xl mb-4 font-bold font-heading">
        <span>Eagle Monitor</span>
        <span class="text-blue-600">helps you</span>
        <span>register for courses.</span>
      </h2>
      <p class="text-blueGray-400 leading-relaxed">Sign up today to get notifications about Boston College courses.</p>
    </div>
    <div><a class="block sm:inline-block py-4 px-8 mb-4 sm:mb-0 sm:mr-3 text-xs text-white text-center font-semibold leading-none bg-blue-600 hover:bg-blue-700 rounded" href="#">Join Now</a><a class="block sm:inline-block py-4 px-8 text-xs text-blueGray-500 hover:text-blueGray-600 text-center font-semibold leading-none bg-white rounded" href="faq">Learn More</a></div>
    </div>
    <div class="relative max-w-3xl mt-6 mb-8 mx-auto">
      <img src="https://shuffle.dev/metis-assets/elements/pattern-small.png" alt=""><div class="absolute" style="top: 3%; left: 9%; width: 82%; height: 90%;">
      <img class="object-contain w-full h-full" src="https://shuffle.dev/metis-assets/elements/app-mockup.png" alt=""></div>
    </div>
    <div class="flex flex-wrap justify-between pt-8 pb-16">
      <div class="flex w-1/2 lg:w-auto py-4">
        <div class="flex justify-center items-center bg-blueGray-50 rounded-xl h-12 w-12 sm:h-20 sm:w-20">
          <img class="h-6 sm:h-8" src="https://shuffle.dev/metis-assets/icons/user-circle.svg" alt=""></div>
        <div class="sm:py-2 ml-2 sm:ml-6">
          {#await numProfiles}
            <span class="sm:text-2xl font-bold font-heading">loading...</span>
          {:then result}
            <span class="sm:text-2xl font-bold font-heading">{result}</span>
          {:catch error}
            <span class="sm:text-2xl font-bold font-heading">Error getting the number of users.</span>
          {/await}
          <p class="text-xs sm:text-base text-blueGray-400">Users</p>
        </div>
      </div>
      <div class="flex w-1/2 lg:w-auto py-4">
        <div class="flex justify-center items-center bg-blueGray-50 rounded-xl h-12 w-12 sm:h-20 sm:w-20">
          <img class="h-6 sm:h-8" src="https://shuffle.dev/metis-assets/icons/puzzle.svg" alt=""></div>
        <div class="sm:py-2 ml-2 sm:ml-6">
          {#await numListedCourses}
            <span class="sm:text-2xl font-bold font-heading">loading...</span>
          {:then result}
            <span class="sm:text-2xl font-bold font-heading">{result}</span>
          {:catch error}
            <span class="sm:text-2xl font-bold font-heading">Error getting the number of listed courses.</span>
          {/await}
          <p class="text-xs sm:text-base text-blueGray-400">Courses Listed</p>
        </div>
      </div>
      <div class="flex w-1/2 lg:w-auto py-4">
        <div class="flex justify-center items-center bg-blueGray-50 rounded-xl h-12 w-12 sm:h-20 sm:w-20">
          <img class="h-6 sm:h-8" src="https://shuffle.dev/metis-assets/icons/thumb-up.svg" alt=""></div>
        <div class="sm:py-2 ml-2 sm:ml-6">
          {#await numMonitoredCourses}
            <span class="sm:text-2xl font-bold font-heading">loading...</span>
          {:then result}
            <span class="sm:text-2xl font-bold font-heading">{result}</span>
          {:catch error}
            <span class="sm:text-2xl font-bold font-heading">Error getting the number of monitored courses.</span>
          {/await}
          <p class="text-xs sm:text-base text-blueGray-400">Monitored Courses</p>
        </div>
      </div>
      <div class="flex w-1/2 lg:w-auto py-4">
        <div class="flex justify-center items-center bg-blueGray-50 rounded-xl h-12 w-12 sm:h-20 sm:w-20">
          <img class="h-6 sm:h-8" src="https://shuffle.dev/metis-assets/icons/terminal.svg" alt=""></div>
        <div class="sm:py-2 ml-2 sm:ml-6">
          {#await numNotifications}
            <span class="sm:text-2xl font-bold font-heading">loading...</span>
          {:then result}
            <span class="sm:text-2xl font-bold font-heading">{result}</span>
          {:catch error}
            <span class="sm:text-2xl font-bold font-heading">Error getting the number of notifications sent.</span>
          {/await}
          <p class="text-xs sm:text-base text-blueGray-400">Notifications Sent</p>
        </div>
      </div>
    </div>
    </div>
  </section>