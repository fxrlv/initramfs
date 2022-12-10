#
# STAGE: kernel
#
FROM scratch as kernel
COPY --from=linuxkit/kernel:5.15.15 kernel kernel

#
# STAGE: build-init
#
FROM golang:1.19 as build-init

WORKDIR /build

COPY Makefile go.mod go.sum .
RUN go mod download && go mod verify

COPY cmd cmd
RUN make build && cp build/init /init

#
# STAGE: build-initramfs
#
FROM alpine as build-initramfs

WORKDIR /initramfs

RUN mkdir proc sys
COPY --from=build-init /build/build/init init
RUN find . -print0 \
    | cpio -0 -o -H newc \
    | gzip -9 > /initramfs.gz

#
# STAGE: initramfs
#
FROM scratch AS initramfs
COPY --from=build-initramfs initramfs.gz initramfs.gz
