package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int64 }
func genFn0(p0 uint8) uint8 {
    fmt.Println(p0)
    var v4 string = ((("a" + ("a" + "")) + "foo bar") + "a")
    if ((false && ((2147483647 < 42) && false)) || true) {
        v5 := []int{1, -1, 669}
        _ = v5
    }
    _ = p0
    _ = v4
    return 127
}
func genFn1(p0 uint16) bool {
    var v6 int32 = -122
    v7 := []int{-54, -199}
    fmt.Println(p0, p0)
    v6 = int32(p0)
    _ = p0
    _ = v6
    _ = v7
    return true
}

func main() {
    var v1 int = 127
    var v2 uint64 = 1
    v3 := make(map[uint32]string)
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            if true {
                var v10 string = "\n\t"
                v11 := GenS1{F0: 889, F1: uint(v2), F2: 9223372036854775807}
                _ = v10
                _ = v11
            }
            var v12 uint8 = 233
            _ = fi9
            _ = v12
        }
        fmt.Println(v2, fi8)
        for fi13 := 0; fi13 < 2; fi13++ {
            fmt.Println(v2, v2)
            _ = fi13
        }
        _ = fi8
    }
    fmt.Println(v1, v1)
    if (747 > -9223372036854775808) {
        var v14 int32 = 979
        if (107 <= -272) {
            var v15 uint32 = 4294967295
            var v16 int32 = v14
            _ = v15
            _ = v16
        } else {
            var v17 float32 = (1e10 - ((3.14159 * (float32(3.14159) / 1.5)) - 0.0))
            switch "\n\t" {
            case "\n\t":
                v18 := []int{635}
                _ = v18
            case "a":
                v2 = (v2 | v2)
            case "hello":
                v2 = (v2 * v2)
            default:
                var v19 uint8 = 255
                _ = v19
            }
            _ = v17
        }
        var v20 int8 = 0
        _ = v14
        _ = v20
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        v2 -= v2
        v2 -= (v2 ^ (v2 / ((1) | 1)))
        _ = fi21
    }
    var v22 int64 = 1
    if ((v1 << ((469) & 63)) >= (-(v1))) {
        fmt.Println(v2, v1, v22)
        v23 := GenS1{F0: 65535, F1: 2, F2: (-(v22))}
        _ = v23
    }
    var v24 uint64 = 0
    var v25 uint8 = 0
    _ = len(v3)
    switch v1 {
    case -2147483648:
        var v26 uint64 = 4294967295
        _ = v26
    case 32768:
        v22 -= (v22 ^ (v22 / ((v22) | 1)))
    }
    delete(v3, 295)
    fmt.Println(v1, v22)
    for fi27 := 0; fi27 < 8; fi27++ {
        v25 = 0
        for fi28 := 0; fi28 < 12; fi28++ {
            switch v1 {
            case 570:
                var v29 uint64 = 545
                _ = v29
            default:
                var v30 uint = 390
                _ = v30
            }
            _ = fi28
        }
        _ = fi27
    }
    v1 = (v1 - int(v24))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
}
