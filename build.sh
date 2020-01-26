#!/bin/sh
operator-sdk build quay.io/tpokki/gatling-operator && \
	docker push quay.io/tpokki/gatling-operator
