module.exports = {
  pluginOptions: {
    autoRouting: {
      chunkNamePrefix: "page-"
    }
  },
  devServer: {
    port: 8080,
    proxy: "http://localhost:9300"
  }
};
