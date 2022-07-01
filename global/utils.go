package global

import (
    "encoding/json"
    "net"
)

// ServerIP 获取当前 IP
func ServerIP() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        return ""
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP.String()
}

// DepCopy 复制source的字段到target target需要指针类型
func DepCopy(source interface{}, target interface{}) error {
    if b, err := json.Marshal(source); err != nil {
        return err
    } else {
        if err := json.Unmarshal(b, target); err != nil {
            return err
        }
    }
    return nil
}
