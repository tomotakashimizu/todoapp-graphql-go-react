FROM node:18.16.1-alpine3.18
WORKDIR /front
RUN apk add --update make g++ python3
COPY package*.json ./
RUN npm install
COPY . .
CMD ["npm", "start"]