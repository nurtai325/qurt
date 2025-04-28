FROM nginx:stable-bookworm
COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY landing /usr/share/nginx/html
COPY assets /usr/share/nginx/html/assets
COPY editor /usr/share/nginx/html/editor
