{#await listPromise}
  Loading…
{:then list}
  <h2 class="text-2xl font-bold mb-2 text-gray-800">Journal entries on {day}</h2>
  {#if !list.length}
    <div class="bg-blue-100 border-l-4 border-blue-500 text-blue-700 p-4 my-2 mx-6">
      Nothing posted on {day}!
    </div>
  {:else}
    <div class="ml-6 space-y-2">
      {#each list as el}
        {#if el && el.oid && el.date}
          <a class="block bg-white shadow py-2 px-4 w-2/3 lg:w-3/5"
            href={'#/view/' + el.oid}>
              {new Date(el.date * 1000).toLocaleTimeString([], {hour: '2-digit', minute: '2-digit'})}
              <b>{el.title || '(no title)'}</b>
          </a>
        {/if}
      {/each}
    </div>
  {/if}
{:catch err}
  <ErrorBox {err} />
{/await}

<script>
import ErrorBox from './ErrorBox.svelte'
import {token} from '../stores.js'
import {LoadList} from '../lib/Requests.js'

export let date = null

$: day = (new Date(date * 1000)).toLocaleDateString()

let listPromise = Promise.resolve([])
$: listPromise = LoadList(date, $token)
</script>