package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
func genFn0(p0 int8, p1 bool, p2 string) uint64 {
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := GenS1{F0: 18446744073709551615, F1: 905}
            switch (p2 + p2) {
            case "a":
                p0 += (p0 | (p0 - p0))
            case "hello":
                v9 := GenS1{F0: 4294967295, F1: 684}
                _ = v9
            case "":
                p0 += p0
            default:
                v10 := GenS1{F0: 579, F1: 613}
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        p0 -= p0
        _ = fi6
    }
    var v11 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 1
}

func main() {
    var v1 float64 = float64(-2.25)
    var v2 int32 = 764
    var v3 int = 32768
    var v4 int64 = 54
    var v5 float64 = float64(-2.25)
    fmt.Println(v3)
    if true {
        var v12 float64 = (v5 * -2.25)
        _ = v12
    }
    var v13 uint8 = 2
    var v14 int64 = (^(v4))
    v15 := []int{}
    v14 = 57
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS1{F0: 18446744073709551615, F1: 1}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        if (0 == -1) {
            var v19 int16 = 42
            v4 = v4
            var v20 bool = false
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: 1, F1: 184}
            v3 -= (v3 ^ (-(v3)))
            if (true || false) {
                v4 -= v14
            } else {
                v15 = append(v15, fi18)
                v22 := GenS1{F0: 813, F1: 4294967295}
                switch "a" {
                case "hello":
                    var v23 int8 = 127
                    _ = v23
                case "\n\t":
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                case "a":
                    switch "a" {
                    case "":
                        switch (fi18 * 32767) {
                        case 128:
                            switch "foo bar" {
                            case "":
                                v13 += (v13 | 127)
                            }
                        case 864:
                            v1 = float64((0.0 + (1.5 * 1.5)))
                        }
                    case "a":
                        fmt.Println(v3, v1)
                    default:
                        if len(v15) > 0 {
                            v15[0] = (fi18 / ((632) | 1))
                        }
                    }
                }
                _ = v22
            }
            _ = v21
        }
        _ = fi18
    }
    var v24 float32 = float32(((v5 + 3.14159) * v1))
    if ((75 >= (v13 + (v13 << ((602) & 63)))) || true) {
        var v25 int = (v3 % (((v3 >> ((554) & 63))) | 1))
        _ = v25
    } else {
        if (349 >= 42) {
            var v26 int16 = 32767
            var v27 int8 = -1
            v4 = v14
            _ = v26
            _ = v27
        } else {
            var v28 uint8 = 42
            _ = v28
        }
    }
    v5 = ((v1 / 0.1) + (float64((v24 / (((v24 / float32(v1)) * 0.0) / v24))) * v1))
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
