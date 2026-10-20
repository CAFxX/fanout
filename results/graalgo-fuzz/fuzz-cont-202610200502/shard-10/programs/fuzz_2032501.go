package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint64; F1 int16; F2 bool }
type GenS4 struct { F0 uint64; F1 int16; F2 uint64 }

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 int64 = 42
    if (4294967295 <= 2) {
        if ("a" != "\n\t") {
            v5 := []int{}
            v6 := GenS1{F0: (v1 % (((v1 | v1)) | 1)), F1: 0, F2: "a"}
            _ = v5
            _ = v6
        } else {
            if false {
                switch -2147483648 {
                case 127:
                    v7 := []int{966}
                    _ = v7
                case -9223372036854775808:
                    v1 += v1
                }
            } else {
                fmt.Println(v1, v1)
                var v8 bool = true
                _ = v8
            }
            for fi9 := 0; fi9 < 11; fi9++ {
                v1 = v1
                var v10 uint16 = 0
                v1 = (v1 / (((v1 >> ((657) & 63))) | 1))
                _ = fi9
                _ = v10
            }
        }
        var v11 int8 = (v1 / ((v1) | 1))
        if false {
            _ = len(v3)
        }
        _ = v11
    }
    var v12 uint8 = 2
    v4 = v4
    v13 := GenS2{F0: "\n\t", F1: 42}
    if true {
        v12 -= v12
        v14 := GenS1{F0: v1, F1: 209, F2: "hello"}
        v15 := []int{-347, -2147483648}
        _ = v14
        _ = v15
    }
    v1 += v1
    var v16 int64 = 327
    v16 = -73
    v12 -= v12
    var v17 uint32 = 4294967295
    var v18 int32 = -416
    if (2 <= uint(v17)) {
        switch -1 {
        case 0:
            for fi19 := 0; fi19 < 19; fi19++ {
                v4 = 42
                var v20 float32 = float32(3.14159)
                _ = fi19
                _ = v20
            }
        case 919:
            delete(v3, "a")
        case -499:
            v17 = (v17 | 101)
        }
        for fi21 := 0; fi21 < 11; fi21++ {
            v16 = v16
            v22 := GenS3{F0: 1, F1: 42, F2: true}
            _ = fi21
            _ = v22
        }
    }
    v17 = (^(v17))
    switch 0 {
    case 42:
        fmt.Println(v17, v17)
    case -2147483648:
        v23 := GenS4{F0: 18446744073709551615, F1: -1, F2: 271}
        _ = v23
    default:
        v16 = -202
    }
    var v24 uint64 = 42
    v25 := v3["a"]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
