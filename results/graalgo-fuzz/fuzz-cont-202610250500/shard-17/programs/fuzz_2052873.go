package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 bool; F2 int32 }
type GenS5 struct { F0 int8; F1 int }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]bool)
    v3 := []uint32{1, 600, 2147483647, 65535}
    var v4 string = "a"
    var v5 int32 = 2147483647
    v6 := []string{(v4 + "hello"), v4, (v4 + v4)}
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            if (true || (!(-398 >= 32768))) {
                v4 = v4
                fmt.Println(v5, fi7)
                var v9 string = ("a" + (("hello" + "") + v4))
                switch -124 {
                case 1:
                    v10 := GenS1{F0: 1, F1: 555, F2: 226}
                    _ = v10
                case -15:
                    var v11 uint64 = 4294967295
                    _ = v11
                }
                switch fi7 {
                case 938:
                    v12 := GenS1{F0: v5, F1: 734, F2: 255}
                    _ = v12
                case 629:
                    v4 = (("hello" + "a") + (v9 + v9))
                case 42:
                    var v13 float32 = 1.5
                    _ = v13
                default:
                    var v14 uint32 = 561
                    _ = v14
                }
                _ = v9
            }
            fmt.Println(fi7, fi8, fi7)
            _ = fi8
        }
        _ = fi7
    }
    var v15 int32 = (v5 - v5)
    for fi16 := 0; fi16 < 15; fi16++ {
        fmt.Println(v5, v1, fi16)
        fmt.Println(v1, v5)
        v17 := GenS2{F0: fi16}
        _ = fi16
        _ = v17
    }
    if true {
        v5 = v5
    }
    v18 := GenS3{F0: 36}
    v15 = v5
    v19 := GenS4{F0: 246, F1: false, F2: (v5 << ((336) & 63))}
    switch -834 {
    case 1:
        if (false && true) {
            if (9223372036854775807 != 0) {
                var v20 float32 = v1
                var v21 uint8 = 128
                _ = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 20; fi22++ {
                    v23 := v2[2]
                    v24 := v2[4294967295]
                    v2[1] = (v24 || (!true))
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v15 -= (-(v15))
            }
        }
    default:
        v25 := GenS5{F0: 68, F1: 127}
        _ = v25
    }
    fmt.Println(v5, v4)
    var v26 float32 = (3.14159 / v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
