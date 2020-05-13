{#await contentPromise}
    Loading…
{:then content}
    {@html content}
{:catch err}
    <ErrorBox {err} />
{/await}

<script>
import ErrorBox from './ErrorBox.svelte'

import MarkdownIt from 'markdown-it'
import {token} from '../stores'

const markdown = new MarkdownIt()

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

    const responseHtml = markdown.render(responseText)

    return responseHtml
}
</script>
