package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int8 }
type GenS2 struct { F0 int16; F1 int16; F2 bool }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int8; F1 int32; F2 int16 }
func genFn0(p0 int32, p1 uint16) uint32 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int16 = 538
        _ = fi5
        _ = v6
    }
    p1 = 312
    var v7 uint8 = 255
    _ = p0
    _ = p1
    _ = v7
    return 2
}
func genFn1() string {
    v8 := []int{-137}
    if len(v8) > 0 {
        v8[0] = 479
    }
    var v9 string = "foo bar"
    fmt.Println(v9)
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 int = 32768
    v2 := []uint32{4294967295, 4294967295, 161, 65535}
    var v3 int32 = -1
    var v4 int8 = 100
    var v10 uint8 = 0
    v2 = append(v2, 2147483647)
    for fi11 := 0; fi11 < 7; fi11++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v12 int = (v1 + v1)
        if (!(!((v1 | (v12 / ((-739) | 1))) != 0))) {
            if false {
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v1)
                }
            }
            v10 = 1
        } else {
            switch "hello" {
            case "a":
                var v13 float32 = 3.14159
                _ = v13
            case "foo bar":
                v10 += (^(v10))
            default:
                if true {
                    var v14 bool = true
                    v15 := GenS1{F0: "\n\t", F1: v3, F2: -1}
                    var v16 float64 = 3.14159
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    if false {
                        v3 -= v3
                    }
                }
            }
            if (0.0 > float64((0.1 - 0.0))) {
                var v17 int32 = v3
                v18 := GenS2{F0: 42, F1: int16(v10), F2: (v4 != -121)}
                v17 += -2147483648
                _ = v17
                _ = v18
            }
            v10 += v10
        }
        _ = fi11
        _ = v12
    }
    _ = len(v2)
    var v19 bool = (((0.0 + (0.1 / (1.5 * 1.5))) <= 1e10) && true)
    v19 = v19
    fmt.Println(v10, v1, v3)
    var v20 uint64 = 18446744073709551615
    v2 = append(v2, 4294967295)
    v21 := GenS3{F0: v4, F1: (("\n\t" + "foo bar") >= ("foo bar" + "a"))}
    v22 := GenS3{F0: v4, F1: false}
    v23 := GenS4{F0: (v4 >> ((65535) & 63)), F1: (v3 - v3), F2: 1000}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        v1 = -192
        fmt.Println(fi24, v3)
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
