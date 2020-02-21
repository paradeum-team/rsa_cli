FROM centos:7
#FROM alpine:3.9
RUN  yum install -y openssl lzma && yum clean all

COPY ./bin/bfs_cli_rsa /data/
COPY afs-x64 /data/

#RUN useradd  -u 1001 ks && chown -R ks:ks /data
#USER ks
RUN chmod 755 /data/afs-x64

EXPOSE 5151
WORKDIR /data

CMD ["./bfs_cli_rsa"]
