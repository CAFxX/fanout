package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "a"
    var v3 int32 = 1
    var v4 uint32 = 2
    var v5 string = v2
    delete(v1, v4)
    if (-2147483648 > (v3 % (((v3 >> ((480) & 63))) | 1))) {
        fmt.Println(v4, v3, v5)
        switch 1 {
        case 767:
            fmt.Println(v3, v3, v2)
        case 127:
            v3 += (v3 >> ((65535) & 63))
        default:
            v5 = (v2 + "a")
        }
        fmt.Println(v4, v2)
    }
    v6 := GenS1{F0: 0}
    var v7 string = v5
    if true {
        v4 += (v4 * (v4 + v4))
    }
    var v8 uint = 4294967295
    v9 := GenS2{F0: 255, F1: 498, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
