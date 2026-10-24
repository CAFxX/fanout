package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 string }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int }
func genFn0(p0 int32, p1 bool) string {
    v5 := GenS1{F0: -9223372036854775808, F1: 1, F2: "\n\t"}
    p0 += p0
    var v6 uint16 = 0
    var v7 int32 = 2147483647
    v8 := []int{127, -2147483648, 1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1(p0 int16, p1 string) uint64 {
    var v9 float64 = ((3.14159 / 1.0) - (0.0 - float64((3.14159 + 0.0))))
    var v10 int16 = 1
    v11 := GenS1{F0: -935, F1: 834, F2: "foo bar"}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 0
}

func main() {
    var v1 float64 = (-2.25 - 3.14159)
    var v2 uint = 42
    var v3 uint8 = 1
    v4 := []int{-2147483648, 9223372036854775807}
    for fi12 := 0; fi12 < 4; fi12++ {
        _ = len(v4)
        _ = fi12
    }
    v1 = 0.0
    var v13 int32 = 541
    switch -504 {
    case -58:
        if (!true) {
            var v14 uint64 = 2
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v14
        } else {
            if (!true) {
                if true {
                    var v15 bool = ((!true) && true)
                    for fi16 := 0; fi16 < 8; fi16++ {
                        for fi17 := 0; fi17 < 6; fi17++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi17
                        }
                        _ = fi16
                    }
                    _ = v15
                }
                var v18 float64 = v1
                if len(v4) > 0 {
                    v4[0] = 42
                }
                _ = v18
            }
            if (false || ((-986 == 276) && true)) {
                var v19 int16 = -539
                _ = v19
            }
        }
    case 128:
        var v20 int64 = -1
        _ = v20
    case -2147483648:
        var v21 string = "hello"
        _ = v21
    default:
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 uint16 = 446
            v13 -= v13
            v24 := GenS1{F0: 5, F1: (v23 & (v23 & v23)), F2: ("\n\t" + "hello")}
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        v13 += (v13 % ((42) | 1))
        _ = fi25
    }
    v26 := GenS2{F0: (^(v3)), F1: 42}
    v27 := GenS3{F0: -463}
    v2 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
