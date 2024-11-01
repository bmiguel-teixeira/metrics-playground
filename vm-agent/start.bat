docker run --rm --network host -v .:/config victoriametrics/vmagent -streamAggr.config=/config/config.yaml -usePromCompatibleNaming -remoteWrite.url=http://localhost:9090/api/v1/write
