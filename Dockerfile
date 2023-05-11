FROM alpine:latest

WORKDIR /app/

COPY bin/mongo-ping .

ENTRYPOINT [ "./mongo-ping" ]