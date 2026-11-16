package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int32 }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
type GenS3 struct { F0 int8; F1 int; F2 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 1
    var v3 uint64 = 0
    var v4 uint32 = 0
    var v5 int = 533
    v6 := make(map[int]int)
    if true {
        fmt.Println(v5, v5)
        _ = len(v6)
        fmt.Println(v4, v5)
    }
    var v7 uint = 318
    _ = len(v6)
    switch "\n\t" {
    case "a":
        v5 = (v5 << (((v7 & 0)) & 63))
    case "hello":
        var v8 uint8 = 0
        _ = v8
    case "":
        if (255 > 556) {
            switch v5 {
            case 182:
                if false {
                    var v9 bool = (!(534 != uint32(v3)))
                    _ = v9
                }
            case -2147483648:
                v10 := []int{(v5 & (^(v5)))}
                _ = v10
            default:
                if false {
                    var v11 uint32 = (v4 ^ (v4 + 65535))
                    fmt.Println(v5, v5)
                    v5 -= v5
                    _ = v11
                }
            }
        } else {
            v5 -= v5
        }
    }
    switch (^(v5)) {
    case 2147483647:
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := []int{4294967295, (v5 * v5), fi12, (fi12 * (^(v5)))}
            var v14 float64 = 0.0
            _ = fi12
            _ = v13
            _ = v14
        }
    case 1:
        if false {
            v15 := GenS1{F0: int64(v5), F1: "a", F2: int32(v3)}
            fmt.Println(v2, v2)
            _ = v15
        } else {
            v5 = (v5 << (((v7 ^ (v7 - (^(v7))))) & 63))
            var v16 uint64 = v3
            _ = len(v1)
            _ = v16
        }
    case 127:
        var v17 uint = (v7 | (v7 | (v7 % ((v7) | 1))))
        _ = v17
    }
    fmt.Println(v3)
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 int8 = -128
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: 0, F1: 292, F2: int(v4)}
    for fi21 := 0; fi21 < 10; fi21++ {
        fmt.Println(fi21, fi21)
        if (!(false && (2 < 0))) {
            switch "hello" {
            case "hello":
                fmt.Println(v4, fi21, v3)
            }
        }
        v22 := GenS3{F0: -128, F1: v5, F2: (v5 >> ((751) & 63))}
        _ = fi21
        _ = v22
    }
    var v23 uint32 = v4
    if false {
        var v24 uint32 = v4
        v2 = v2
        v25 := GenS3{F0: 127, F1: (-(v5)), F2: v5}
        _ = v24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
