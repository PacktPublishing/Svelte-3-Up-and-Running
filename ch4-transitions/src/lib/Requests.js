/**
 * Requests an object from the server
 * @param {string} oid - ID of the object to request
 * @param {string} token - Auth token
 * @returns {Object} Object's data, which is a dictionary with the date, title and content properties
 * @async
 */
export async function LoadObject(oid, token) {
    // Send the request to the server
    const response = await fetch(process.env.API_URL + '/object/' + oid, {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + token
        }
    })
    if (response.status < 200 || response.status >= 400) {
        throw Error('Invalid response status code: ' + response.status)
    }

    // Parse the response
    const responseText = await response.text()
    if (!responseText) {
        throw Error('Response is empty')
    }

    // Get the date and title from the response headers
    const date = response.headers.get('x-object-date')
    const title = response.headers.get('x-object-title')

    return {
        date,
        title,
        content: responseText
    }
}

/**
 * Requests the list of objects for a given day.
 * @param {number} start - Start time, as UNIX timestamp
 * @param {string} token - Auth token
 * @returns {Array} List of objects as returned by the server
 * @async
 */
export async function LoadList(start, token) {
    // No date
    if (!start) {
        return Promise.resolve([])
    }

    // Remove all decimals
    start = parseInt(start, 10)

    // Calculate the end of the day
    // We can't just add 24 hours because in most time zones, one day per year is 23 hours, and one is 25 (when DST starts and ends)
    const end = new Date(start * 1000)
    end.setDate(end.getDate() + 1)

    // Request body: a FormData object with start and end times
    const reqBody = new FormData()
    reqBody.set('start', start)
    reqBody.set('end', end.getTime() / 1000)

    // Make the request to the server
    const response = await fetch(process.env.API_URL + '/search', {
        method: 'POST',
        body: reqBody,
        headers: {
            'Authorization': 'Bearer ' + token
        }
    })
    if (response.status < 200 || response.status >= 400) {
        throw Error('Invalid response status code: ' + response.status)
    }

    // Parse the response as JSON and return it
    const responseData = await response.json()
    if (!responseData) {
        throw Error('Response is empty')
    }

    return responseData
}

/**
 * Sends a request to create a new object.
 * @param {string} title - Object title
 * @param {string} content - Object content
 * @param {string} token - Auth token
 * @returns {string} Object ID of the created object
 * @async
 */
export async function AddRequest(title, content, token) {
    // Request body: a FormData object with the title and content
    const formData = new FormData()
    formData.append('title', title)
    formData.append('content', content)

    // Send the request to the server
    const response = await fetch(process.env.API_URL + '/object', {
        method: 'POST',
        body: formData,
        headers: {
            'Authorization': 'Bearer ' + token
        }
    })
    if (response.status < 200 || response.status >= 400) {
        throw Error('Invalid response status code: ' + response.status)
    }

    // Parse the response as JSON
    const responseData = await response.json()
    if (!responseData || !responseData.objectId) {
        throw Error('Invalid response: no objectId')
    }
    return responseData.objectId
}
