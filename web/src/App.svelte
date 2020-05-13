<div class="container">
  <div class="row">
    <div class="col-12">
        <Navbar />
    </div>
  </div>
  <div class="col-12">
    {#if $view == 'add'}
        <ViewAdd />
    {:else if $view && $view.startsWith('view/')}
        <ViewObject objectId={$view.substring(5)} />
    {:else}
        <ViewList />
    {/if}
  </div>
</div>

<script>
// Components
import Navbar from './components/Navbar.svelte'
import ViewAdd from './components/View-Add.svelte'
import ViewObject from './components/View-Object.svelte'
import ViewList from './components/View-List.svelte'

// Stores
import {view} from './stores'
</script>