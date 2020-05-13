<AddForm on:added={added}/>

<script>
import AddForm from './AddForm.svelte'

import {view} from '../stores'

function added(event) {
    if (event && event.detail && event.detail.objectId) {
        $view = 'view/' + event.detail.objectId
    }
}
</script>
