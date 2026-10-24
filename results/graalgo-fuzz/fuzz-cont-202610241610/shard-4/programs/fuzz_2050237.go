package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }

func main() {
    var v1 uint64 = 901
    v2 := make(map[string]bool)
    var v3 int8 = 127
    v4 := []string{}
    switch 127 {
    case -1:
        v4 = append(v4, ("hello" + "foo bar"))
    case 32767:
        v5 := GenS1{F0: (v3 - 51), F1: -648}
        _ = v5
    default:
        v3 += (v3 << ((65535) & 63))
    }
    var v6 uint8 = 184
    switch 32767 {
    case 2147483647:
        var v7 int32 = -295
        _ = v7
    default:
        for fi8 := 0; fi8 < 17; fi8++ {
            v6 = v6
            var v9 float64 = 1.5
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        _ = len(v4)
        if true {
            fmt.Println(v6, fi10)
        } else {
            v1 = (^(v1))
            var v11 uint32 = 2147483647
            _ = v11
        }
        _ = fi10
    }
    if false {
        v2[""] = false
    }
    fmt.Println(v3)
    var v12 uint = 65535
    var v13 int = -9223372036854775808
    v2["hello"] = true
    var v14 uint16 = 0
    var v15 uint16 = v14
    var v16 string = ("" + "")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
