FROM ubuntu:xenial

RUN apt-get update --fix-missing && apt-get install --yes software-properties-common wget

RUN wget -c https://github.com/btcz/bitcoinz/releases/download/2.0.7/bitcoinz-2.0.7-8e85a88df-amd64.deb -O - 
RUN dpkg -i bitcoinz-2.0.7-8e85a88df-amd64.deb

# TODO FINISH THIS + add /run.sh

COPY run.sh /root/run.sh
RUN chmod +x /root/run.sh

EXPOSE 11337

ENTRYPOINT ["./root/run.sh"]
