FROM node:alpine

WORKDIR /app

EXPOSE 3333

COPY package*.json .

RUN npm i

COPY . .

CMD ["npm", "run", "dev"]