'use strict'

const ProductRegisterHook = exports = module.exports = {}
const GRPC = use('GRPC')
ProductRegisterHook.Register = async (modelInstance) => {
    const product = modelInstance.toJSON()
    GRPC.RegisterProduct({
        ID: product.id,
        Price: product.price_in_cents
    }, (err, resp) => {
        if(err) {
            console.log(err)
        }
        console.log(resp)
    })
}
