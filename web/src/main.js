import App from './App.svelte'

import {HandleSession} from './lib/Session'

import {profile, token} from './stores'

const app = (async function() {
    // Load profile and check session
    const [profileData, tokenData] = await HandleSession()
    profile.set(profileData || null)
    token.set(tokenData || null)

    // Crete a Svelte app by loading the main view
    window.svelteApp = new App({
        target: document.body
    })
})()

export default app
