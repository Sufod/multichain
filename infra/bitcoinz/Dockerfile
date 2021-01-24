FROM ubuntu:xenial

WORKDIR /root

RUN apt-get update --fix-missing && apt-get install --yes software-properties-common wget libgomp1

RUN wget -c https://github.com/btcz/bitcoinz/releases/download/2.0.7/bitcoinz-2.0.7-8e85a88df-amd64.deb 
RUN dpkg -i bitcoinz-2.0.7-8e85a88df-amd64.deb
RUN mkdir -p /root/.zcash-params && zcash-fetch-params

COPY bitcoinz.conf /root/bitcoinz.conf
COPY run.sh /root/run.sh
RUN chmod +x /root/run.sh

EXPOSE 11337

ENTRYPOINT ["run.sh"]
