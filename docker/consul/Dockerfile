FROM consul:1.1.0

RUN mkdir -p /consul-initial

ADD ./docker/consul/entrypoint.sh   /root/
ADD ./.consul                       /consul-initial/

ENTRYPOINT sh /root/entrypoint.sh
