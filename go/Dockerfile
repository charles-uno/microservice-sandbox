
FROM centos

RUN yum install -y epel-release

RUN yum install -y \
    go \
    make

RUN mkdir -p /sandbox
ADD . /sandbox
WORKDIR /sandbox

RUN make app

CMD ./app
