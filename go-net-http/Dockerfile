
FROM centos

RUN yum install -y epel-release

RUN yum -y update

RUN yum install -y go

RUN yum install -y make

RUN mkdir -p /sandbox

ADD . /sandbox

WORKDIR /sandbox

# TODO -- Would be better to already have built it and install via RPM
RUN make app
