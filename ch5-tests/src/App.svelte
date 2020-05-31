<Navbar />

<div class="container mx-auto w-full lg:w-3/5 px-2 pt-2 mt-2">
    <Router {routes} />
</div>

<script>
// Router
import Router from 'svelte-spa-router'
import routes from './routes.js'

// Components
import Navbar from './components/Navbar.svelte'
</script>