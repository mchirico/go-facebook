#!/bin/bash
# git archive --format=tar --prefix=go-facebook/ HEAD ../| gzip > go-facebook.tar.gz
docker build --no-cache -t mchirico/go-facebook .
if [ "$(whoami)" != "mchirico" ]; then
        echo "You need permissions to push to repo"
        exit 1
fi
docker run --rm -it mchirico/go-facebook fgrab  
docker push mchirico/go-facebook
