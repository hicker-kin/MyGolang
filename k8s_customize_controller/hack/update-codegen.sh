#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# 注意:
# 1. kubebuilder2.3.2版本生成的api目录结构code-generator无法直接使用(将api由api/${VERSION}移动至api/${GROUP}/${VERSION}即可)

# 2. modify: MODULE、GROUP_VERSION、APIS_PKG is necessary

# corresponding to go mod init <module>
MODULE=k8s_customize_controller
# api package
APIS_PKG=apis
# generated output package
OUTPUT_PKG=pkg/client
# group-version, modify as you info: such as foo:v1alpha1
GROUP_VERSION=stable:v1


SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
echo $SCRIPT_ROOT
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}
echo $CODEGEN_PKG

rm -rf ${OUTPUT_PKG}/{clientset,informers,listers}

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
#bash "${CODEGEN_PKG}"/generate-groups.sh "client,informer,lister" \
bash "${CODEGEN_PKG}"/generate-groups.sh all \
  ${MODULE}/${OUTPUT_PKG} ${MODULE}/pkg/${APIS_PKG} \
  ${GROUP_VERSION} \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt
#  --output-base "${SCRIPT_ROOT}"
#  --output-base "${SCRIPT_ROOT}/../../.."
