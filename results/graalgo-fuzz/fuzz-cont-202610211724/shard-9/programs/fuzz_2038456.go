package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int16; F1 int8; F2 int16 }
func genFn0(p0 uint32) int {
    if true {
        var v7 uint16 = 255
        switch "a" {
        case "hello":
            v8 := GenS1{F0: uint(v7), F1: -9223372036854775808}
            _ = v8
        default:
            var v9 int = -382
            _ = v9
        }
        _ = v7
    } else {
        var v10 int64 = -1
        var v11 int64 = (v10 + int64(p0))
        _ = v10
        _ = v11
    }
    fmt.Println(p0)
    var v12 uint8 = 42
    _ = p0
    _ = v12
    return 9223372036854775807
}
func genFn1(p0 int64, p1 uint32, p2 uint) int {
    v13 := GenS2{F0: 1, F1: p2}
    var v14 int16 = 42
    fmt.Println(p0, p2, p2)
    for fi15 := 0; fi15 < 7; fi15++ {
        fmt.Println(v14, fi15)
        var v16 int = (fi15 % (((-(fi15))) | 1))
        _ = fi15
        _ = v16
    }
    if (p0 == (p0 ^ p0)) {
        var v17 uint64 = 1
        _ = v17
    } else {
        for fi18 := 0; fi18 < 13; fi18++ {
            fmt.Println(fi18, p0, p2)
            for fi19 := 0; fi19 < 5; fi19++ {
                v20 := GenS1{F0: (^(p2)), F1: (p0 << (((p2 - (p2 % (((p2 & 255)) | 1)))) & 63))}
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        fmt.Println(v14, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return genFn0(p1)
}

func main() {
    var v1 int = 2147483647
    var v2 int32 = -975
    v3 := make(map[uint32]bool)
    v4 := []string{"hello"}
    v5 := make(map[uint32]int)
    var v6 uint64 = 339
    delete(v5, 2)
    var v21 int16 = 32767
    v1 = (v1 ^ genFn1(-1, 0, 65535))
    v5[649] = genFn1(-1, 541, 255)
    v22 := GenS3{F0: (v21 | (^(v21))), F1: 127, F2: (v21 / (((v21 | v21)) | 1))}
    v21 = v21
    v6 -= v6
    v23 := GenS3{F0: v21, F1: 42, F2: (v21 - (^(v21)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
