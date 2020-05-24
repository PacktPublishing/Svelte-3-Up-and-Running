{#await contentPromise}
    Loading…
{:then el}
    <p class="mb-2 italic">Saved on {new Date(el.date).toLocaleString()}</p>
    <Renderer title={el.title} content={el.content} />
{:catch err}
    <ErrorBox {err} />
{/await}

<script>
import Renderer from './Renderer.svelte'
import ErrorBox from './ErrorBox.svelte'
import {LoadObject} from '../lib/Requests.js'
import {token} from '../stores.js'

export let objectId = null
let contentPromise = Promise.resolve('')
$: contentPromise = LoadObject(objectId, $token)
</script>
