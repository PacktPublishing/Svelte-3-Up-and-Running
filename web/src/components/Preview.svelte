<div class="bg-white shadow rounded px-8 pt-6 pb-8 mb-4 text-sm">
    <h1 class="text-3xl mb-2">Preview</h1>
    <div class="border border-gray-600 p-4">
        <Renderer {title} {content} />
    </div>
</div>

<script>
import Renderer from './Renderer.svelte'

export let content
export let title
</script>
