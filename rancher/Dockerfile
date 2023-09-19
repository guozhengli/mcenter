FROM bh-harbor.suixingpay.com/zhaikun/ops-go:v1
MAINTAINER jaxzhai
RUN yum -y install kde-l10n-Chinese \
    && localedef -c -f UTF-8 -i zh_CN zh_CN.UFT-8 \
    && echo 'LANG="zh_CN.UFT-8"' > /etc/locale.conf  && source /etc/locale.conf \
    && yum clean all \
    && mkdir /opt/mcenter \
    && mkdir /opt/mcenter/etc
ENV LANG=zh_CN.UFT-8 \
    LC_ALL=zh_CN.UFT-8
WORKDIR /opt/mcenter
COPY ./mcenter /opt/mcenter/
EXPOSE 8080
CMD ["./mcenter","start","-f","/opt/mcenter/etc/config.toml"]
