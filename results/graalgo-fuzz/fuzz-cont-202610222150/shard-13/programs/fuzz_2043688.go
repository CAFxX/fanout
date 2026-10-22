package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int32; F1 int8; F2 int64 }
func genFn0() bool {
    var v5 uint16 = 913
    v5 = (v5 >> ((1) & 63))
    var v6 bool = false
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 int64) uint32 {
    var v7 string = "foo bar"
    var v8 uint = 306
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int64 = p0
        v11 := GenS1{F0: 662, F1: 2}
        _ = fi9
        _ = v10
        _ = v11
    }
    v8 -= v8
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int32 = 42
    var v2 uint8 = 93
    var v3 uint = 917
    var v4 uint64 = 18446744073709551615
    switch 9223372036854775807 {
    case -1:
        v12 := GenS1{F0: v4, F1: 2}
        _ = v12
    default:
        v13 := GenS2{F0: 2147483647, F1: 127, F2: int64(v2)}
        _ = v13
    }
    v14 := []int{128, -807, -2147483648}
    var v15 uint = v3
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    if (false && ("hello" > (("" + "\n\t") + "\n\t"))) {
        v4 -= (v4 + v4)
        var v16 uint8 = (v2 ^ v2)
        if true {
            switch 127 {
            case -9223372036854775808:
                v4 = (v4 * v4)
            case -626:
                if false {
                    if false {
                        if true {
                            v17 := GenS1{F0: 951, F1: (^(v2))}
                            var v18 float32 = (3.14159 / float32(0.1))
                            _ = v17
                            _ = v18
                        } else {
                            var v19 int16 = 1
                            var v20 float32 = (0.1 * 1.5)
                            v21 := GenS1{F0: v4, F1: v2}
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    }
                    var v22 uint32 = 2147483648
                    _ = v22
                }
            default:
                _ = len(v14)
            }
        } else {
            var v23 int64 = 429
            v4 -= (v4 >> ((10) & 63))
            _ = v23
        }
        _ = v16
    } else {
        var v24 uint64 = (v4 ^ (v4 / ((v4) | 1)))
        v15 = 18446744073709551615
        _ = v24
    }
    v25 := GenS2{F0: (-(v1)), F1: 1, F2: 255}
    switch 128 {
    case -960:
        v3 += v3
    case -92:
        _ = len(v14)
    }
    v4 -= v4
    v4 += uint64(v2)
    v15 -= (v3 * (v15 * 65535))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
