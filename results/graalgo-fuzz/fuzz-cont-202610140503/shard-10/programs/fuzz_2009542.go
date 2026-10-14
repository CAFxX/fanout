package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0() uint16 {
    var v5 string = "a"
    if (((9223372036854775807 == 906) || (1 != 9223372036854775807)) && (-128 >= -104)) {
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 int32 = 0
            _ = fi6
            _ = v7
        }
        var v8 int32 = 2147483647
        fmt.Println(v5)
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 float64 = 3.14159
        if (false || (0 == 2147483647)) {
            var v11 uint64 = 0
            _ = v11
        } else {
            v5 = ((v5 + (v5 + v5)) + v5)
            var v12 int64 = -9223372036854775808
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    _ = v5
    return 65535
}

func main() {
    var v1 int = -817
    var v2 uint32 = 0
    v3 := []int{v1}
    var v4 float32 = float32(((0.0 - (3.14159 - ((0.1 / 1.0) / -2.25))) / 1.5))
    var v13 int8 = -128
    switch 127 {
    case -218:
        if len(v3) > 0 {
            v3[0] = (v1 >> ((18446744073709551615) & 63))
        }
    default:
        v14 := GenS1{F0: (-(v1))}
        _ = v14
    }
    var v15 uint16 = 564
    fmt.Println(v2, v15, v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, (v1 & 4294967295))
    fmt.Println(v4)
    v16 := GenS1{F0: (^(v1))}
    fmt.Println(v13, v1)
    v17 := GenS2{F0: -1, F1: -275}
    var v18 uint8 = 42
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 19; fi20++ {
            if (-364 <= 975) {
                switch "foo bar" {
                case "":
                    if true {
                        v3 = append(v3, (^(v1)))
                        if len(v3) > 0 {
                            v3[0] = fi19
                        }
                    }
                }
                _ = len(v3)
                if len(v3) > 0 {
                    v3[0] = (v1 % (((v1 ^ fi19)) | 1))
                }
            }
            v21 := GenS2{F0: 1, F1: -428}
            _ = fi20
            _ = v21
        }
        fmt.Println(v1, fi19)
        _ = fi19
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 uint8 = (v18 & 42)
        if false {
            fmt.Println(v13, v2, v18)
            var v24 bool = (392 <= 18446744073709551615)
            _ = v24
        }
        var v25 float32 = 3.14159
        _ = fi22
        _ = v23
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
