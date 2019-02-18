'use strict'

const UserBirthdateHook = exports = module.exports = {}
const GRPC = use('GRPC')
UserBirthdateHook.RegisterBirthdate = async (modelInstance) => {
    const user = modelInstance.toJSON()
    GRPC.RegisterBirthdate({
        ID: user.id,
        Birthdate: user.birthdate
    }, (err, resp) => {
        if(err) {
            console.log(err)
        }
        console.log(resp)
    })
}
