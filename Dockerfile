FROM progrium/busybox
RUN mkdir -p /etc/ssl && mkdir -p /etc/ssl/certs
ADD certs /etc/ssl/certs/
ADD deployAgent/deployAgent deployAgent
ENTRYPOINT ["./deployAgent"]