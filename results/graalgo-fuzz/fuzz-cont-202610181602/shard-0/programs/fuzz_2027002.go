package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int; F1 uint16 }
func genFn0(p0 string, p1 int16) int8 {
    v7 := GenS1{F0: 0, F1: 2, F2: 933}
    p0 = ("\n\t" + (p0 + "a"))
    p0 = ""
    _ = p0
    _ = p1
    _ = v7
    return 100
}

func main() {
    v1 := []int{2147483647, -2147483648, 1}
    var v2 int32 = 42
    var v3 int = 9223372036854775807
    var v4 uint8 = 0
    var v5 int8 = 127
    v6 := make(map[uint32]uint32)
    for fi8 := 0; fi8 < 20; fi8++ {
        if (1 >= 255) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi8
    }
    if (42 != 42) {
        _ = len(v1)
        fmt.Println(v4, v5)
        if false {
            if true {
                v9 := v6[772]
                var v10 uint = 18446744073709551615
                _ = v10
                _ = v9
            }
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v11 := GenS2{F0: v5, F1: 18446744073709551615, F2: (v4 | (v4 + v4))}
        fmt.Println(v3)
        _ = v11
    }
    var v12 float32 = -2.25
    fmt.Println(v4)
    if false {
        var v13 int32 = (v2 | v2)
        _ = len(v1)
        _ = v13
    }
    var v14 uint64 = 18446744073709551615
    var v15 int = 210
    v16 := GenS3{F0: uint64(v4), F1: 9223372036854775807}
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint32 = 2147483647
        v19 := GenS4{F0: v3, F1: 382}
        v15 = fi17
        _ = fi17
        _ = v18
        _ = v19
    }
    v4 = (v4 << ((604) & 63))
    switch (-(v15)) {
    case -9223372036854775808:
        switch v15 {
        case 0:
            v14 = v14
        }
    case 32767:
        if false {
            v20 := v6[869]
            if false {
                var v21 bool = ((v3 * (v15 ^ v15)) != v3)
                fmt.Println(v14, v4)
                _ = v21
            } else {
                var v22 int = (^(v15))
                fmt.Println(v4, v22)
                _ = v22
            }
            _ = v20
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v15 -= v15
            v4 += 42
        }
    case 42:
        v12 = (v12 / v12)
    default:
        v2 -= (v2 >> ((42) & 63))
    }
    v3 += int(v4)
    fmt.Println(v4, v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
