package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint }
type GenS2 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int32, p1 int64, p2 uint) uint16 {
    v5 := GenS1{F0: 1, F1: -1, F2: (^(p2))}
    switch -9223372036854775808 {
    case 0:
        var v6 float64 = 1.5
        _ = v6
    case -9223372036854775808:
        v7 := GenS1{F0: 18446744073709551615, F1: (-(p1)), F2: (p2 - p2)}
        _ = v7
    default:
        v8 := []int{-1, 42, 32767, -355}
        _ = v8
    }
    v9 := GenS2{F0: 42, F1: 965}
    var v10 bool = false
    var v11 uint16 = 593
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v5
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    v12 := GenS1{F0: 42, F1: 9223372036854775807, F2: 42}
    var v13 uint = 314
    for fi14 := 0; fi14 < 14; fi14++ {
        switch 128 {
        case -929:
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 int16 = 803
                _ = fi15
                _ = v16
            }
        case 0:
            var v17 int32 = 42
            _ = v17
        }
        v13 = (^(v13))
        var v18 int = fi14
        _ = fi14
        _ = v18
    }
    fmt.Println(v13)
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 bool = ((fi19 % (((fi19 - -9223372036854775808)) | 1)) >= (fi19 % (((fi19 / ((fi19) | 1))) | 1)))
        v21 := GenS2{F0: 42, F1: 65535}
        _ = fi19
        _ = v20
        _ = v21
    }
    v13 -= 4294967295
    delete(v1, 2147483647)
    switch -2147483648 {
    case 127:
        for fi22 := 0; fi22 < 8; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 int = (-(fi23))
                _ = fi23
                _ = v24
            }
            v13 -= 0
            _ = fi22
        }
    case 9223372036854775807:
        var v25 float32 = 0.0
        _ = v25
    default:
        var v26 uint = v13
        _ = v26
    }
    switch 4294967295 {
    case -386:
        var v27 int64 = 0
        _ = v27
    case 1:
        var v28 string = "a"
        _ = v28
    case -658:
        v2[("a" + ("a" + "a"))] = "hello"
    default:
        v29 := []int{937, 302, -1}
        _ = v29
    }
    if true {
        var v30 int16 = -32768
        v31 := v2["a"]
        _ = v30
        _ = v31
    } else {
        v13 += v13
        fmt.Println(v13)
        var v32 uint8 = 3
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
