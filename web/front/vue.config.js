module.exports = {
  publicPath: '/static/',
  outputDir: 'dist',
  transpileDependencies: ['vuetify'],
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = '欢迎来到TsuhaoBlog'
      return args
    })
  },
  productionSourceMap: false
}
