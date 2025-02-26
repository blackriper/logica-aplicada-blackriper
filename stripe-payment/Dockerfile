#first stage 
FROM golang:alpine AS builder  

RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY go.mod go.sum ./ 
RUN go mod download

COPY  . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /stripe-store 


#second stage 
FROM scratch 

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /stripe-store /stripe-store

COPY . .

EXPOSE 3000

CMD [ "/stripe-store" ]
