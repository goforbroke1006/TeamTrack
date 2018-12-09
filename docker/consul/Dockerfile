FROM consul:1.1.0

ADD ./docker/consul/entrypoint.sh   /root
ADD ./.consul                       /consul-initial

#VOLUME [ "/consul-initial/" ]

ENTRYPOINT sh /root/entrypoint.sh
