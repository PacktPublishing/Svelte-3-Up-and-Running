<div class="mt-2 mb-4">
    <h2 class="text-2xl font-bold mb-2 text-gray-800">Choose a day</h2>
    <div class="ml-6">
        <Datepicker bind:selected />
    </div>
</div>

<script>
import Datepicker from 'svelte-calendar'

// Set this at midnight (local time zone)
const today = new Date()
today.setHours(0)
today.setMinutes(0)
today.setSeconds(0)
today.setMilliseconds(0)

let selected = today
let lastDate = null
export let date = null

$: {
    if (selected.getTime() != lastDate) {
        lastDate = selected.getTime()
        date = lastDate / 1000
    }
}
</script>
