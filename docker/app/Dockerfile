#FROM golang
#
#RUN mkdir -p /go/src/github.com/goforbroke1006/teamtrack
#
#ADD . /go/src/github.com/goforbroke1006/teamtrack
#
#RUN go get  -t -v ./...
#RUN go get  github.com/canthefason/go-watcher
#RUN go install github.com/canthefason/go-watcher/cmd/watcher
#
#ENTRYPOINT  watcher -run github.com/goforbroke1006/teamtrack/cmd  -watch github.com/goforbroke1006/teamtrack


FROM golang

ENV PROJECT_NAME=github.com/goforbroke1006/teamtrack

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get github.com/canthefason/go-watcher && \
    go install github.com/canthefason/go-watcher/cmd/watcher

RUN mkdir -p /go/src/${PROJECT_NAME}
ADD ./ /go/src/${PROJECT_NAME}
RUN cd /go/src/${PROJECT_NAME}/ && make dep gen

ENV APP_ARGS ''

ENTRYPOINT watcher -run "/go/src/${PROJECT_NAME}/cmd/" \
    -watch ${PROJECT_NAME} \
    ${APP_ARGS}
