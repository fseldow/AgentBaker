#!/bin/bash

set -uxo pipefail

DEFAULT_CLIENT_VERSION="v0.1.0-alpha.2"

EVENTS_LOGGING_DIR=/var/log/azure/Microsoft.Azure.Extensions.CustomScript/events/

RETRY_PERIOD_SECONDS=180 # 3 minutes
RETRY_WAIT_SECONDS=5

CLIENT_BINARY_DOWNLOAD_URL="${CLIENT_BINARY_DOWNLOAD_URL:-https://k8sreleases.blob.core.windows.net/aks-tls-bootstrap-client/${DEFAULT_CLIENT_VERSION}/linux/amd64/tls-bootstrap-client}"
CLIENT_BINARY_PATH="${CLIENT_BINARY_PATH:-/opt/azure/tlsbootstrap/tls-bootstrap-client}"
KUBECONFIG_PATH="${KUBECONFIG_PATH:-/var/lib/kubelet/kubeconfig}"
API_SERVER_NAME="${API_SERVER_NAME:-""}"
AZURE_CONFIG_PATH="${AZURE_CONFIG_PATH:-/etc/kubernetes/azure.json}"
CLUSTER_CA_FILE_PATH="${CLUSTER_CA_FILE_PATH:-/etc/kubernetes/certs/ca.crt}"
AAD_RESOURCE="${AAD_RESOURCE:-""}"

retrycmd_if_failure() {
    retries=$1; wait_sleep=$2; timeout=$3; shift && shift && shift
    for i in $(seq 1 $retries); do
        timeout $timeout "${@}" && break || \
        if [ $i -eq $retries ]; then
            echo Executed \"$@\" $i times;
            return 1
        else
            sleep $wait_sleep
        fi
    done
    echo Executed \"$@\" $i times;
}

logs_to_events() {
    local task=$1; shift
    local eventsFileName=$(date +%s%3N)

    local startTime=$(date +"%F %T.%3N")
    ${@}
    ret=$?
    local endTime=$(date +"%F %T.%3N")

    json_string=$( jq -n \
        --arg Timestamp   "${startTime}" \
        --arg OperationId "${endTime}" \
        --arg Version     "1.23" \
        --arg TaskName    "${task}" \
        --arg EventLevel  "Informational" \
        --arg Message     "Completed: $*" \
        --arg EventPid    "0" \
        --arg EventTid    "0" \
        '{Timestamp: $Timestamp, OperationId: $OperationId, Version: $Version, TaskName: $TaskName, EventLevel: $EventLevel, Message: $Message, EventPid: $EventPid, EventTid: $EventTid}'
    )
    echo ${json_string} > ${EVENTS_LOGGING_DIR}${eventsFileName}.json

    if [ "$ret" != "0" ]; then
      return $ret
    fi
}

downloadClient() {
    [ -f "$CLIENT_BINARY_PATH" ] && exit 0
    DOWNLOAD_DIR=$(dirname $CLIENT_BINARY_PATH)

    if ! retrycmd_if_failure 30 5 60 curl -fSL -o "$CLIENT_BINARY_PATH" "$CLIENT_BINARY_DOWNLOAD_URL"; then
        echo "ERROR: unable to download secure TLS bootstrapping client binary from $CLIENT_BINARY_DOWNLOAD_URL"
        exit 1
    fi
    chown -R root:root "$DOWNLOAD_DIR"
    chmod -R 755 "$DOWNLOAD_DIR"
}

bootstrap() {
    if [ -z "$API_SERVER_NAME" ]; then
        echo "ERROR: missing apiserver FQDN, cannot continue bootstrapping"
        exit 1
    fi
    if [ ! -f "$CLIENT_BINARY_PATH" ]; then
        echo "ERROR: bootstrap client binary does not exist at path $CLIENT_BINARY_PATH"
        exit 1
    fi

    chmod +x $CLIENT_BINARY_PATH

    deadline=$(($(date +%s) + RETRY_PERIOD_SECONDS))
    while true; do
        now=$(date +%s)
        if [ $((now - deadline)) -ge 0 ]; then
            echo "ERROR: bootstrapping deadline exceeded"
            exit 1
        fi

        $CLIENT_BINARY_PATH bootstrap \
         --aad-resource="$AAD_RESOURCE" \
         --apiserver-fqdn="${API_SERVER_NAME}:443" \
         --cluster-ca-file="$CLUSTER_CA_FILE_PATH" \
         --azure-config="$AZURE_CONFIG_PATH" \
         --next-proto="aks-tls-bootstrap" \
         --kubeconfig="$KUBECONFIG_PATH"

        [ $? -eq 0 ] && exit 0

        sleep $RETRY_WAIT_SECONDS
    done
}

SUB_COMMAND=$1
if [ "${SUB_COMMAND,,}" == "download" ]; then
    logs_to_events "AKS.downloadSecureTLSBootstrapClient" downloadClient
elif [ "${SUB_COMMAND,,}" == "bootstrap" ]; then
    logs_to_events "AKS.performSecureTLSBootstrapping" bootstrap
else
    echo "ERROR: unknown subcommand $SUB_COMMAND for secure-tls-bootstrap.sh"
    exit 1
fi

#EOF