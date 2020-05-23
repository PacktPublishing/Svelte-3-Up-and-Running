{#await listPromise}
    Loading…
{:then list}
    <h2 class="text-2xl font-bold mb-2 text-gray-800">Journal entries on {day}</h2>
    {#if !list.length}
        <div class="bg-blue-100 border-l-4 border-blue-500 text-blue-700 p-4 my-2 mx-6">
            Nothing posted on {day}!
        </div>
    {:else}
        <ul class="ml-6 space-y-2">
            {#each list as el}
                {#if el && el.oid && el.date}
                    <li class="cursor-pointer bg-white shadow py-2 px-4 w-2/3 lg:w-3/5" on:click={showObject(el.oid)}>
                        {new Date(el.date * 1000).toLocaleTimeString([], {hour: '2-digit', minute: '2-digit'})}
                        <b>{el.title || '(no title)'}</b>
                    </li>
                {/if}
            {/each}
        </ul>
    {/if}
{:catch err}
    <ErrorBox {err} />
{/await}

<script>
import ErrorBox from './ErrorBox.svelte'
import {view, token} from '../stores'

export let date = null

$: day = (new Date(date * 1000)).toLocaleDateString()

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