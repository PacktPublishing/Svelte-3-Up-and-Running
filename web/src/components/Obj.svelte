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

import {token} from '../stores'

export let objectId = null
let contentPromise = Promise.resolve('')
$: contentPromise = loadObject(objectId)

async function loadObject(oid) {
    const response = await fetch(process.env.API_URL + '/object/' + oid, {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + $token
        }
    })

    if (response.status < 200 || response.status >= 400) {
        throw Error('Invalid response status code: ' + response.status)
    }

    const responseText = await response.text()
    if (!responseText) {
        throw Error('Response is empty')
    }

    const date = response.headers.get('x-object-date')
    const title = response.headers.get('x-object-title')

    return {
        date,
        title,
        content: responseText
    }
}
</script>
