package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint64 }
func genFn0(p0 uint32, p1 uint16, p2 uint16) int8 {
    p0 = 0
    v6 := []int{0, 32768, -2147483648, 1}
    var v7 int = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 uint16 = 65535
    v2 := []uint32{2, 42, 705, 663}
    var v3 float32 = (3.14159 * 0.1)
    v4 := make(map[uint32]int)
    var v5 int16 = -742
    var v8 uint32 = 0
    v1 -= 32767
    delete(v4, (v8 & v8))
    var v9 uint64 = 780
    var v10 int = 9223372036854775807
    for fi11 := 0; fi11 < 19; fi11++ {
        v3 -= 1.5
        var v12 float32 = (v3 / v3)
        var v13 int64 = -9223372036854775808
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 127}
    var v15 uint32 = 16
    switch v10 {
    case 117:
        var v16 int = (^(v10))
        _ = v16
    case 9223372036854775807:
        v15 = 0
    case 0:
        for fi17 := 0; fi17 < 8; fi17++ {
            fmt.Println(v9, v1, fi17)
            _ = fi17
        }
    default:
        for fi18 := 0; fi18 < 19; fi18++ {
            v10 = fi18
            v19 := GenS2{F0: 1}
            for fi20 := 0; fi20 < 19; fi20++ {
                v9 = v9
                switch v10 {
                case 0:
                    v21 := GenS3{F0: "a", F1: (v9 ^ 18446744073709551615)}
                    _ = v21
                case -9223372036854775808:
                    fmt.Println(fi18, v15)
                }
                var v22 bool = (0 == -751)
                _ = fi20
                _ = v22
            }
            _ = fi18
            _ = v19
        }
    }
    v10 = 32767
    for fi23 := 0; fi23 < 9; fi23++ {
        fmt.Println(v5, v15, v5)
        if (v9 > (v9 / ((v9) | 1))) {
            switch fi23 {
            case -2147483648:
                if (v10 < (v10 << ((42) & 63))) {
                    fmt.Println(v8, v9, v3)
                    v4[4294967295] = (fi23 % (((^(fi23))) | 1))
                    var v24 uint16 = v1
                    _ = v24
                }
            }
            v9 = v9
        }
        _ = fi23
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 int16 = (-(v5))
        _ = fi25
        _ = v26
    }
    v8 = (^(v8))
    _ = len(v2)
    fmt.Println(v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
