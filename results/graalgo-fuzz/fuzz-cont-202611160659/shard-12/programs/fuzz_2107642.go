package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int16 }

func main() {
    var v1 int8 = -85
    var v2 uint32 = 586
    var v3 int16 = 42
    var v4 uint64 = 286
    v5 := []int{32767, 127, 42, int(v4)}
    var v6 bool = false
    fmt.Println(v2, v4)
    v7 := GenS1{F0: 0}
    _ = len(v5)
    v8 := GenS1{F0: -1}
    for fi9 := 0; fi9 < 20; fi9++ {
        _ = len(v5)
        for fi10 := 0; fi10 < 18; fi10++ {
            v2 -= (v2 % ((v2) | 1))
            v1 = 0
            _ = fi10
        }
        _ = fi9
    }
    v11 := GenS2{F0: 492}
    switch "a" {
    case "":
        var v12 uint64 = 917
        _ = v12
    case "hello":
        fmt.Println(v2, v2, v6)
    default:
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint64 = (v4 + 387)
            _ = fi13
            _ = v14
        }
    }
    fmt.Println(v2)
    switch 127 {
    case 32768:
        for fi15 := 0; fi15 < 11; fi15++ {
            fmt.Println(v4, v4)
            var v16 uint16 = 88
            fmt.Println(v16, v1)
            _ = fi15
            _ = v16
        }
    case 32767:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        for fi17 := 0; fi17 < 8; fi17++ {
            v2 += (v2 >> ((65535) & 63))
            v2 = 283
            _ = fi17
        }
    }
    var v18 int32 = 195
    _ = len(v5)
    _ = len(v5)
    v19 := GenS3{F0: 88, F1: (v3 ^ (v3 - v3))}
    if len(v5) > 0 {
        v5[0] = 42
    }
    switch "a" {
    case "a":
        var v20 float32 = -2.25
        _ = v20
    default:
        var v21 uint = 2
        _ = v21
    }
    var v22 uint = 91
    var v23 int = -1
    v5 = append(v5, v23)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
