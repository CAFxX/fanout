package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint16 }
func genFn0(p0 uint32, p1 uint8) uint {
    if true {
        var v6 string = "hello"
        var v7 int64 = 0
        var v8 uint16 = 2
        _ = v6
        _ = v7
        _ = v8
    }
    fmt.Println(p1, p0)
    p1 = (p1 | 169)
    _ = p0
    _ = p1
    return 74
}

func main() {
    var v1 int = -2147483648
    var v2 int16 = 1000
    var v3 uint64 = 1
    var v4 int8 = -128
    var v5 uint = 18446744073709551615
    fmt.Println(v4, v5, v1)
    if true {
        v4 += (^(v4))
    } else {
        var v9 uint8 = 127
        var v10 uint64 = (v3 ^ v3)
        if true {
            var v11 int = (-(v1))
            for fi12 := 0; fi12 < 8; fi12++ {
                v10 -= (v3 << ((2) & 63))
                _ = fi12
            }
            var v13 uint32 = 692
            _ = v11
            _ = v13
        } else {
            v3 += (v10 % ((v10) | 1))
        }
        _ = v10
        _ = v9
    }
    v5 -= (v5 / ((917) | 1))
    v14 := []int{}
    v2 += 1000
    var v15 uint32 = 42
    var v16 uint64 = (v3 * 602)
    if true {
        if (2 < 984) {
            fmt.Println(v15)
            v16 = 328
        } else {
            switch 32767 {
            case 1:
                v14 = append(v14, v1)
            default:
                var v17 uint8 = 2
                _ = v17
            }
            v3 += (^(v16))
            v1 = (v1 ^ v1)
        }
    }
    fmt.Println(v5, v2, v4)
    var v18 uint8 = 138
    v19 := GenS1{F0: int16(v4), F1: v15, F2: 32767}
    if len(v14) > 0 {
        v14[0] = (v1 / ((int(v15)) | 1))
    }
    var v20 uint64 = v3
    v15 = v15
    switch v1 {
    case 500:
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    case 170:
        if false {
            v1 += (v1 << ((1) & 63))
            if true {
                _ = len(v14)
            }
        } else {
            v5 = (v5 & 641)
        }
    default:
        if len(v14) > 0 {
            v14[0] = -799
        }
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 string = ("a" + "a")
        var v23 int = (fi21 ^ (fi21 ^ (v1 << ((genFn0((v15 / ((v15) | 1)), uint8(v15))) & 63))))
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
