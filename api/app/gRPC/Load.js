const grpc = require('grpc')
const loader = require('@grpc/proto-loader')
const { ioc } = require('@adonisjs/fold')

module.exports = () => {
    const path = __dirname + "/service.proto"
    const clientDefinition = loader.loadSync(path, {
        keepCase: true
    })
    const definition = grpc.loadPackageDefinition(clientDefinition)

    const client = new definition.Service('hash-server-svc.default.svc.cluster.local:9090', grpc.credentials.createInsecure())

    ioc.singleton('GRPC', (_) => {
        return client
    })
    return 'GRPC Client load successfully.'
}