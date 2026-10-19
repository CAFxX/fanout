package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint64 }
type GenS2 struct { F0 int64; F1 int64; F2 int }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 string }
func genFn0() bool {
    v6 := GenS1{F0: 2, F1: 877, F2: 292}
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 float32 = 0.1
        v9 := []int{fi7, 32768, (fi7 % (((fi7 << ((1) & 63))) | 1))}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v6
    return (!true)
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]bool)
    var v3 float32 = -2.25
    v4 := []uint32{1, 2, 4294967295}
    v5 := make(map[uint32]uint32)
    v1 = ((v1 + (v1 + v1)) + v1)
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int8 = 44
        for fi12 := 0; fi12 < 19; fi12++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi12
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            if false {
                if (-1 >= -32768) {
                    v14 := GenS2{F0: 335, F1: 863, F2: fi10}
                    var v15 string = v1
                    _ = v14
                    _ = v15
                } else {
                    var v16 uint8 = 142
                    var v17 uint16 = 634
                    switch 9 {
                    case -763:
                        v11 = v11
                    case -9223372036854775808:
                        v18 := GenS2{F0: 1, F1: -510, F2: 160}
                        _ = v18
                    case 9223372036854775807:
                        v1 = ("\n\t" + v1)
                        v16 = 42
                    default:
                        v11 -= v11
                    }
                    _ = v16
                    _ = v17
                }
                switch v1 {
                case "\n\t":
                    v19 := GenS3{F0: -54, F1: 714}
                    _ = v19
                case "a":
                    v20 := GenS4{F0: false, F1: (((v1 + v1) + (v1 + v1)) + (v1 + (((v1 + ("\n\t" + v1)) + ("a" + (v1 + v1))) + v1)))}
                    _ = v20
                }
                v21 := v5[2147483647]
                v11 = (^(v11))
                var v22 uint32 = v21
                _ = v21
                _ = v22
            } else {
                v23 := GenS4{F0: true, F1: ("foo bar" + "foo bar")}
                _ = v23
            }
            fmt.Println(v3)
            _ = fi13
        }
        _ = fi10
        _ = v11
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        v3 -= (((v3 * v3) * v3) / v3)
        var v25 float64 = (0.1 - 0.0)
        v1 = v1
        _ = fi24
        _ = v25
    }
    v3 = -2.25
    var v26 string = ((v1 + v1) + ("" + v1))
    v27 := GenS4{F0: (2 == 867), F1: v1}
    var v28 bool = (!true)
    v3 = (v3 * ((3.14159 / v3) + v3))
    v26 = v26
    v29 := GenS5{F0: (v26 + "a")}
    delete(v5, 65535)
    var v30 uint8 = 2
    fmt.Println(v28)
    v31 := GenS1{F0: 864, F1: 9223372036854775807, F2: 829}
    v32 := GenS4{F0: (!true), F1: v1}
    var v33 uint16 = 283
    v34 := GenS1{F0: v33, F1: 365, F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
