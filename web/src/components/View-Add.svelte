<div class="row justify-content-md-center">
    <div class="col-6">
        <AddForm on:added={added}/>
    </div>
</div>

<script>
import AddForm from './AddForm.svelte'

import {view} from '../stores'

function added(event) {
    if (event && event.detail && event.detail.objectId) {
        $view = 'view/' + event.detail.objectId
    }
}
</script>
