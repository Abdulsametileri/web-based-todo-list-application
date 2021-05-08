FROM nginx:stable-alpine as production-stage
RUN mkdir /app
COPY dist /app
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
