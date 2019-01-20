
FROM centos

RUN yum install -y epel-release

RUN yum install -y \
    python \
    python-pip

RUN pip install \
    connexion \
    flask

RUN mkdir -p /sandbox
ADD . /sandbox
WORKDIR /sandbox

CMD ./app.py
