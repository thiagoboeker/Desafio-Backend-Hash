'use strict'

/** @typedef {import('@adonisjs/framework/src/Request')} Request */
/** @typedef {import('@adonisjs/framework/src/Response')} Response */
/** @typedef {import('@adonisjs/framework/src/View')} View */

/**
 * Resourceful controller for interacting with products
 */
const Model = use('App/Models/Product')
const GRPC = use('GRPC')
const bluebird = require('bluebird')
const client = bluebird.promisifyAll(GRPC)
const _ = require('lodash')
class ProductController {

  // Retorna a lista de produtos
  async index ({ request, response, view, auth }) {
    // Pega todos os produtos
    let Products = await Model.query().fetch()
    Products = Products.toJSON()

    // Array de promises para pegar o disconto de cada produto
    // TODO: Consulta em forma de batch
    const product_with_discounts = await bluebird.all(_.map(Products, async prod => {
      let discount;
      try {
        discount = await client.GetDiscountAsync({
          ProductID: prod.id,
          UserID: auth.user.$attributes.id,
        })
      }
      //TODO: tratamento de erro 
      catch(err) {
        console.log(err)
        discount = null;
      }
      return {
        ...prod,
        discount: discount,
      }
    }))
    return response.status(200).send({
      success: true,
      data: product_with_discounts
    })
  }

  // CRUD - Inseri um novo produto
  async store ({ request, response }) {
    const body = request.only(['price_in_cents', 'title', 'description'])
    const resp = await Model.create(body)
    return response.status(201).send({
      success: true,
      data: resp
    })
  }

  // CRUD - Retorna um produto especifico.
  async show ({ params, request, response, view }) {
    const resp = await Model.findOrFail(params.id)
    return response.status(302).send({
      success: true,
      data: resp
    })
  }
  
  // CRUD - Update
  async update ({ params, request, response }) {
    const body = request.only(['price_in_cents', 'title', 'description'])
    const product = await Model.findOrFail(params.id)
    product.merge(body)
    await product.save()
    return response.status(200).send({
      success: true,
      data: resp
    })
  }

  // CRUD - Delete
  async destroy ({ params, request, response }) {
    const product = await Model.query().where({id: params.id}).delete()
    return response.status(200).send({
      success: true,
      data: product
    })
  }
}

module.exports = ProductController
