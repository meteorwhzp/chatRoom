#!/bin/bash
workspace=$(cd $(dirname $0) && pwd -P)
echo "workspace: $workspace"
cd `dirname $0` || exit
#ulimit -c unlimited
mkdir -p ./status/city-info


nodemgr_conf=conf/nodemgr.json
config_conf=conf/config.ini

change_conf_by_cluster(){
    local clusterfile="$workspace/.deploy/service.cluster.txt"
    if [[ -f "$clusterfile" ]]; then
        local cluster=`cat $clusterfile`
        local nodemgr_conf_c="$nodemgr_conf.$cluster"
        if [[ -f $nodemgr_conf_c ]]; then
            cp $nodemgr_conf_c $nodemgr_conf
            echo "nodemgr_conf: using $nodemgr_conf_c"
        fi

        local config_conf_c="$config_conf.$cluster"
        if [[ -f $config_conf_c ]]; then
            cp $config_conf_c $config_conf
            echo "config_conf: using $config_conf_c"
        fi

    fi

}

start(){
    stop
    sleep 1
	#change_conf_by_cluster
    setsid ./supervise.city-info -u status/city-info ./city-info
}

stop(){
    killall -9 supervise.city-info
    killall -9 city-info
}

restart(){
    start
}


case C"$1" in
    Cstart)
        start
        echo "start Done!"
        ;;
    Cstop)
        stop
        echo "stop Done!"
        ;;
    Crestart)
        restart
        echo "restart Done!"
        ;;
    C*)
        echo "Usage: $0 {start|stop|restart}"
        ;;
esac
