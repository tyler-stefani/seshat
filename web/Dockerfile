FROM node:lts-alpine3.16

RUN mkdir /app

WORKDIR /app

COPY package*.json .

RUN npm ci

COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev"]
