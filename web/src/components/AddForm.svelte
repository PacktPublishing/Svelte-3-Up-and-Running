<form on:submit|preventDefault={submit}>
    <div class="form-group">
        <label for="addform-content">Content:</label>
        <textarea class="form-control" id="addform-content" rows="4" bind:value={content}></textarea>
    </div>
    <button type="submit" class="btn btn-primary" disabled={running}>Save</button>
    {#if formError}
        <div class="alert alert-danger" role="alert">
            {formError}
        </div>
    {/if}
</form>

<script>
import {createEventDispatcher} from 'svelte'

import {token} from '../stores'

const dispatch = createEventDispatcher()

let content
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
