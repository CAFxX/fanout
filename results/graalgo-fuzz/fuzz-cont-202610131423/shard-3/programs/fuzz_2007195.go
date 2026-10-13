package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
func genFn0(p0 int16, p1 uint64, p2 int32) uint8 {
    v4 := GenS1{F0: 1, F1: 42, F2: 2147483647}
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := []int{int(p0), 571}
        switch ("foo bar" + "hello") {
        case "":
            var v7 float32 = -2.25
            _ = v7
        case "foo bar":
            fmt.Println(fi5)
        case "hello":
            if true {
                _ = len(v6)
            } else {
                fmt.Println(p0)
            }
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    var v1 int8 = 100
    var v2 float32 = (((float32(float64(-2.25)) * (0.1 + 0.0)) / 1.5) / 0.1)
    var v3 int16 = -1
    v8 := []int{757}
    var v9 float32 = 1.5
    v10 := GenS1{F0: 933, F1: 288, F2: 32768}
    v2 += v9
    v1 -= (v1 % ((-71) | 1))
    switch ("hello" + "") {
    case "foo bar":
        v9 = float32(1e10)
    }
    switch -9223372036854775808 {
    case 127:
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS1{F0: 42, F1: 65535, F2: fi11}
            var v13 int32 = 1
            _ = fi11
            _ = v12
            _ = v13
        }
    case 0:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    default:
        v14 := GenS1{F0: 1, F1: 42, F2: -38}
        _ = v14
    }
    v9 = v9
    _ = len(v8)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v9 = float32(float64(v2))
    switch "hello" {
    case "":
        v1 = (v1 / ((v1) | 1))
    case "foo bar":
        v3 = (v3 * v3)
    default:
        v3 -= (v3 << ((42) & 63))
    }
    v15 := GenS2{F0: 743, F1: 298, F2: 1}
    switch 32767 {
    case 127:
        v3 -= (^(v3))
    default:
        var v16 uint32 = 578
        _ = v16
    }
    var v17 uint16 = 350
    _ = len(v8)
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
