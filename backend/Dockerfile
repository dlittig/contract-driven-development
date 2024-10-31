FROM alpine:3.20

ENV TZ="Europe/Berlin"

USER root

COPY container_files/ /usr/bin

RUN apk --no-cache add openjdk17 python3 py3-pip wget; \
    wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.9.0/openapi-generator-cli-7.9.0.jar -O openapi-generator-cli.jar; \
    apk remove wget; \
    rm -rf /var/cache/apk/*; \
    pip install poetry --break-system-packages; \
    python --version; \
    pip --version; \
    poetry --version; \
    openapi-generator-cli version;