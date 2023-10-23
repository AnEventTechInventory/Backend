
FROM arm64v8/alpine:latest
WORKDIR /app
COPY myapp /app/
EXPOSE 5678
CMD ["/app/myapp"]
ENV DB_USERNAME=$DB_USERNAME
ENV DB_PASSWORD=$DB_PASSWORD