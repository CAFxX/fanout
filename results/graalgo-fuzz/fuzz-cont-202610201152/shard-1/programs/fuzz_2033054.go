package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint, p1 int32) int8 {
    switch "hello" {
    case "a":
        fmt.Println(p1, p1)
    case "":
        var v7 bool = (true || (100 == 127))
        _ = v7
    case "\n\t":
        v8 := []int{-9223372036854775808}
        _ = v8
    }
    if ((true || (2 >= 4294967295)) || false) {
        v9 := GenS1{F0: 255, F1: -111, F2: 65535}
        _ = v9
    }
    _ = p0
    _ = p1
    return 100
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]uint32)
    var v3 float64 = 1e10
    var v4 uint = 0
    var v5 int16 = 42
    v6 := []uint32{42, 641, 42}
    v10 := GenS2{F0: 293, F1: 42}
    var v11 uint16 = 0
    v4 = (^(v4))
    switch "foo bar" {
    case "foo bar":
        v12 := GenS2{F0: -1, F1: 42}
        _ = v12
    case "\n\t":
        v4 -= (v4 + v4)
    }
    v13 := GenS3{F0: ("\n\t" + "a"), F1: 255}
    var v14 uint32 = 4294967295
    var v15 uint8 = 0
    fmt.Println(v3, v15)
    if v1 {
        v16 := GenS3{F0: "hello", F1: (^(v15))}
        fmt.Println(v11, v15, v14)
        _ = v16
    }
    delete(v2, (v14 & (v14 - (v14 + 2147483647))))
    v6 = append(v6, v14)
    var v17 uint64 = 1
    for fi18 := 0; fi18 < 17; fi18++ {
        if (!(genFn0(v4, 0) <= -128)) {
            switch int(v4) {
            case -9223372036854775808:
                v19 := GenS2{F0: 446, F1: fi18}
                _ = v19
            case 806:
                var v20 float32 = float32(v3)
                _ = v20
            case 127:
                var v21 string = (("hello" + "foo bar") + "hello")
                _ = v21
            default:
                v15 += v15
            }
            for fi22 := 0; fi22 < 7; fi22++ {
                fmt.Println(v11, v17, v15)
                var v23 int8 = -1
                if v1 {
                    fmt.Println(v3, v3)
                } else {
                    v2[807] = v14
                }
                if v1 {
                    fmt.Println(v3, v17)
                    v24 := GenS4{F0: -9223372036854775808, F1: (v15 % (((v15 << (((v4 % ((uint(v17)) | 1))) & 63))) | 1))}
                    fmt.Println(v4, v11)
                    _ = v24
                }
                _ = len(v6)
                var v25 uint8 = v15
                v26 := GenS5{F0: genFn0(uint(v14), 1)}
                _ = fi22
                _ = v23
                _ = v25
                _ = v26
            }
            v27 := GenS1{F0: v15, F1: -43, F2: 2147483647}
            _ = v27
        }
        _ = fi18
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
