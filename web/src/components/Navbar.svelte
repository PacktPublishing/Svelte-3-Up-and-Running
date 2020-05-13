<nav class="navbar navbar-expand-lg navbar-light bg-light">
  <span class="navbar-brand">
    <i class="fa fa-pencil-square-o" aria-hidden="true"></i> Journal
  </span>
  <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
    <div class="navbar-nav">
      <span class="nav-item nav-link {(active == 'home') ? 'active' : ''}" on:click={() => $view = null}>
        Home
        {#if active == 'home'}
          <span class="sr-only">(current)</span>
        {/if}
      </span>
      <span class="nav-item nav-link {(active == 'add') ? 'active' : ''}" on:click={() => $view = 'add'}>
        Add
        {#if active == 'add'}
          <span class="sr-only">(current)</span>
        {/if}
      </span>
    </div>
  </div>
</nav>

<script>
import {view} from '../stores'

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
