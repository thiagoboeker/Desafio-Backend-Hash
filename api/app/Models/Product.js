'use strict'

/** @type {typeof import('@adonisjs/lucid/src/Lucid/Model')} */
const Model = use('Model')

class Product extends Model {
    static boot() {
        super.boot()
        this.addHook('afterSave', 'ProductRegisterHook.Register')
    }
}

module.exports = Product
