FROM node:lts-alpine AS frontend-deps

COPY assets/package.json package.json
COPY assets/yarn.lock yarn.lock

RUN yarn install --frozen-lockfile

FROM node:lts-alpine AS frontend-builder

COPY assets/ .
COPY --from=frontend-deps node_modules node_modules

RUN yarn build

FROM node:lts-alpine AS frontend-embed

COPY --from=frontend-builder build build

RUN apk update && apk upgrade
RUN apk add zip
RUN zip -q assets.zip -r build

FROM golang:alpine AS go-builder

WORKDIR /app

RUN apk add build-base

COPY . .

RUN rm -rf assets
COPY --from=frontend-embed assets.zip .

RUN go build -a -o randomdonate .

FROM alpine AS runner

WORKDIR /app

COPY --from=go-builder /app/randomdonate .

VOLUME /app
EXPOSE 8080

RUN chmod +x /app/randomdonate
CMD /app/randomdonate
