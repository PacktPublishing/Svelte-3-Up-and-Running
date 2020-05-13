import {writable, derived} from 'svelte/store'

// Contains the user's profile
export const profile = writable(null)

// Contains the token
export const token = writable(null)

// Derived store that returns true if the user is authenticated
export const isAuthenticated = derived([token, profile], (a) => a && a[0] && a[1])

// Current view
export const view = writable(null)
