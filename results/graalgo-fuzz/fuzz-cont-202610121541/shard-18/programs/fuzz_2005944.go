package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int64; F2 string }
func genFn0(p0 int, p1 uint32, p2 uint32) int8 {
    switch p0 {
    case -103:
        var v5 int16 = 1
        _ = v5
    case 42:
        p2 = p1
    default:
        v6 := GenS1{F0: 18446744073709551615, F1: 1, F2: 1}
        _ = v6
    }
    p1 = p1
    var v7 int8 = 100
    v8 := GenS2{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return v7
}
func genFn1(p0 int64) uint32 {
    v9 := []int{}
    if (!((!false) && false)) {
        if len(v9) > 0 {
            v9[0] = 805
        }
    }
    _ = p0
    _ = v9
    return 280
}

func main() {
    var v1 float32 = 0.1
    var v2 int8 = 127
    v3 := make(map[int]uint32)
    var v4 bool = (-32768 >= -520)
    v2 = (^(v2))
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS2{F0: 1}
        _ = fi10
        _ = v11
    }
    v1 -= float32(1e10)
    var v12 int = -568
    v13 := GenS1{F0: 619, F1: 2147483647, F2: 473}
    fmt.Println(v2)
    fmt.Println(v1)
    v14 := GenS3{F0: 156}
    var v15 int16 = 1000
    var v16 int32 = 1
    fmt.Println(v4)
    v15 = (v15 - v15)
    fmt.Println(v2)
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 string = ("" + "")
        if ((!(v4 || ((1 > 168) || v4))) || (v2 < -1)) {
            for fi19 := 0; fi19 < 20; fi19++ {
                v2 = v2
                _ = fi19
            }
        }
        if v4 {
            var v20 uint32 = 2
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    v16 = v16
    if (v2 == v2) {
        fmt.Println(v16)
        if (-9223372036854775808 == 42) {
            fmt.Println(v1)
        }
        v21 := GenS4{F0: 4294967295, F1: -1, F2: "a"}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
