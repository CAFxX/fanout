package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int8 }

func main() {
    var v1 uint = 0
    var v2 int8 = -1
    var v3 int64 = 609
    var v4 bool = (!false)
    switch -2147483648 {
    case 127:
        v5 := []int{-1, 0, 32768}
        _ = v5
    case -2147483648:
        v6 := []int{664, 32768, -457, -984}
        _ = v6
    case 208:
        v7 := GenS1{F0: v4, F1: (v4 || v4), F2: (-(v2))}
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        if v4 {
            fmt.Println(fi8, fi8)
        }
        v9 := []int{fi8, (fi8 % ((fi8) | 1)), (fi8 % (((fi8 ^ 32768)) | 1)), int(v2)}
        _ = fi8
        _ = v9
    }
    v4 = (!v4)
    v10 := []int{127, 32767}
    var v11 uint = v1
    v12 := GenS2{F0: 0, F1: 42, F2: 255}
    var v13 int64 = -9223372036854775808
    v14 := GenS3{F0: 534, F1: v2}
    _ = len(v10)
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(fi15, v1, v4)
        var v16 int16 = 42
        _ = fi15
        _ = v16
    }
    v17 := GenS3{F0: 65535, F1: 0}
    v1 -= 703
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
