FROM node:alpine

WORKDIR /app

COPY package*.json ./
COPY vite.config.ts ./
COPY . .

RUN npm install --loglevel verbose

COPY . .

EXPOSE 3080

CMD [ "sh", "-c", "cd /app && npm i && npm run dev" ]