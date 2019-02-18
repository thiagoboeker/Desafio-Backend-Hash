FROM node:11.0

RUN mkdir -p /app

COPY ./package.json /app

WORKDIR /app

RUN npm install

COPY ./ .

EXPOSE 3333

CMD ["npm", "start"]