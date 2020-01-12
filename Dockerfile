FROM ubuntu:latest

RUN apt-get update
RUN apt-get -y install git gnupg curl build-essential libssl-dev libreadline-dev zlib1g-dev libsqlite3-dev tzdata

COPY rails-init rails-init

CMD ./rails-init
