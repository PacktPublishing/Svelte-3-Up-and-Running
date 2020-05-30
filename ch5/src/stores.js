import {writable, derived} from 'svelte/store'
export const profile = writable(null)
export const token = writable(null)
export const isAuthenticated = derived([token, profile], (a) => a && a[0] && a[1])
export const view = writable(null)
