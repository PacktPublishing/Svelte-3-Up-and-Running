/* eslint-env mocha */

const assert = require('assert')

describe('sample test', function() {
    // Do not call .end() at the end of each test as we need to continue the same session

    // Increase timeouts
    this.slow(2000)
    this.timeout(3000)

    // Bail after a test fails
    this.bail(true)

    it('authenticate', (browser) => {
        browser
            .url('http://localhost:5000/')
            .waitForElementVisible('body', 1000)
            // We should be redirected to the auth server
            .assert.title('Svelte 3 Up and Running API Server')
            .assert.visible('input[id="auth-username"]')
            .assert.visible('input[id="auth-password"]')
            .assert.visible('button[type="submit"]')
            // Submit the form
            .clearValue('input[id="auth-username"]')
            .clearValue('input[id="auth-password"]')
            .setValue('input[id="auth-username"]', 'svelte')
            .setValue('input[id="auth-password"]', 'svelte')
            .click('button[type="submit"]')
            .pause(500, () => {
                // Ensure we are redirected to the app
                browser
                    .waitForElementVisible('body', 1000)
                    .assert.title('Svelte Journal')
                    .url((url) => {
                        assert(url)
                        assert(url.value)
                        assert.equal(url.value, 'http://localhost:5000/#')
                    })
            })
    })

    it('navigate to add form', (browser) => {
        browser
            .assert.visible('a[href="#/add"]')
            .click('a[href="#/add"]')
            .pause(50, () => {
                browser.url((url) => {
                    assert(url)
                    assert(url.value)
                    assert.equal(url.value, 'http://localhost:5000/#/add')
                })
            })
    })

    it('add form renders', (browser) => {
        browser
            .waitForElementVisible('body', 1000)
            .assert.visible('input[id="addform-title"]')
            .assert.visible('textarea[id="addform-content"]')
            .assert.visible('button[type="submit"]')
    })

    it('submit add form', (browser) => {
        browser
            .clearValue('input[id="addform-title"]')
            .clearValue('textarea[id="addform-content"]')
            .setValue('input[id="addform-title"]', 'Test entry')
            .setValue('textarea[id="addform-content"]', 'This is **bold**')
            .click('button[type="submit"]')
            .pause(500, () => {
                // Ensure we are redirected to the page showing the entry
                browser.url((url) => {
                    assert(url)
                    assert(url.value)
                    assert.match(url.value, /http:\/\/localhost:5000\/#\/view\/[0-9a-f-]{36}/)
                })
            })
    })

    // This test is just to end the session in the browser
    it('end session', (browser) => {
        browser.end()
    })
})
