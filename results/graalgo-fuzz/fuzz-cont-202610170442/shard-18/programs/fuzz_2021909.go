package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int16; F2 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 int32; F1 int }
func genFn0(p0 bool, p1 uint) int8 {
    var v5 int32 = 42
    p0 = false
    fmt.Println(p0, v5, v5)
    v6 := GenS1{F0: p1, F1: 0}
    var v7 int8 = 42
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (v7 / ((-1) | 1))
}
func genFn1(p0 int64) int {
    v8 := GenS2{F0: -209}
    p0 += -635
    v9 := GenS3{F0: (p0 - p0), F1: 1, F2: 4294967295}
    if (!false) {
        v10 := GenS4{F0: (-2147483648 != -1)}
        _ = v10
    }
    _ = p0
    _ = v8
    _ = v9
    return -747
}

func main() {
    var v1 int16 = -1
    v2 := []string{"", ""}
    var v3 uint = 42
    v4 := []int{-311, 32768}
    v11 := GenS5{F0: 24, F1: 261}
    var v12 uint = (v3 + v3)
    var v13 int = -1
    if len(v4) > 0 {
        v4[0] = genFn1(42)
    }
    v3 -= v3
    var v14 int32 = -483
    v15 := GenS6{F0: 0, F1: v14}
    v3 = (v12 >> ((v3) & 63))
    for fi16 := 0; fi16 < 1; fi16++ {
        v12 -= 65535
        _ = fi16
    }
    if (v14 != v14) {
        fmt.Println(v12)
        v17 := GenS6{F0: 255, F1: (v14 << (((v3 & v12)) & 63))}
        _ = v17
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch (v13 ^ v13) {
        case 128:
            var v18 uint32 = 2147483647
            _ = v18
        case -9223372036854775808:
            v19 := GenS7{F0: 2, F1: 0}
            _ = v19
        default:
            var v20 int16 = v1
            _ = v20
        }
        switch v13 {
        case 32767:
            if (!((^(v1)) > v1)) {
                fmt.Println(v3, v14, v12)
                fmt.Println(v1, v3)
                v21 := GenS1{F0: v12, F1: v12}
                _ = v21
            }
        case 1:
            var v22 bool = false
            _ = v22
        default:
            for fi23 := 0; fi23 < 2; fi23++ {
                v14 = 0
                _ = fi23
            }
        }
    }
    if false {
        fmt.Println(v1)
    }
    if false {
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 int16 = int16(v12)
            _ = fi24
            _ = v25
        }
    } else {
        v26 := GenS8{F0: v14, F1: (v13 * 486)}
        switch "\n\t" {
        case "":
            fmt.Println(v12, v1)
        case "hello":
            v27 := GenS7{F0: 4294967295, F1: 971}
            _ = v27
        default:
            _ = len(v4)
        }
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
