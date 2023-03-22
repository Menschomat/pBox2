FROM golang:1.19-alpine AS build

WORKDIR /src/
COPY *.go go.* /src/
COPY utils/ /src/utils/
COPY model/ /src/model/
RUN CGO_ENABLED=0 go build -o /bin/pbox

FROM scratch
COPY --from=build /bin/pbox /bin/pbox
ENTRYPOINT ["/bin/pbox"]