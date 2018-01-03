FROM scratch
#FROM alpine

ADD ./build/score /
# CMD ["mkdir ./schemas"]
ADD ./schemas /schemas

CMD ["/score"]

EXPOSE 8000