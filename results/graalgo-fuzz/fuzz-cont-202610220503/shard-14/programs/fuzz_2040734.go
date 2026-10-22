package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint16; F2 string }
type GenS4 struct { F0 int16; F1 uint8 }
func genFn0(p0 int8) int16 {
    p0 += -98
    p0 -= (p0 << ((2) & 63))
    fmt.Println(p0)
    fmt.Println(p0)
    var v4 bool = (4294967295 >= 341)
    _ = p0
    _ = v4
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 2
    v3 := make(map[string]uint32)
    var v5 int32 = -1
    v3["hello"] = 1
    if true {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 uint = uint(v2)
            v8 := GenS1{F0: (fi6 & -9223372036854775808), F1: (v5 / (((v5 ^ (-(v5)))) | 1)), F2: (v2 ^ v2)}
            _ = fi6
            _ = v7
            _ = v8
        }
        fmt.Println(v5)
    }
    v9 := GenS1{F0: 42, F1: -2147483648, F2: v2}
    if (1 > 17) {
        var v10 uint16 = 42
        v11 := GenS2{F0: v2, F1: 0}
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int16 = -32768
            _ = fi12
            _ = v13
        }
        _ = v10
        _ = v11
    }
    v14 := []int{}
    var v15 float32 = 0.1
    for fi16 := 0; fi16 < 12; fi16++ {
        for fi17 := 0; fi17 < 8; fi17++ {
            if ((!(2 < 42)) && false) {
                v18 := GenS2{F0: 255, F1: 65535}
                if len(v14) > 0 {
                    v14[0] = -660
                }
                if (("\n\t" > ("a" + "")) && (((3.14159 * -2.25) + (1e10 - 1.5)) <= (3.14159 / (3.14159 * 0.1)))) {
                    var v19 bool = true
                    switch fi16 {
                    case 128:
                        switch fi16 {
                        case 0:
                            if (v19 || (true || false)) {
                                var v20 float64 = (0.0 / 1.0)
                                v20 = 0.0
                                v15 -= (v15 - v15)
                                _ = v20
                            }
                        }
                    default:
                        v19 = (!false)
                    }
                    _ = v19
                }
                _ = v18
            } else {
                v21 := GenS3{F0: -192, F1: 2, F2: "foo bar"}
                _ = v21
            }
            v5 += (-(v5))
            _ = fi17
        }
        _ = fi16
    }
    if (136 != 814) {
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 bool = (759 < 42)
            _ = fi22
            _ = v23
        }
    } else {
        fmt.Println(v15)
        var v24 int8 = -68
        var v25 uint16 = 65535
        _ = v24
        _ = v25
    }
    v26 := GenS4{F0: 625, F1: 255}
    var v27 uint32 = 775
    v2 = v2
    if len(v14) > 0 {
        v14[0] = 32768
    }
    var v28 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v5
    _ = v9
}
