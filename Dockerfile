
FROM centos

RUN yum install -y epel-release

RUN yum -y update

RUN yum install -y python python-pip

RUN pip install flask connexion

RUN mkdir -p /sandbox

ADD . /sandbox

WORKDIR /sandbox

EXPOSE 5000
