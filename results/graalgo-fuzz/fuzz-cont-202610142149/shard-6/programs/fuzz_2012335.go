package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }
type GenS4 struct { F0 string }

func main() {
    v1 := []int{0, -69}
    v2 := []int{-368, 220, 128, -2147483648}
    var v3 uint = 4294967295
    v4 := make(map[uint32]bool)
    var v5 int32 = 0
    var v6 int = 740
    fmt.Println(v5, v5)
    v6 = (v6 + (v6 & (v6 << (((v3 - v3)) & 63))))
    if (255 <= 588) {
        v7 := GenS1{F0: v6}
        switch (v6 - (v6 + 762)) {
        case -2147483648:
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 int16 = 48
                var v10 uint64 = 1
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        switch v6 {
        case 643:
            v5 = v5
        case -461:
            v3 += v3
        case 1:
            var v11 bool = (v6 >= (v6 + (v6 & 42)))
            _ = v11
        default:
            if ((v5 | v5) == 878) {
                v3 -= (v3 / (((v3 ^ v3)) | 1))
            }
        }
        _ = v7
    }
    var v12 uint64 = 4294967295
    var v13 int32 = v5
    v5 += v5
    switch (v6 * (^(v6))) {
    case 128:
        v14 := GenS2{F0: 255, F1: 0}
        _ = v14
    case -720:
        if (2147483648 != 1) {
            var v15 float64 = 0.1
            _ = v15
        }
    case 4294967295:
        _ = len(v4)
    }
    _ = len(v1)
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS3{F0: v6, F1: -1, F2: uint64(v3)}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            v18 := GenS2{F0: 255, F1: v6}
            for fi19 := 0; fi19 < 2; fi19++ {
                v12 = 4294967295
                _ = len(v1)
                _ = fi19
            }
            _ = v18
        } else {
            fmt.Println(fi16)
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v3, v6)
    if ("\n\t" != "foo bar") {
        var v20 int64 = -1
        var v21 uint64 = (^(v12))
        var v22 uint8 = 128
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: v6}
    var v24 uint8 = 127
    var v25 uint64 = (v12 ^ 904)
    var v26 int32 = (v5 ^ (v5 - 42))
    switch v6 {
    case -94:
        var v27 uint16 = 42
        _ = v27
    case 9223372036854775807:
        if (!false) {
            v4[4294967295] = true
        } else {
            for fi28 := 0; fi28 < 12; fi28++ {
                v29 := GenS4{F0: "hello"}
                _ = fi28
                _ = v29
            }
            v13 += int32(v24)
            delete(v4, uint32(v24))
        }
    case -9223372036854775808:
        v4[65535] = false
    default:
        v6 += v6
    }
    var v30 uint32 = 193
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
