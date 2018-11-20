const merge = require('webpack-merge');
const common = require('./webpack.common.js');
const Uglify = require('uglifyjs-webpack-plugin');

module.exports = merge(common, {
  mode: 'production',
  optimization: {
    minimizer: [
      new Uglify({
        uglifyOptions: {
          mangle: true,
          compress: {
            drop_console: true,
            conditionals: true,
            unused: true,
            comparisons: true,
            dead_code: true,
            if_return: true,
            join_vars: true,
            warnings: false,
          },
          output: {
            comments: false,
          },
          exclude: [/\.min\.js$/gi],
        },
      }),
    ],
  },
});
