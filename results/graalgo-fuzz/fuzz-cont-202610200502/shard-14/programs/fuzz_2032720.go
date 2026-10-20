package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 bool) int64 {
    var v6 int8 = 108
    v7 := GenS1{F0: 0, F1: v6, F2: 1}
    v8 := GenS1{F0: 32768, F1: v6, F2: 1}
    if p0 {
        var v9 float64 = 3.14159
        _ = v9
    } else {
        v10 := []int{32767, 9223372036854775807, 127}
        p0 = (p0 && p0)
        var v11 uint8 = 2
        _ = v10
        _ = v11
    }
    v6 = v6
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -97
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{}
    v3 := make(map[int]string)
    var v4 uint64 = 395
    v5 := make(map[uint32]bool)
    var v12 int8 = 1
    v13 := GenS2{F0: 42, F1: 287}
    v12 -= (v12 / (((v12 / ((127) | 1))) | 1))
    for fi14 := 0; fi14 < 10; fi14++ {
        switch fi14 {
        case -9223372036854775808:
            v4 = 42
        case 9223372036854775807:
            var v15 int = fi14
            _ = v15
        case 32768:
            v4 -= 188
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 uint = 255
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 uint64 = 0
            _ = fi18
            _ = v19
        }
        _ = fi14
    }
    v20 := GenS3{F0: 2}
    var v21 float32 = (1e10 / 1e10)
    _ = len(v3)
    var v22 int8 = (v12 << ((42) & 63))
    fmt.Println(v21, v22, v12)
    for fi23 := 0; fi23 < 7; fi23++ {
        for fi24 := 0; fi24 < 13; fi24++ {
            v4 -= (^(v4))
            if (true || false) {
                v12 -= v12
                v12 -= (v12 & v12)
                v21 = 0.0
                var v25 string = "hello"
                delete(v1, fi24)
                _ = v25
            }
            v26 := GenS4{F0: 1}
            _ = fi24
            _ = v26
        }
        v27 := GenS5{F0: 602, F1: -32768}
        _ = fi23
        _ = v27
    }
    if false {
        v21 = ((v21 * -2.25) + v21)
    }
    fmt.Println(v21, v22, v4)
    for fi28 := 0; fi28 < 13; fi28++ {
        for fi29 := 0; fi29 < 19; fi29++ {
            v30 := GenS6{F0: 32767}
            fmt.Println(fi29, fi29, v21)
            v21 = float32(0.1)
            _ = fi29
            _ = v30
        }
        v31 := GenS2{F0: -2147483648, F1: 65535}
        _ = fi28
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
