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

import {token} from '../stores'

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
    const formData = new FormData()
    formData.append('title', title)
    formData.append('content', content)
    try {
        const response = await fetch(process.env.API_URL + '/object', {
            method: 'POST',
            body: formData,
            headers: {
                'Authorization': 'Bearer ' + $token
            }
        })

        if (response.status < 200 || response.status >= 400) {
            throw Error('Invalid response status code: ' + response.status)
        }

        const responseData = await response.json()
        if (!responseData || !responseData.objectId) {
            throw Error('Invalid response: no objectId')
        }

        dispatch('added', {objectId: responseData.objectId})
    }
    catch (err) {
        console.error(err)
        formError = 'Request error: ' + err
    }

    running = false
}
</script>
