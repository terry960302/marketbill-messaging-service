FROM public.ecr.aws/lambda/provided:al2 as build
LABEL Terry Kim <terry960302@gmail.com>

RUN yum install -y golang
RUN go env -w GOPROXY=direct

ADD go.mod go.sum ./
RUN go mod download

ADD . .
RUN env GOOS=linux GOARCH=amd64 go build -o /main

FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /main /main

ENV PROFILE=prod
ENV SENS_HOST=https://sens.apigw.ntruss.com
ENV SENS_SERVICE_ID=ncp:sms:kr:290881020329:marketbill-project
ENV SENS_ACCESS_KEY_ID=2aJkrtHdUtk5NP4oG8yh
ENV SENS_SECRET_KEY=x2A3OXOz0P1qmaLDnTiqo2dQ7if6BzOElQEPNg6b

ENTRYPOINT [ "/main" ]