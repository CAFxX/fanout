package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int16 }

func main() {
    var v1 bool = (128 > 72)
    v2 := []string{"\n\t", "a", "hello", ("hello" + "")}
    v3 := []uint32{2147483648, 573, 2147483648}
    v4 := make(map[uint32]int)
    var v5 int64 = 0
    v6 := []uint32{1}
    var v7 int32 = -612
    v1 = true
    v7 = (-(v7))
    v1 = v1
    v5 -= (v5 % ((42) | 1))
    if false {
        fmt.Println(v1, v7)
    }
    if v1 {
        v8 := GenS1{F0: (v5 ^ (-(v5)))}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v9 float32 = (3.14159 + -2.25)
        _ = v8
        _ = v9
    } else {
        v5 = (v5 * v5)
        var v10 float32 = 1.5
        _ = v10
    }
    v11 := GenS1{F0: (-(v5))}
    var v12 uint8 = 160
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v13 := v4[2147483648]
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS1{F0: (v5 & (v5 % (((v5 | (v5 ^ -9223372036854775808))) | 1)))}
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: uint32(v12), F1: 1, F2: 2147483648}
        v18 := GenS1{F0: (-(v5))}
        _ = fi14
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := GenS2{F0: 1, F1: -128, F2: 4294967295}
            if (1 != 2) {
                v1 = false
            } else {
                v7 -= (^(v7))
                v22 := GenS3{F0: 2147483647, F1: -495}
                _ = v22
            }
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
