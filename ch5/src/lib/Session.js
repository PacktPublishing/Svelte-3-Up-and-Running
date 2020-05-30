import IdTokenVerifier from 'idtoken-verifier'

/**
 * Nonce is a class that helps with generating and verifying nonce's
 */
class Nonce {
    constructor() {
        this._nonceKeyName = process.env.KEY_STORAGE_PREFIX + '-nonce'
        this._nonceLength = 7
    }

    /**
     * Generates a new nonce and stores it in the session storage
     *
     * @returns {string} A nonce
     */
    generate() {
        // Generate a nonce
        const nonce = RandomString(this._nonceLength)

        // Store the nonce in the session
        window.sessionStorage.setItem(this._nonceKeyName, nonce)

        return nonce
    }

    /**
     * Retrieves the last nonce from session storage
     *
     * @returns {string} A nonce
     */
    retrieve() {
        const read = window.sessionStorage.getItem(this._nonceKeyName)

        const regExp = new RegExp('^[A-Za-z0-9_\\-]{' + this._nonceLength + '}$')
        if (!read || !read.match(regExp)) {
            return null
        }

        return read
    }
}

/**
 * AuthenticationAttempts is a class that helps with managing authentication attempts
 */
export class AuthenticationAttempts {
    constructor() {
        this._attemptsKeyName = process.env.KEY_STORAGE_PREFIX + '-attempts'
    }

    getAttempts() {
        return parseInt((window.sessionStorage.getItem(this._attemptsKeyName) || 0), 10)
    }

    increaseAttempts() {
        const attempts = this.getAttempts()
        window.sessionStorage.setItem(this._attemptsKeyName, attempts + 1)
        return attempts
    }

    resetAttempts() {
        window.sessionStorage.setItem(this._attemptsKeyName, 0)
    }
}

/**
 * Credentials is a class that helps with managing credentials
 */
export class Credentials {
    constructor() {
        this._sessionKeyName = process.env.KEY_STORAGE_PREFIX + '-jwt'
        this._tokenValidated = false
        this._nonce = new Nonce()
        this._profile = null
    }

    /**
     * Returns the authorization URL to point users to, storing the nonce
     *
     * @returns {string} Authorization URL
     */
    authorizationUrl() {
        if (!process.env.AUTH_URL) {
            throw Error('Empty AUTH_URL env variable')
        }
        if (!process.env.AUTH_CLIENT_ID) {
            throw Error('Empty AUTH_CLIENT_ID env variable')
        }

        // Return URL
        // Use APP_URL env var or fallback to the current URL
        let appUrl
        if (process.env.APP_URL) {
            appUrl = process.env.APP_URL
        }
        else {
            // Remove the fragment from the URL
            let href = window.location.href
            if (window.location.hash && window.location.hash.length) {
                href = href.slice(0, -1 * window.location.hash.length)
            }
            if (href.endsWith('#')) {
                href = href.slice(0, -1)
            }
            // Remove index.html from the end if present
            if (href.endsWith('index.html')) {
                href = href.slice(0, -10)
            }
            appUrl = href
        }

        // Ensure it ends with a /
        appUrl += appUrl.endsWith('/') ? '' : '/'

        // Generate a nonce
        const nonce = this._nonce.generate()

        // Generate the URL
        const map = {
            '{clientId}': process.env.AUTH_CLIENT_ID,
            '{nonce}': nonce,
            '{appUrl}': encodeURIComponent(appUrl)
        }
        return process.env.AUTH_URL.replace(
            new RegExp(Object.keys(map).join('|'), 'g'),
            (matched) => map[matched]
        )
    }
    /**
     * Returns the profile object from the JWT token
     *
     * @returns {Object} Profile for the authenticated user
     * @async
     */
    async getProfile() {
        // If we have a pre-parsed and pre-validated profile in memory, return that
        if (this._profile) {
            return this._profile
        }

        // Get the token
        const jwt = this.getToken()
        if (!jwt) {
            return {}
        }

        // Get the profile out of the token
        try {
            let profile = await this._validateToken(jwt)
            if (!profile) {
                profile = {}
            }
            this._profile = profile
            return profile
        }
        catch (e) {
            this._profile = {}
            throw e
        }
    }

    /**
     * Returns the JWT token for the session
     *
     * @returns {string|null} JWT Token, or null if no token
     */
    getToken() {
        const read = window.sessionStorage.getItem(this._sessionKeyName)
        if (!read || !read.length) {
            return null
        }

        let data
        try {
            data = JSON.parse(read)
        }
        catch (error) {
            /* eslint-disable-next-line no-console */
            console.error('Error while parsing JSON from sessionStorage', error)
            throw Error('Could not get the token from session storage')
        }

        if (!data || !data.jwt) {
            return null
        }

        return data.jwt
    }

    /**
     * Stores the JWT token for the session
     *
     * @param {string} jwt - JWT Token
     * @async
     */
    async setToken(jwt) {
        // Delete the profile in memory
        this._profile = null

        // First, validate the token
        const profile = await this._validateToken(jwt)
        if (!profile) {
            throw Error('Token validation failed')
        }

        // Store the token
        window.sessionStorage.setItem(this._sessionKeyName, JSON.stringify({jwt}))

        // Set the profile in memory
        this._profile = profile
    }

    /**
     * Validates a token
     *
     * @param {string} jwt - JWT token to validate
     * @returns {Object} Extracted payload
     * @private
     */
    async _validateToken(jwt) {
        if (!process.env.AUTH_ISSUER) {
            throw Error('Empty AUTH_ISSUER env variable')
        }
        if (!process.env.AUTH_CLIENT_ID) {
            throw Error('Empty AUTH_CLIENT_ID env variable')
        }
        if (!process.env.AUTH_JWKS_URL) {
            throw Error('Empty AUTH_JWKS_URL env variable')
        }

        // Check the format
        if (!jwt) {
            throw Error('Invalid token')
        }

        // Validate the token
        const verifier = new IdTokenVerifier({
            jwksURI: process.env.AUTH_JWKS_URL,
            issuer: process.env.AUTH_ISSUER,
            audience: process.env.AUTH_CLIENT_ID
        })
        const payload = await new Promise((resolve, reject) => {
            verifier.verify(jwt, this._nonce.retrieve(), (error, payload) => {
                if (error) {
                    // eslint-disable-next-line no-console
                    console.error('Validation error', error)
                    reject(error)
                }

                resolve(payload)
            })
        })

        return payload
    }
}

/**
 * Returns a random string, useful for example as nonce.
 *
 * @param {number} length - Length of the string
 * @returns {string} Random string
 */
export function RandomString(length = 7) {
    const bytes = new Uint8Array(length)
    const random = window.crypto.getRandomValues(bytes)
    const result = []
    const charset = '0123456789ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvwxyz-_'
    random.forEach((c) => {
        result.push(charset[c % charset.length])
    })
    return result.join('')
}

/**
 * Export the credentials singleton
 */
export const credentials = new Credentials()

/**
 * Export the attempts singleton
 */
export const attempts = new AuthenticationAttempts()

/**
 * Main function that handles sessions: checks if there's an id_token in the page's URL, parses and validates it.
 * If there's no session, redirects the user to the authentication URL.
 *
 * @param {number} [maxAttempts=1] - Maximum number of attempts before stopping redirecting users; set to 0 to disable (e.g. always redirect)
 * @returns {[Object,string]} An array where the first element is the dictionary with the user's profile, made of all the claims included in the JWT, and the second element is the raw JWT
 * @async
 */
export async function HandleSession(maxAttempts) {
    // Default value
    if (maxAttempts === undefined || maxAttempts === null) {
        maxAttempts = 1
    }

    // Get the profile
    const profile = await GetProfile()

    // If we're not authenticated, and this is the first attempt, automatically redirect users
    if (!profile) {
        if (maxAttempts < 1 || attempts.increaseAttempts() < maxAttempts) {
            window.location.href = credentials.authorizationUrl()
            return [null, '']
        }
        else {
            console.error('No profile returned, but do not redirect because maximum attempts have passed')
        }
    }

    const token = credentials.getToken()
    return [profile, token]
}

/**
 * Returns the profile of the user based on the JWT id_token.
 * Returns false if there's no profile, or if the JWT is invalid.
 *
 * @returns {Object} The dictionary with the user's profile, made of all the claims included in the JWT
 * @async
 */
async function GetProfile() {
    // Check if we have an id_token
    if (window.location.hash) {
        const matchIdToken = window.location.hash.match(/id_token=(([A-Za-z0-9\-_~+/]+)\.([A-Za-z0-9\-_~+/]+)(?:\.([A-Za-z0-9\-_~+/]+)))/)
        if (matchIdToken && matchIdToken[1]) {
            // First thing: remove the token from the URL (for security)
            history.replaceState(undefined, undefined, '#')

            // Validate and store the JWT
            // If there's an error, redirect to auth page
            try {
                // Set (and validate) the JWT
                await credentials.setToken(matchIdToken[1])

                // Reset attempts counter
                attempts.resetAttempts()
            }
            catch (error) {
                /* eslint-disable-next-line no-console */
                console.error('Token error', error)

                return false
            }
        }
    }

    // If we don't have credentials stored, redirect the user to the authentication page
    if (!credentials.getToken()) {
        return false
    }

    // Get the profile
    // If there's no session or it has expired, redirect to auth page
    let profile = null
    try {
        profile = await credentials.getProfile()
    }
    catch (error) {
        /* eslint-disable-next-line no-console */
        console.error('Token error', error)

        return false
    }

    return profile
}
