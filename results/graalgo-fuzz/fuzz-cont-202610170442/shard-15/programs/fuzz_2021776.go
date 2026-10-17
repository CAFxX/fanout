package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint32) uint64 {
    var v6 float32 = (1.5 - float32(1e10))
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 42
    v2 := []uint32{uint32(v1)}
    var v3 uint64 = 2
    var v4 string = ""
    v5 := make(map[uint32]uint32)
    var v7 uint16 = 0
    v7 += (v7 % (((v7 - 2)) | 1))
    for fi8 := 0; fi8 < 8; fi8++ {
        if ((-2.25 * (((3.14159 * float32((1.5 - (0.0 / 1e10)))) * 0.1) - -2.25)) == 3.14159) {
            if len(v2) > 0 {
                v2[0] = 1
            }
            fmt.Println(v3)
            v4 = v4
        } else {
            v9 := GenS1{F0: 482, F1: -107, F2: fi8}
            var v10 int64 = -1
            _ = v10
            _ = v9
        }
        fmt.Println(v4, v7)
        _ = fi8
    }
    v5[0] = uint32(v7)
    var v11 int16 = -1
    var v12 uint16 = (v1 ^ v7)
    for fi13 := 0; fi13 < 5; fi13++ {
        v1 = v1
        _ = fi13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        if true {
            v2 = append(v2, 460)
            v15 := GenS1{F0: 1, F1: int8(v3), F2: fi14}
            _ = v15
        }
        v16 := GenS1{F0: 412, F1: 1, F2: (fi14 | 504)}
        _ = fi14
        _ = v16
    }
    v17 := GenS2{F0: 2147483648}
    if false {
        v3 = v3
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 int = fi18
            var v20 string = v4
            _ = fi18
            _ = v19
            _ = v20
        }
        v21 := GenS3{F0: 2147483647, F1: 287}
        _ = v21
    } else {
        var v22 int = 9223372036854775807
        switch (v22 ^ v22) {
        case -470:
            fmt.Println(v4)
        case 127:
            switch 4294967295 {
            case 251:
                fmt.Println(v4)
            case 9223372036854775807:
                v11 = v11
            case 1:
                v22 = (^(v22))
            default:
                var v23 bool = true
                _ = v23
            }
        default:
            v24 := GenS1{F0: 2, F1: 42, F2: 113}
            _ = v24
        }
        v25 := GenS4{F0: 104}
        _ = v22
        _ = v25
    }
    v12 = (v1 / ((255) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
