docker:
	docker build -t ping:1 -f ping/Dockerfile .
	docker build -t pong:1 -f pong/Dockerfile .

apply:
	kubectl apply -f manifests

apply-extra:
	kubectl apply -f manifests/pong-service.yaml
	kubectl apply -f extra-manifests/configmap.yaml
	kubectl apply -f extra-manifests/ping-pod.yaml
	kubectl apply -f extra-manifests/pong-pod.yaml

cleanup:
	kubectl delete pods --all
	kubectl delete service pong --ignore-not-found
	kubectl delete configmaps --all

redeploy:
	$(MAKE) cleanup
	$(MAKE) apply

redeploy-extra:
	$(MAKE) cleanup
	$(MAKE) apply-extra
