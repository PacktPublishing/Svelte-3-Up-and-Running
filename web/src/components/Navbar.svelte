<nav class="bg-white w-full shadow">
  <div class="w-full lg:w-3/5 container mx-auto px-2 flex flex-wrap items-center justify-between">
    <div class="my-4 pl-4 md:pl-0 text-lg">
      <i class="fa fa-pencil-square-o" aria-hidden="true"></i> Journal
    </div>
    <div class="my-4 pr-4 md:pr-0">
      <span
        class="ml-4 cursor-pointer {(active == 'home') ? 'text-blue-600 underline' : ''}"
        on:click={() => $view = null}>
        Home
      </span>
      <span
        class="ml-4 cursor-pointer {(active == 'add') ? 'text-blue-600 underline' : ''}"
        on:click={() => $view = 'add'}>
        Add
      </span>
    </div>
  </div>
</nav>

<script>
import {view} from '../stores.js'

let active
$: {
    switch ($view) {
        case 'add':
            active = $view
            break

        default:
            active = 'home'
            break
    }
}
</script>
