<div in:fade>
    <AddForm on:added={added} bind:content bind:title />
    <div class="bg-white shadow rounded px-8 pt-6 pb-8 mb-4 text-sm">
        <h1 class="text-3xl mb-2">Preview</h1>
        <div class="border border-gray-600 p-4">
            <Renderer {title} {content} />
        </div>
    </div>
</div>

<script>
import {fade} from 'svelte/transition'
import {push} from 'svelte-spa-router'
import AddForm from './AddForm.svelte'
import Renderer from './Renderer.svelte'

let content = ''
let title = ''
function added(event) {
    if (event && event.detail && event.detail.objectId) {
        push('/view/' + event.detail.objectId)
    }
}
</script>
