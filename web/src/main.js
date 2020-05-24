import './global.css'

import {HandleSession} from './lib/Session.js'
import {profile, token} from './stores.js'

import App from './App.svelte'

const app = (async function() {
    // Load profile and check session
    const [profileData, tokenData] = await HandleSession(0)
    profile.set(profileData || null)
    token.set(tokenData || null)

    // Crete a Svelte app by loading the main view
    return new App({
        target: document.body
    })
})()

export default app
