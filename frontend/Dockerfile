FROM node:lts-alpine as build-stage
ARG ENV
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run test:unit && npm run build-for-$ENV

FROM nginx:stable-alpine as production-stage
ARG ENV
RUN mkdir /app
COPY --from=build-stage /app/dist /app
COPY --from=build-stage /app/nginx-$ENV.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
