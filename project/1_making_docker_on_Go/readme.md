## Building a basic Go docker file

- Create a newfolder and `cd` into it
- Create a `main.go` and `Dockerfile`
- Insert your golang logic in main.go
- Edit the docker file and insert the following (explanation of code is in the Dockerfile)
```Dockerfile
FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

```
- Use the following command to build an image. (Can change `golang-hi-server` of your choice) 
```bash
docker build -t golang-hi-server .
```
- Check if image is made
```bash
docker images
```
- Run the docker image created using (`golang-hi-server` is the name of img)
```bash
docker run -d -p 8080:8082 -it golang-hi-server
```
- open `http://localhost/8080` to see the project in action