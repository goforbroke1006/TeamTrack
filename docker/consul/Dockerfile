FROM consul:1.1.0

ADD ./entrypoint.sh /root

RUN ls -lah /
RUN ls -lah /root

ENTRYPOINT sh /root/entrypoint.sh
#ENTRYPOINT tail -f /dev/null