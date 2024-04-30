###########################################
# BASE IMAGE
###########################################
FROM golang:alpine AS build

WORKDIR /build
COPY . .
RUN go mod download 
EXPOSE 8080
RUN go build -o ./bin

############################################
# HERE STARTS THE MAGIC OF MULTI STAGE BUILD
############################################

FROM scratch

# Copy the compiled binary from the build stage
WORKDIR /app
COPY --from=build /build/bin ./bin

CMD ["/app/bin"]
