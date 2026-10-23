package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 bool }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint32, p1 uint32) int16 {
    v7 := GenS1{F0: -2147483648}
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 float32 = 1e10
        v10 := GenS2{F0: 919, F1: 221, F2: false}
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(p1, p1)
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v7
    return 32767
}
func genFn1(p0 int8, p1 uint, p2 int) bool {
    p0 = (^(p0))
    v11 := []int{}
    var v12 uint32 = 96
    var v13 int16 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return false
}

func main() {
    var v1 int16 = 1
    var v2 uint8 = 1
    var v3 int8 = -128
    v4 := make(map[uint32]uint32)
    var v5 float32 = float32((((3.14159 / 1.0) + 0.0) - ((float64(((0.0 / 1.5) / 1e10)) / 1e10) / 1.0)))
    v6 := make(map[int]string)
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(v5, fi14)
        for fi15 := 0; fi15 < 4; fi15++ {
            v16 := GenS2{F0: 799, F1: 209, F2: false}
            v17 := []int{(fi14 + int(v1))}
            switch "" {
            case "foo bar":
                var v18 int64 = -9223372036854775808
                _ = v18
            case "hello":
                fmt.Println(v3)
            case "":
                v2 -= v2
            default:
                var v19 string = (("\n\t" + "") + "a")
                _ = v19
            }
            if false {
                var v20 bool = (!false)
                _ = v20
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        var v21 int16 = -1
        _ = fi14
        _ = v21
    }
    v22 := GenS3{F0: 1}
    v2 += (v2 - (v2 | (v2 % ((0) | 1))))
    delete(v4, 693)
    v23 := []int{9223372036854775807, -1}
    var v24 int16 = (v1 >> ((203) & 63))
    var v25 int64 = 0
    var v26 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
