package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int; F1 int32; F2 uint64 }
func genFn0(p0 int8, p1 uint, p2 uint8) uint {
    if false {
        if ((360 > 4294967295) && false) {
            v4 := GenS1{F0: 32767, F1: 397}
            _ = v4
        }
        p0 += -128
        fmt.Println(p2, p1)
    }
    var v5 uint8 = (^(p2))
    fmt.Println(p2)
    v5 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}
func genFn1() int64 {
    var v6 uint64 = 972
    v6 = v6
    _ = v6
    return 0
}

func main() {
    v1 := []uint32{2147483647, 0}
    var v2 bool = (true && (-128 > 8))
    var v3 uint8 = 255
    fmt.Println(v2, v3)
    v7 := GenS1{F0: -1, F1: 2147483647}
    v3 += (v3 | v3)
    if v2 {
        for fi8 := 0; fi8 < 3; fi8++ {
            v2 = false
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch fi8 {
            case -1:
                _ = len(v1)
            case 32767:
                var v9 bool = ((!true) && true)
                _ = v9
            default:
                v3 = (v3 * (v3 | (v3 >> ((65535) & 63))))
            }
            _ = fi8
        }
        v10 := GenS1{F0: 1000, F1: 2147483647}
        _ = v10
    }
    var v11 uint16 = 451
    var v12 uint32 = 42
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint8 = 127
            var v16 float32 = 3.14159
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 uint8 = (v3 ^ (v3 | v3))
        v2 = v2
        _ = fi13
        _ = v17
    }
    fmt.Println(v11, v11)
    v18 := GenS2{F0: 122, F1: 255}
    fmt.Println(v11, v12)
    v19 := GenS1{F0: 32767, F1: v12}
    _ = len(v1)
    switch "foo bar" {
    case "foo bar":
        var v20 int64 = -1
        _ = v20
    case "\n\t":
        if (v12 != (v12 & (v12 | v12))) {
            var v21 uint64 = 1
            _ = v21
        }
    case "a":
        for fi22 := 0; fi22 < 15; fi22++ {
            fmt.Println(fi22, v2, fi22)
            switch "" {
            case "\n\t":
                if (v2 && v2) {
                    v23 := GenS2{F0: 42, F1: uint(v12)}
                    _ = v23
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            if ((v3 ^ (^(v3))) != (^(v3))) {
                fmt.Println(v2, fi22, v2)
            }
            _ = fi22
        }
    default:
        v24 := GenS3{F0: -9223372036854775808, F1: -1, F2: 389}
        _ = v24
    }
    v2 = true
    v25 := GenS1{F0: -1, F1: v12}
    var v26 bool = true
    var v27 int64 = -963
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v7
}
