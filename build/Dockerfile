FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/simple-clusteroperator-operator \
    USER_UID=1001 \
    USER_NAME=simple-clusteroperator-operator

# install operator binary
COPY build/_output/bin/simple-clusteroperator-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
