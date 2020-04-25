#!/bin/bash

IFACES="eno1"
MTU=9000
TXQ=500
RXU=300
MBUF=204217728

# Don't use MPTCP on the management interface
#ip link set dev eth0 multipath off

# Setup RFS
for CPU in 0 1 2 3 4 5 6 7
do
	for iface in $IFACES; do
		echo ff > /sys/class/net/$iface/queues/rx-${CPU}/rps_cpus
		echo ff > /sys/class/net/$iface/queues/tx-${CPU}/xps_cpus
		echo 1024 > /sys/class/net/$iface/queues/rx-${CPU}/rps_flow_cnt
	done
done

echo 1024 > /proc/sys/net/core/rps_sock_flow_entries

service irqbalance stop

# Let's use jumbo frames
for iface in $IFACES; do
	ifconfig $iface mtu $MTU txqueuelen $TXQ
done

# Interrupt coalescing
for iface in $IFACES; do
	ethtool -C $iface rx-usecs $RXU
done

# IRQ CPU Affinity
for i in {104..157}; do
        echo "04" > /proc/irq/$i/smp_affinity
done


# Configure sysctl's

sysctl -w net.mptcp.mptcp_checksum=0

sysctl -w net.ipv4.tcp_rmem="4096 524288 $MBUF"
sysctl -w net.ipv4.tcp_wmem="4096 524288 $MBUF"
sysctl -w net.ipv4.tcp_mem="768174 10242330 15363480"

sysctl -w net.ipv4.tcp_low_latency=1
sysctl -w net.ipv4.tcp_no_metrics_save=1
sysctl -w net.ipv4.tcp_timestamps=1
sysctl -w net.ipv4.tcp_sack=1

sysctl -w net.core.rmem_max=524287
sysctl -w net.core.wmem_max=524287
sysctl -w net.core.optmem_max=524287

sysctl -w net.core.netdev_max_backlog=10000

sysctl -w net.ipv4.tcp_congestion_control=cubic
