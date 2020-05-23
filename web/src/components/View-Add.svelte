<AddForm on:added={added} bind:content bind:title />

<Preview {title} {content} />

<script>
import AddForm from './AddForm.svelte'
import Preview from './Preview.svelte'

import {view} from '../stores'

let content = ''
let title = ''

function added(event) {
    if (event && event.detail && event.detail.objectId) {
        $view = 'view/' + event.detail.objectId
    }
}
</script>
