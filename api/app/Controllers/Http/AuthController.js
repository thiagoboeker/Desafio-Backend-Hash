'use strict'
const Model = use('App/Models/User')
class AuthController {

    // Registra um novo Usuario na tabela User
    async register({ request, response }) {
        const body = request.only(['username', 'email', 'password', "birthdate"])
        const resp = await Model.create(body)
        return response.status(201).send({
            success: true,
        })
    }

    // Loga um usuario, retorna um JWT na resposta 
    async login({ request, response, auth }) {
        const body = request.only(['user'])
        const token = await auth.authenticator('jwt').attempt(body.user.email, body.user.password)
        return response.status(302).send({
            success: true,
            token: token
        })
    }
}

module.exports = AuthController
