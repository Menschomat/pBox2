FROM node:18-alpine AS build_dashboard
WORKDIR /src/
COPY web/ web/
WORKDIR /src/web/
RUN npm install
RUN npm run build

FROM golang:1.22.4-alpine AS build_backend
WORKDIR /src/
COPY *.go go.* /src/
COPY docs/ /src/docs/
COPY api/ /src/api/
COPY utils/ /src/utils/
COPY model/ /src/model/

RUN CGO_ENABLED=0 go build -o /bin/pbox



FROM scratch
COPY --from=build_backend /bin/pbox /src/pbox
COPY --from=build_dashboard /src/static/dashboard /src/static/dashboard
COPY config.json /src/config.json
WORKDIR /src/
ENTRYPOINT ["./pbox"]