package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int16 }

func main() {
    var v1 uint16 = 65535
    var v2 float64 = (((-2.25 - (1e10 - (3.14159 * 1e10))) + (-2.25 * 0.1)) / 3.14159)
    v3 := make(map[int]uint32)
    var v4 uint16 = (v1 % (((v1 - 681)) | 1))
    var v5 uint64 = 226
    v4 = (^(v4))
    v6 := GenS1{F0: 30, F1: 65535, F2: 0}
    var v7 float64 = float64(0.0)
    v8 := GenS2{F0: 42}
    var v9 uint16 = v4
    v2 -= v2
    var v10 int32 = -749
    fmt.Println(v5)
    fmt.Println(v5, v7)
    var v11 uint32 = 65535
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        if (!true) {
            v1 = (v4 ^ v9)
        } else {
            var v12 int16 = 0
            _ = v12
        }
    case "":
        for fi13 := 0; fi13 < 20; fi13++ {
            if (128 < uint8(v1)) {
                v9 += v1
            }
            _ = fi13
        }
    default:
        v14 := []int{42}
        _ = v14
    }
    v15 := GenS1{F0: 100, F1: (v11 / ((v11) | 1)), F2: -128}
    if false {
        v3[2147483647] = 65535
    }
    v10 += v10
    var v16 uint32 = 42
    if true {
        var v17 int8 = 0
        fmt.Println(v9, v10)
        var v18 uint32 = v16
        _ = v17
        _ = v18
    }
    v9 += v1
    if (32767 < 128) {
        v19 := GenS1{F0: 0, F1: 42, F2: 0}
        _ = len(v3)
        _ = v19
    } else {
        for fi20 := 0; fi20 < 8; fi20++ {
            if false {
                fmt.Println(v4, v9, v2)
                fmt.Println(v11, fi20, v10)
                var v21 int = -647
                _ = v21
            } else {
                v22 := GenS3{F0: -1, F1: -987}
                v23 := []int{fi20, (fi20 << ((2) & 63))}
                v4 = v1
                _ = v22
                _ = v23
            }
            v24 := []int{}
            v25 := GenS3{F0: 1, F1: 32767}
            _ = fi20
            _ = v24
            _ = v25
        }
        v16 -= (v16 >> ((502) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
