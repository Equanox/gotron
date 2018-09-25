var path = require("path");

var configuration = {
  /*
   * app.ts represents the entry point to your web application. Webpack will
   * recursively go through every "require" statement in app.ts and
   * efficiently build out the application's dependency tree.
   */
  entry: ["./src/app.tsx"],

  /*
   * The combination of path and filename tells Webpack what name to give to
   * the final bundled JavaScript file and where to store this file.
   */
  output: {
    path: path.resolve(__dirname, "build"),
    filename: "bundle.js"
  },

  /*
   * resolve lets Webpack now in advance what file extensions you plan on
   * "require"ing into the web application, and allows you to drop them
   * in your code.
   */
  resolve: {
    extensions: ['.ts', '.tsx', '.js', '.scss', '.json', '.css'],
    modules: [
      path.resolve(__dirname + '/src'),
      path.resolve(__dirname + '/node_modules')
    ]
  },

  module: {
    /*
     * Each loader needs an associated Regex test that goes through each
     * of the files you've included (or in this case, all files but the
     * ones in the excluded directories) and finds all files that pass
     * the test. Then it will apply the loader to that file. I haven't
     * installed ts-loader yet, but will do that shortly.
     */
    rules: [
      {
        test: /\.css$/,
        use: ['typings-for-css-modules-loader?modules&namedExport&camelCase'],
        exclude: /node_modules/,
      },
      {
        test: /\.(ttf)$/,
        loader: 'url-loader?limit=100000',
        exclude: /node_modules/,
      },
      {
        test: /\.(ts|tsx)$/,
        loader: "awesome-typescript-loader",
        exclude: /node_modules/,
      }
    ]
  }
};

module.exports = configuration;
