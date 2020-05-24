<form class="bg-white shadow rounded px-8 pt-6 pb-8 mb-4" on:submit|preventDefault={submit}>
    <div class="mb-4">
        <label for="addform-title" class="block text-gray-700 font-bold mb-2">Title:</label>
        <input type="text" class="shadow-inner appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" bind:value={title} />
    </div>
    <div class="mb-4">
        <label for="addform-content" class="block text-gray-700 font-bold mb-2">Content:</label>
        <textarea id="addform-content" class="shadow-inner appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline textarea-tall" bind:value={content}></textarea>
    </div>
    <button type="submit" class="bg-blue-500 hover:bg-blue-700 text-white py-1 px-3 rounded focus:outline-none focus:shadow-outline" disabled={running}>Save</button>
    {#if formError}
        <div class="mt-2 text-sm text-red-600">
            {formError}
        </div>
    {/if}
</form>

<style>
.textarea-tall {
    height: calc(80vh - 14em);
}
</style>

<script>
import {createEventDispatcher} from 'svelte'
import {AddRequest} from '../lib/Requests.js'
import {token} from '../stores.js'

export let content
export let title

const dispatch = createEventDispatcher()

let formError = null
let running = false
async function submit() {
    // Semaphore
    if (running) {
        return
    }
    running = true

    // Errors
    if (!content) {
        formError = 'Please write some content'
        running = false
        return
    }
    formError = null

    // Send request
    try {
        const objectId = await AddRequest(title, content, $token)
        dispatch('added', {objectId})
    }
    catch (err) {
        console.error(err)
        formError = 'Request error: ' + err
    }

    running = false
}
</script>
