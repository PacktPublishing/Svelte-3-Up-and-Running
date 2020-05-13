<Datepicker bind:selected />

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
