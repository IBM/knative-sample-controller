FROM debian:stretch-slim
RUN apt-get update && apt-get install --no-install-recommends -y ca-certificates && rm -rf /var/lib/apt/lists/*
ADD ./sample-controller /usr/bin/main
CMD ["/usr/bin/main"]