const {WebpackPluginServe} = require('webpack-plugin-serve')
const {CleanWebpackPlugin} = require('clean-webpack-plugin')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const Dotenv = require('dotenv-webpack')
const SriPlugin = require('webpack-subresource-integrity')
const path = require('path')

const mode = process.env.NODE_ENV || 'development'
const prod = mode === 'production'

const htmlMinifyOptions = {
    collapseWhitespace: true,
    conservativeCollapse: true,
    removeComments: true,
    collapseBooleanAttributes: true,
    decodeEntities: true,
    html5: true,
    keepClosingSlash: false,
    processConditionalComments: true,
    removeEmptyAttributes: true
}

// Entry points
const entry = {
    app: [path.resolve(__dirname, 'src/main.js')],
}
if (!prod) {
    // Required for webpack-plugin-serve (dev-only)
    entry.webpackServe = 'webpack-plugin-serve/client'
}

const addPlugins = []
if (prod) {
    // Enable the clean plugin in prod only
    addPlugins.push(new CleanWebpackPlugin({
        cleanOnceBeforeBuildPatterns: ['**/*', '!assets', '!assets/*']
    }))
}
else {
    // Enable the serve plugin in dev only
    addPlugins.push(new WebpackPluginServe({
        static: 'public',
        host: '0.0.0.0',
        port: 3000
    }))
}

module.exports = {
    entry,
    resolve: {
        mainFields: ['svelte', 'browser', 'module', 'main'],
        extensions: ['.mjs', '.js', '.svelte']
    },
    output: {
        path: path.resolve(__dirname, 'public'),
        filename: prod ? '[name].[hash:8].js' : '[name].js',
        chunkFilename: prod ? '[name].[contenthash:8].js' : '[name].js',
        crossOriginLoading: 'anonymous'
    },
    module: {
        rules: [
            {
                test: /\.(svelte)$/,
                exclude: [],
                use: {
                    loader: 'svelte-loader'
                }
            }
        ]
    },
    mode,
    plugins: [
        // Supports reading the .env file
        new Dotenv(),

        // Generate the index.html file
        new HtmlWebpackPlugin({
            filename: 'index.html',
            template: path.resolve(__dirname, 'src/main.html'),
            chunks: ['app'],
            minify: prod ? htmlMinifyOptions : false
        }),

        // Enable subresource integrity check
        new SriPlugin({
            hashFuncNames: ['sha384'],
            enabled: prod,
        })
    ].concat(addPlugins),
    watch: !prod,
    devtool: prod ? false : 'source-map'
}
