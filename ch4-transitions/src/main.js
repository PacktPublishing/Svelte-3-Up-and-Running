import './global.css'
import {HandleSession} from './lib/Session.js'
import {profile, token} from './stores'
import App from './App.svelte'

const app = (async function() {
    const [profileData, tokenData] = await HandleSession(0)
    profile.set(profileData || null)
    token.set(tokenData || null)

    return new App({
        target: document.body
    })
})()
export default app
