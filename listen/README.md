# grpc-listen

proves it's possible to serve grpc request on the same port with REST gateway

## how-to check

build

    make && ./grpc-listen

send REST request on shared port:

    echo "{}" | http -v http://127.0.0.1:5556/topic

check grpc:

    grpcurl -plaintext 127.0.0.1:5556 api.Publisher.Topic


how to install [grpcurl](https://github.com/fullstorydev/grpcurl) on MacOSX:

    brew install grpcurl
