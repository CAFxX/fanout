package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0() string {
    if true {
        var v4 uint64 = 18446744073709551615
        _ = v4
    } else {
        v5 := []int{0}
        var v6 int64 = -9223372036854775808
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(fi7)
            v8 := GenS1{F0: fi7, F1: 65535}
            for fi9 := 0; fi9 < 11; fi9++ {
                switch "hello" {
                case "foo bar":
                    if false {
                        v6 += -429
                    } else {
                        v5 = append(v5, fi9)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        fmt.Println(fi7)
                    }
                default:
                    v10 := GenS1{F0: fi7, F1: 2147483648}
                    _ = v10
                }
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
        _ = v5
        _ = v6
    }
    switch 867 {
    case 0:
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 float32 = float32(0.1)
            var v13 uint = 111
            _ = fi11
            _ = v12
            _ = v13
        }
    case 32767:
        var v14 uint16 = 571
        _ = v14
    case 596:
        if ((-886 == -589) && (2 > 18446744073709551615)) {
            if (0 >= 2) {
                v15 := []int{}
                _ = v15
            }
        }
    }
    var v16 uint32 = 642
    _ = v16
    return ("\n\t" + ("a" + ("" + "")))
}

func main() {
    var v1 int32 = 1
    var v2 bool = false
    v3 := []int{127, 1}
    var v17 float64 = float64(((1.5 - (float32(0.0) * float32((-2.25 * (((3.14159 - 1.5) * 1.5) * ((((0.0 - -2.25) * (1.5 + (float64(((0.1 * 1.5) / -2.25)) + 0.0))) / (float64((float32(-2.25) + (float32(0.0) * 1e10))) * (3.14159 + (((-2.25 / (0.0 - 1.5)) + (float64((1.5 / ((((float32(-2.25) / 3.14159) * 1e10) - (1e10 - float32((0.0 - 3.14159)))) * 0.1))) - (1.5 + float64((float32(1e10) + 0.1))))) / 1e10)))) * 1e10)))))) / 1.5))
    var v18 float32 = 1.5
    fmt.Println(v2)
    var v19 int16 = -636
    var v20 int8 = 21
    fmt.Println(v2, v18)
    var v21 uint8 = 2
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 bool = ((18446744073709551615 == 42) || (843 == 255))
        v23 = v23
        var v24 int16 = (v19 ^ (v19 >> ((255) & 63)))
        _ = fi22
        _ = v23
        _ = v24
    }
    if len(v3) > 0 {
        v3[0] = -196
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
