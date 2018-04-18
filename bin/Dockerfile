FROM ubuntu
MAINTAINER Sergey Sergeev
RUN mkdir /home/test
COPY start.sh /home/test/
COPY go_build_main_go /usr/bin/go_build_main.go
ENTRYPOINT /home/test/start.sh