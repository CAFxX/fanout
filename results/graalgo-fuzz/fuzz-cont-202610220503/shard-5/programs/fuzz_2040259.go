package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint; F2 int32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint8) uint16 {
    var v7 bool = false
    var v8 float32 = (1e10 + (1.5 + 0.0))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1() int8 {
    var v9 uint8 = 127
    var v10 int32 = 42
    fmt.Println(v10, v10)
    fmt.Println(v10, v10)
    _ = v10
    _ = v9
    return 100
}

func main() {
    var v1 float32 = -2.25
    var v2 uint32 = 923
    v3 := make(map[int]string)
    var v4 int64 = -9223372036854775808
    v5 := make(map[uint32]bool)
    v6 := make(map[string]uint32)
    v2 += (v2 ^ v2)
    switch -96 {
    case 128:
        v1 -= float32(0.1)
    default:
        v11 := GenS1{F0: true, F1: 115, F2: v2}
        _ = v11
    }
    v2 -= v2
    for fi12 := 0; fi12 < 17; fi12++ {
        fmt.Println(v2)
        _ = fi12
    }
    var v13 bool = (436 <= 938)
    if true {
        v14 := []int{128}
        v1 = float32((1e10 - (float64(((v1 * float32(1.5)) - v1)) - 3.14159)))
        for fi15 := 0; fi15 < 17; fi15++ {
            if (262 < 42) {
                v13 = (!v13)
                var v16 uint = 906
                v17 := GenS2{F0: 2, F1: (v16 | (v16 | v16)), F2: 42}
                _ = v16
                _ = v17
            } else {
                v18 := GenS1{F0: (13 <= -42), F1: 2, F2: 2147483647}
                var v19 uint = 1
                _ = v18
                _ = v19
            }
            _ = fi15
        }
        _ = v14
    }
    v20 := GenS1{F0: true, F1: 255, F2: 2}
    v13 = (!(!true))
    var v21 uint8 = 2
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 uint16 = 513
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 float32 = (v1 - 0.0)
            v26 := GenS2{F0: 4294967295, F1: 28, F2: 0}
            _ = len(v3)
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    var v27 int64 = v4
    var v28 uint16 = 16
    v4 = v4
    if true {
        v13 = (v1 >= v1)
        v4 -= v4
        for fi29 := 0; fi29 < 5; fi29++ {
            delete(v3, 32767)
            v13 = ((^(v28)) < v28)
            _ = fi29
        }
    } else {
        v30 := GenS3{F0: (v21 & v21)}
        v4 += (-(v4))
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
