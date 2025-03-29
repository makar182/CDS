FROM node:18

WORKDIR /app

COPY package.json ./
RUN npm install

COPY . .

RUN npm run build

EXPOSE 3000
RUN chmod +x /app/frontend-client-data
CMD ["npm", "start"]
