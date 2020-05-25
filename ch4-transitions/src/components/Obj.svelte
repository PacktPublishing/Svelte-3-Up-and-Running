{#await contentPromise}
  Loading…
{:then el}
  <p class="mb-2 italic">
    Saved on
    {(el && el.date) ? new Date(el.date).toLocaleString() : '(null)'}
  </p>
  <Renderer title={el && el.title || ''} content={el && el.content || ''} />
{:catch err}
  <ErrorBox {err} />
{/await}

<script>
import Renderer from './Renderer.svelte'
import ErrorBox from './ErrorBox.svelte'
import {LoadObject} from '../lib/Requests.js'
import {token} from '../stores.js'

export let objectId = null
let contentPromise = Promise.resolve({})
$: contentPromise = LoadObject(objectId, $token)
</script>
