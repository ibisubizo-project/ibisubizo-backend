const path = require('path');
const webpack = require('webpack');


const ExtractTextPlugin = require("extract-text-webpack-plugin")

module.exports = {
    entry: './src/index.js',
    module: {
        rules: [
            {
                test:/\.(js|jsx)$/,
                exclude: /node_modules/,
                use: ['babel-loader']
            },
            {
                test: /\.css$/,
                loader: ExtractTextPlugin.extract({
                    use: 'css-loader',
                })
            },
            {
                test: /\.(png|gif|jpg|jpeg)/,
                use: [
                    'file-loader'
                ]
            }
        ]
    },
    resolve:{
        extensions: ['*', '.js', '.jsx']
    },
    output: {
        path: path.join(__dirname, 'dist'),
        publicPath: '/',
        filename: 'bundle.js'
    },
    plugins:[
        new webpack.HotModuleReplacementPlugin(),
        new ExtractTextPlugin('style.bundle.css')
    ],
    devServer: {
        contentBase: './dist',
        hot: true
    }
}