{#await listPromise}
    Loading…
{:then list}
    {#if !list.length}
        <div class="alert alert-info" role="alert">
            Nothing posted on {(new Date(date * 1000)).toLocaleDateString()}!
        </div>
    {:else}
        <ul>
            {#each list as el}
                {#if el && el.oid && el.date}
                    <li on:click={showObject(el.oid)}>{new Date(el.date * 1000).toLocaleTimeString()}</li>
                {/if}
            {/each}
        </ul>
    {/if}
{:catch err}
    <div class="alert alert-danger" role="alert">
        {err}
    </div>
{/await}

<script>
import {view, token} from '../stores'

export let date = null

let listPromise = Promise.resolve([])
$: listPromise = loadList(date)
async function loadList(start) {
    if (!start) {
        return Promise.resolve([])
    }
    // Calculate the end of the day
    // We can't just add 24 hours because in most time zones, one day per year is 23 hours, and one is 25 (when DST starts and ends)
    const end = new Date(start * 1000)
    end.setDate(end.getDate() + 1)

    const reqBody = new FormData()
    reqBody.set('start', start)
    reqBody.set('end', end.getTime() / 1000)
    
    const response = await fetch(process.env.API_URL + '/search', {
        method: 'POST',
        body: reqBody,
        headers: {
            'Authorization': 'Bearer ' + $token
        }
    })

    if (response.status < 200 || response.status >= 400) {
        throw Error('Invalid response status code: ' + response.status)
    }

    const responseData = await response.json()
    if (!responseData) {
        throw Error('Response is empty')
    }

    return responseData
}

function showObject(oid) {
    $view = 'view/' + oid
}
</script>