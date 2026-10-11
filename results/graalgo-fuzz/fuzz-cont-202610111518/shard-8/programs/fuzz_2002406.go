package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0(p0 int8) int {
    fmt.Println(p0)
    var v6 int32 = 2147483647
    v6 -= v6
    _ = p0
    _ = v6
    return -1
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{4294967295, 91}
    var v3 uint8 = 127
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]uint32)
    v3 += v3
    var v7 uint64 = 18446744073709551615
    var v8 uint16 = 61
    fmt.Println(v7, v8, v8)
    var v9 uint8 = 114
    if false {
        switch 4294967295 {
        case 450:
            if true {
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
            } else {
                v9 -= (v3 - (v3 ^ 174))
            }
        case 9223372036854775807:
            var v10 uint16 = (v8 >> ((20) & 63))
            _ = v10
        default:
            v7 = (v7 / (((^(v7))) | 1))
        }
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 float64 = 0.0
            var v13 int8 = -128
            _ = len(v4)
            _ = fi11
            _ = v12
            _ = v13
        }
        v1["foo bar"] = ("hello" + (("\n\t" + ("foo bar" + "foo bar")) + ""))
    } else {
        switch ("a" + "") {
        case "foo bar":
            v3 = v9
        }
    }
    switch "" {
    case "hello":
        v14 := GenS1{F0: -1, F1: -74}
        _ = v14
    case "a":
        if false {
            var v15 uint64 = (^(v7))
            var v16 int32 = 42
            _ = v15
            _ = v16
        } else {
            var v17 uint = 65535
            _ = v17
        }
    case "":
        if false {
            if true {
                if len(v2) > 0 {
                    v2[0] = 65535
                }
                if (!(int64(v8) <= -1)) {
                    v3 -= (v9 - v3)
                    fmt.Println(v9, v3, v8)
                    var v18 int32 = 545
                    _ = v18
                } else {
                    v7 -= v7
                }
                if (!false) {
                    if false {
                        if (!((0.0 == -2.25) || (32767 != 206))) {
                            var v19 int8 = -1
                            _ = v19
                        }
                    } else {
                        var v20 uint = 1
                        _ = v20
                    }
                    fmt.Println(v7, v3)
                }
            }
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 float64 = ((float64(float32((1e10 + 0.0))) - (3.14159 * float64(1.5))) + (1.5 + 1.5))
                fmt.Println(v7, v3)
                _ = len(v2)
                _ = fi21
                _ = v22
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        if (false && (-743 > -675)) {
            var v24 int16 = 32767
            v24 += v24
            _ = v24
        }
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
