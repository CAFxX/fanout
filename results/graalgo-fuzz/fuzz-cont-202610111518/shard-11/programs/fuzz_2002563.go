package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint64 }

func main() {
    var v1 int16 = -230
    var v2 uint8 = 110
    var v3 int = 1
    v4 := []uint32{}
    fmt.Println(v2)
    v2 += 233
    switch "hello" {
    case "a":
        fmt.Println(v2)
    default:
        _ = len(v4)
    }
    _ = len(v4)
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: 849, F1: -9223372036854775808, F2: 2}
        if len(v4) > 0 {
            v4[0] = 0
        }
        _ = fi5
        _ = v6
    }
    if (true && (790 > 4294967295)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v7 int8 = 11
        _ = v7
    }
    if len(v4) > 0 {
        v4[0] = 65535
    }
    v3 = (^(v3))
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(v1)
        _ = fi8
    }
    switch (-(v3)) {
    case 1:
        v2 -= 0
    default:
        for fi9 := 0; fi9 < 18; fi9++ {
            fmt.Println(v3, v3, v2)
            for fi10 := 0; fi10 < 15; fi10++ {
                v1 -= (v1 / (((v1 / ((v1) | 1))) | 1))
                switch "hello" {
                case "":
                    v1 -= v1
                case "foo bar":
                    fmt.Println(fi9)
                case "hello":
                    v11 := GenS2{F0: 42, F1: 128, F2: 189}
                    _ = v11
                default:
                    var v12 uint16 = 42
                    v1 += -32768
                    var v13 uint = 583
                    _ = v12
                    _ = v13
                }
                _ = fi10
            }
            _ = fi9
        }
    }
    v2 = v2
    v3 = v3
    if ((v1 % (((^(v1))) | 1)) == (v1 / (((v1 ^ v1)) | 1))) {
        var v14 uint16 = 0
        _ = v14
    } else {
        v1 = v1
        var v15 int64 = 164
        _ = v15
    }
    _ = len(v4)
    v1 = (-(v1))
    var v16 uint = 0
    var v17 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
