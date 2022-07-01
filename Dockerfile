FROM node:lts-alpine AS frontend-deps

COPY assets/package.json package.json
COPY assets/yarn.lock yarn.lock

RUN yarn install --frozen-lockfile

FROM node:lts-alpine AS frontend-builder

COPY assets/ .
COPY --from=frontend-deps node_modules node_modules

RUN yarn export

FROM node:lts-alpine AS frontend-embed

COPY --from=builder out out

RUN apk update && upgrade
RUN apk add zip
RUN zip -q out.zip -r out

FROM golang:alpine AS go-builder

COPY . .

RUN rm -rf assets && mkdir assets
COPY --from=frontend-embed out.zip assets/out.zip

RUN go build -a -o randomdonate .

FROM alpine AS runner

COPY --from=go-builder randomdonate .
COPY conf.ini .
COPY random_donate.db .

EXPOSE 8080

RUN chmod +x randomdonate
CMD ./randomdonate
