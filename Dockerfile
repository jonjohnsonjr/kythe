# Assumptions:
#   You're sitting inside something like: $GOPATH/src/github.com/google/kythe
#   `../../knative/serving` exists.

# Build the public resources for the web UI
FROM openjdk as dist
COPY . /src
WORKDIR /src/kythe/web/ui
RUN apt-get update && apt-get install -y wget
RUN /src/third_party/leiningen/lein cljsbuild once prod

# Build all the kythe binaries
FROM openjdk as bins
COPY . /src
WORKDIR /src
RUN echo "deb [arch=amd64] http://storage.googleapis.com/bazel-apt stable jdk1.8" | \
    tee /etc/apt/sources.list.d/bazel.list &&  \
      curl https://bazel.build/bazel-release.pub.gpg | apt-key add -
RUN apt-get update && apt-get install -y bazel clang
RUN bazel build -c opt \
    //kythe/go/extractors/cmd/gotool \
    //kythe/go/indexer/cmd/go_indexer \
    //kythe/go/serving/tools/write_tables \
    //kythe/go/serving/tools/http_server
# Symlinks are bad, m'kay.
RUN mkdir /out && \
    cp /src/bazel-out/k8-opt/bin/kythe/go/extractors/cmd/gotool/linux_amd64_stripped/gotool /out/gotool && \
    cp /src/bazel-out/k8-opt/bin/kythe/go/indexer/cmd/go_indexer/linux_amd64_stripped/go_indexer /out/go_indexer && \
    cp /src/bazel-out/k8-opt/bin/kythe/go/serving/tools/write_tables/linux_amd64_stripped/write_tables /out/write_tables && \
    cp /src/bazel-out/k8-opt/bin/kythe/go/serving/tools/http_server/linux_amd64_stripped/http_server /out/http_server

# Index the code, producing the serving_table
FROM golang as indexer
COPY --from=bins /out/gotool /gotool
COPY --from=bins /out/go_indexer /go_indexer
COPY --from=bins /out/write_tables /write_tables
WORKDIR /go/src/app
COPY ../../knative/serving /go/src/app
COPY index.sh index.sh
RUN chmod +x index.sh && ./index.sh

# Final image, serves the index via web ui
FROM scratch
COPY --from=dist /src/kythe/web/ui/resources/public /public
COPY --from=bins /out/http_server /http_server
COPY /tmp/out /index --from indexer
EXPOSE 8080
CMD ['/http_server', '--listen', ':8080', '--public_resources', '/public', '--serving_table', '/index']
