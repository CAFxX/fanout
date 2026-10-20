package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int8; F1 int; F2 int32 }
type GenS3 struct { F0 int16; F1 int16; F2 uint8 }
func genFn0(p0 int32) bool {
    var v6 uint16 = 113
    switch 127 {
    case 1:
        var v7 bool = false
        _ = v7
    case -341:
        for fi8 := 0; fi8 < 11; fi8++ {
            p0 = (p0 * p0)
            fmt.Println(v6)
            _ = fi8
        }
    default:
        var v9 int16 = 32767
        _ = v9
    }
    if (!true) {
        for fi10 := 0; fi10 < 10; fi10++ {
            v6 -= 255
            _ = fi10
        }
    } else {
        var v11 int64 = -18
        _ = v11
    }
    v12 := GenS1{F0: (0 != int(p0)), F1: 31}
    _ = p0
    _ = v12
    _ = v6
    return false
}
func genFn1(p0 uint16, p1 uint16) int {
    var v13 int32 = 2147483647
    v13 = int32(p0)
    v14 := GenS2{F0: -12, F1: 42, F2: (v13 | v13)}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := []uint32{}
    var v4 bool = false
    v5 := make(map[string]bool)
    var v15 int64 = 42
    switch "" {
    case "a":
        v16 := GenS2{F0: 100, F1: -679, F2: -150}
        _ = v16
    case "":
        var v17 int64 = (v15 + v15)
        _ = v17
    case "hello":
        v15 = v15
    default:
        v15 += v15
    }
    v18 := GenS1{F0: v4, F1: 210}
    var v19 int8 = 114
    v19 -= 42
    fmt.Println(v15)
    if v4 {
        switch 2147483647 {
        case -9223372036854775808:
            v20 := GenS3{F0: -32768, F1: 306, F2: 1}
            _ = v20
        default:
            v15 -= v15
        }
        var v21 int16 = -32768
        var v22 uint = 519
        _ = v21
        _ = v22
    } else {
        if false {
            v15 += (v15 + v15)
        }
        var v23 int64 = (v15 & v15)
        _ = v23
    }
    if (-272 != -280) {
        v4 = true
        var v24 int64 = (v15 | (v15 * v15))
        _ = v24
    } else {
        switch ((("foo bar" + "hello") + "a") + (("\n\t" + "\n\t") + "hello")) {
        case "a":
            v4 = v4
        }
        switch 32768 {
        case 64:
            delete(v2, 1)
        }
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        v4 = (float32(-2.25) == 1e10)
        _ = fi25
    }
    var v26 float64 = float64(float32(0.1))
    if true {
        v26 += float64(3.14159)
        var v27 int = 1
        _ = v27
    }
    fmt.Println(v4)
    v19 += v19
    switch 32767 {
    case 127:
        var v28 uint8 = 2
        _ = v28
    }
    if (2147483647 != 229) {
        v19 = v19
    }
    v19 += v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
