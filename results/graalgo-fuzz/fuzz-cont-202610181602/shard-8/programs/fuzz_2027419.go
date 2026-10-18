package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint16, p2 int64) uint32 {
    v4 := GenS1{F0: 628}
    v5 := GenS1{F0: 889}
    var v6 uint8 = uint8(p0)
    p0 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 234
}
func genFn1(p0 string, p1 int) int8 {
    v7 := []int{p1, 1, 181, (p1 / ((p1) | 1))}
    p0 = p0
    if len(v7) > 0 {
        v7[0] = p1
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(p1, fi8)
        p0 = ("hello" + p0)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return -100
}

func main() {
    v1 := []uint32{2147483647, 2147483647, 65535, 2}
    var v2 int32 = -333
    var v3 uint32 = 2147483647
    for fi9 := 0; fi9 < 10; fi9++ {
        v3 = (^(v3))
        v3 = 350
        _ = fi9
    }
    fmt.Println(v3, v2)
    switch -513 {
    case 32767:
        if len(v1) > 0 {
            v1[0] = (v3 % ((48) | 1))
        }
    default:
        v2 -= v2
    }
    fmt.Println(v2, v3)
    if false {
        for fi10 := 0; fi10 < 12; fi10++ {
            v3 -= v3
            _ = fi10
        }
        if (-128 >= 0) {
            fmt.Println(v2, v2)
            var v11 uint64 = uint64(v3)
            var v12 uint64 = 296
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int16 = 0
        switch "\n\t" {
        case "\n\t":
            _ = len(v1)
        case "a":
            if len(v1) > 0 {
                v1[0] = v3
            }
        default:
            var v15 uint32 = v3
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    var v16 int32 = v2
    v3 = v3
    var v17 int = 122
    if true {
        var v18 float32 = 0.1
        _ = len(v1)
        _ = v18
    } else {
        v1 = append(v1, 65535)
        v19 := GenS1{F0: 460}
        _ = v19
    }
    if false {
        fmt.Println(v17, v16)
    } else {
        fmt.Println(v3, v3, v16)
        var v20 uint32 = (v3 + v3)
        switch v17 {
        case 42:
            if false {
                v16 -= v2
                var v21 string = ("" + "hello")
                var v22 uint16 = 32767
                _ = v21
                _ = v22
            }
        case 128:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            v3 = v3
        }
        _ = v20
    }
    v17 = (v17 >> ((65535) & 63))
    v16 = v16
    var v23 uint8 = 0
    var v24 int8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
