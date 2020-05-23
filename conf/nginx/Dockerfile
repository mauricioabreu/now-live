FROM nginx:alpine

RUN apk add --no-cache curl

RUN \
  rm -f \
    /etc/nginx/sites-available/minio.conf \
    /etc/nginx/sites-enabled/minio.conf \
    /etc/nginx/sites-enabled/default

ADD sites-available/ /etc/nginx/sites-available

COPY nginx.conf /etc/nginx/nginx.conf

RUN \
  mkdir -p /etc/nginx/sites-enabled && \
  ln -s /etc/nginx/sites-available/minio.conf /etc/nginx/sites-enabled/minio.conf