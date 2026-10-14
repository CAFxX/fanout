package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint8 }
type GenS2 struct { F0 int; F1 int8 }
func genFn0(p0 uint) int16 {
    var v7 uint32 = 2147483647
    var v8 int32 = -448
    v8 = -167
    if false {
        if (!true) {
            v9 := GenS1{F0: 189, F1: 0, F2: 0}
            v10 := []int{42, 1}
            _ = v10
            _ = v9
        }
        var v11 float64 = ((float64(-2.25) * 1.5) - float64(0.1))
        switch 1 {
        case -370:
            v8 -= -1
        default:
            v8 = v8
        }
        _ = v11
    }
    _ = p0
    _ = v7
    _ = v8
    return 697
}

func main() {
    v1 := []uint32{205}
    var v2 uint = 18446744073709551615
    v3 := []uint32{2, 1, 2147483648, 0}
    var v4 float64 = (3.14159 * (1.5 / 3.14159))
    var v5 int64 = 42
    var v6 float64 = 3.14159
    var v12 int8 = 95
    switch -599 {
    case 2147483647:
        v13 := GenS1{F0: 42, F1: -1, F2: 86}
        _ = v13
    case -794:
        var v14 uint8 = 1
        _ = v14
    case 42:
        switch "\n\t" {
        case "a":
            for fi15 := 0; fi15 < 3; fi15++ {
                switch "\n\t" {
                case "hello":
                    v1 = append(v1, 2147483647)
                case "":
                    v16 := GenS1{F0: 1, F1: 1, F2: 1}
                    _ = v16
                case "a":
                    v1 = append(v1, 65535)
                default:
                    var v17 uint = (v2 % ((18446744073709551615) | 1))
                    _ = v17
                }
                _ = fi15
            }
        case "":
            v18 := GenS2{F0: 298, F1: (v12 / (((-(v12))) | 1))}
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v1 = append(v1, 2147483648)
        }
        v20 := GenS2{F0: (-(fi19)), F1: v12}
        v6 = v4
        _ = fi19
        _ = v20
    }
    v5 = v5
    v21 := GenS2{F0: 0, F1: (-(v12))}
    v22 := GenS2{F0: 42, F1: (v12 % (((v12 * (v12 >> ((v2) & 63)))) | 1))}
    v23 := GenS2{F0: 9223372036854775807, F1: (v12 + v12)}
    fmt.Println(v12, v2)
    switch ("" + "") {
    case "":
        v6 = 0.0
    case "a":
        fmt.Println(v4, v12, v6)
    default:
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 int64 = v5
            fmt.Println(v12, v6)
            if len(v3) > 0 {
                v3[0] = 1
            }
            _ = fi24
            _ = v25
        }
    }
    if false {
        if true {
            v5 -= (v5 << (((v2 * v2)) & 63))
        } else {
            fmt.Println(v6, v4)
            var v26 bool = false
            _ = v26
        }
        v27 := GenS2{F0: 956, F1: (v12 / (((v12 - v12)) | 1))}
        var v28 uint = (^(v2))
        _ = v27
        _ = v28
    } else {
        var v29 bool = (0 <= 42)
        var v30 int16 = genFn0(v2)
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 uint8 = 42
            _ = fi31
            _ = v32
        }
        _ = v29
        _ = v30
    }
    switch 32767 {
    case -212:
        v12 = v12
    case 0:
        v2 -= v2
    case -9223372036854775808:
        var v33 int = 32767
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
