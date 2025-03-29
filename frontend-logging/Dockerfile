FROM node:18

WORKDIR /app

COPY package.json ./
RUN npm install

COPY . .

RUN npm run build

EXPOSE 3001
RUN chmod +x /app/frontend-logging
CMD ["npm", "start"]
