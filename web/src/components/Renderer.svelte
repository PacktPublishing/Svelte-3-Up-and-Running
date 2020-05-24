{#if title}
    <h1 class="text-3xl mb-3">{title}</h1>
{:else}
    <h1 class="text-3xl mb-3 text-gray-600">No title</h1>
{/if}
<div class="rendered">
    {#if content}
        {@html rendered}
    {:else}
        <p class="text-gray-600">No content</p>
    {/if}
</div>

<script>
// Props for the view
export let title = ''
export let content = ''

// Markdown renderer
import MarkdownIt from 'markdown-it'
const markdown = new MarkdownIt()

$: rendered = markdown.render(content)
</script>
