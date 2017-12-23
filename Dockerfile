FROM scratch

ADD ./build/score /
# CMD ["mkdir ./schemas"]
ADD ./schemas /schemas

CMD ["/score"]

EXPOSE 8080