#!/bin/bash -xe

IMAGE_NAME=go-sample-site
DOCKER_REGISTRY=ccr.ccs.tencentyun.com/koderover-rc
DOCKERFILE=build/Dockerfile

image_registry="${DOCKER_REGISTRY}/${IMAGE_NAME}"
image_tag_specific="${image_registry}:0.3.1"

echo "Build and push image."
docker build -t ${image_tag_specific} -f ${DOCKERFILE} .
docker push ${image_tag_specific}

set +x
echo ""
echo "Image ready:"
echo "   ${image_tag_specific}"
