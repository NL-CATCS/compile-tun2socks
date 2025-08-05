package mobile

import (
    "github.com/xjasonlyu/tun2socks/v2/core"
    "github.com/xjasonlyu/tun2socks/v2/component/socks"
)

var app *core.App

func StartTun2Socks(fd int, ip, mask, mtu, proxyIp string, proxyPort int) {
    app, _ = core.New(core.Config{
        TUN: core.TUNConfig{
            FileDescriptor: fd,
            IP:             ip,
            Netmask:        mask,
            MTU:            mtu,
        },
        Handler: socks.Config{
            Address: proxyIp,
            Port:    uint16(proxyPort),
            UDP:     true,
        },
    })
    go app.Run()
}

func StopTun2Socks() {
    if app != nil {
        app.Close()
    }
}
