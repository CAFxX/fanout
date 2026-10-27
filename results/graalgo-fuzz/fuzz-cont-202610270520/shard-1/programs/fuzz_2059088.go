package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int16; F2 bool }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint32, p1 uint32, p2 int8) int16 {
    var v7 uint8 = 0
    var v8 int8 = p2
    p0 = (p0 * p1)
    var v9 int = 4294967295
    v9 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 1000
}
func genFn1(p0 uint8) int {
    fmt.Println(p0)
    if true {
        p0 = p0
        v10 := GenS1{F0: 1}
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS1{F0: (^(fi11))}
        fmt.Println(p0, p0)
        _ = fi11
        _ = v12
    }
    _ = p0
    return 200
}

func main() {
    var v1 uint64 = 169
    v2 := []uint32{0, 2147483647, 1, 2}
    var v3 int16 = 632
    v4 := make(map[int]bool)
    var v5 bool = true
    var v6 int8 = 127
    v6 = (v6 + (v6 & 0))
    if v5 {
        v13 := GenS1{F0: -19}
        var v14 int8 = (v6 % (((-(v6))) | 1))
        _ = v13
        _ = v14
    }
    if v5 {
        v5 = (820 <= 1)
    }
    var v15 int32 = -2147483648
    _ = len(v4)
    v16 := GenS2{F0: v15, F1: -1, F2: (0 >= -108)}
    v1 = (v1 ^ v1)
    if false {
        fmt.Println(v6, v5)
        var v17 string = ("" + "foo bar")
        var v18 uint8 = 255
        _ = v17
        _ = v18
    }
    var v19 int8 = v6
    for fi20 := 0; fi20 < 6; fi20++ {
        v21 := GenS3{F0: 2, F1: 2}
        if v5 {
            v19 = v6
            var v22 uint16 = 32767
            v23 := GenS2{F0: v15, F1: genFn0(65535, 65535, 42), F2: (1.5 <= (-2.25 - 3.14159))}
            _ = v22
            _ = v23
        } else {
            for fi24 := 0; fi24 < 18; fi24++ {
                v25 := GenS1{F0: (fi24 | (fi20 - fi24))}
                _ = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 9; fi26++ {
                if v5 {
                    var v27 int32 = v15
                    _ = v27
                }
                switch (fi20 + fi26) {
                case 127:
                    switch 128 {
                    case -879:
                        if len(v2) > 0 {
                            v2[0] = 132
                        }
                    case -2147483648:
                        fmt.Println(v3, v6)
                    default:
                        v4[4294967295] = v5
                    }
                default:
                    v6 += (v19 << ((0) & 63))
                }
                _ = fi26
            }
        }
        _ = fi20
        _ = v21
    }
    v19 += (v19 >> ((4294967295) & 63))
    var v28 bool = v5
    v29 := GenS3{F0: 65535, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
