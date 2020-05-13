<div class="row justify-content-md-center">
    <div class="col-6">
        <Calendar bind:date />
    </div>
</div>
<div class="row justify-content-md-center">
    <div class="col-11">
        <List bind:date />
    </div>
</div>

<script>
import Calendar from './Calendar.svelte'
import List from './List.svelte'

let date = null
</script>
