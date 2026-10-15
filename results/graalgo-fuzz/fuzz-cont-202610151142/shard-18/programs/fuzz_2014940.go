package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0(p0 int16, p1 int) int64 {
    switch p1 {
    case 32767:
        fmt.Println(p0)
    case 930:
        p1 += (p1 | p1)
    case 41:
        v4 := GenS1{F0: 42}
        _ = v4
    default:
        var v5 int32 = 2147483647
        _ = v5
    }
    if (92 != 0) {
        fmt.Println(p0, p1)
    }
    p1 = (p1 / ((9223372036854775807) | 1))
    var v6 int8 = 100
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 uint32 = 492
    v2 := []uint32{1, v1}
    var v3 uint16 = 0
    fmt.Println(v3, v1)
    if false {
        if false {
            var v7 uint32 = 2147483648
            _ = v7
        } else {
            v8 := GenS2{F0: 255, F1: 487}
            var v9 bool = false
            v3 = (v3 >> ((255) & 63))
            _ = v8
            _ = v9
        }
    }
    var v10 float32 = (3.14159 + (-2.25 * 1.5))
    v11 := GenS1{F0: 127}
    if ((((v10 + -2.25) * v10) - (float32(0.0) / v10)) < 1e10) {
        switch -9223372036854775808 {
        case -9223372036854775808:
            var v12 float32 = v10
            _ = v12
        case 1:
            v1 = v1
        }
    } else {
        var v13 bool = (0 != 2)
        v1 = 42
        _ = v13
    }
    if true {
        if len(v2) > 0 {
            v2[0] = uint32(v3)
        }
        v10 = (1e10 - (1.5 - (v10 - v10)))
        if len(v2) > 0 {
            v2[0] = v1
        }
    } else {
        var v14 int = 648
        var v15 uint64 = 2
        _ = v14
        _ = v15
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        var v16 int32 = 1
        if false {
            fmt.Println(v1)
            v16 += 0
        }
        v1 = (v1 % (((^(v1))) | 1))
        _ = v16
    } else {
        switch (("\n\t" + "a") + "") {
        case "":
            v10 = -2.25
        case "a":
            v17 := GenS2{F0: 42, F1: 42}
            _ = v17
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v10 += v10
    }
    v10 -= float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
