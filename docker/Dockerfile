FROM golang
ADD . /code
WORKDIR /code
RUN apt-get update
RUN apt-get -y upgrade
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends apt-utils
RUN apt-get install -y net-tools
RUN go get -u github.com/mchirico/go-facebook/...
RUN fgrab



# Build
# docker build --no-cache -t mchirico/go-facebook .

# Run
# docker run --rm -it mchirico/go-facebook fgrab

# Upload to repo
# docker push mchirico/go-facebook

# To get this file
# docker pull mchirico/go-facebook


