package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int32 }
func genFn0() uint {
    v4 := []int{-62, 776}
    if ((818 >= -2147483648) && (1 < 93)) {
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 int8 = 127
            _ = fi5
            _ = v6
        }
        var v7 int32 = 41
        _ = v7
    }
    if true {
        if true {
            var v8 int = 42
            _ = v8
        }
        var v9 bool = false
        _ = v9
    } else {
        if (((0.0 * (float32(0.0) - -2.25)) * (3.14159 * 3.14159)) != (((3.14159 / (3.14159 - 0.1)) + 1.5) * (float32(1e10) / float32((1e10 / ((((0.0 + (-2.25 / 3.14159)) / (3.14159 - -2.25)) + 1.5) + 0.0)))))) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v10 string = "hello"
            v10 = "hello"
            _ = v10
        } else {
            v11 := GenS1{F0: "", F1: 4294967295}
            var v12 int16 = 1
            if true {
                v13 := GenS2{F0: 42}
                switch 1 {
                case 627:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 485:
                    switch -930 {
                    case -2147483648:
                        v12 = (v12 + 32767)
                    case 9223372036854775807:
                        v12 = 788
                    case 1:
                        v12 = (v12 ^ (v12 & v12))
                    default:
                        switch "hello" {
                        case "a":
                            v12 -= (v12 * 496)
                        }
                    }
                case 127:
                    v12 = -283
                }
                _ = v13
            } else {
                v14 := GenS2{F0: 2147483647}
                _ = v14
            }
            _ = v11
            _ = v12
        }
        var v15 int32 = 609
        switch "\n\t" {
        case "a":
            v4 = append(v4, 2147483647)
        case "\n\t":
            v15 -= -812
        case "":
            v4 = append(v4, 0)
        default:
            if len(v4) > 0 {
                v4[0] = 32767
            }
        }
        _ = v15
    }
    _ = v4
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 127
    var v3 uint = 42
    for fi16 := 0; fi16 < 8; fi16++ {
        if false {
            v3 = v3
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := []int{(fi17 >> ((genFn0()) & 63)), (fi16 / ((fi16) | 1)), (fi16 + (fi17 % ((fi17) | 1))), (fi17 * fi16)}
                _ = fi17
                _ = v18
            }
            v3 = v3
        }
        _ = fi16
    }
    v3 += (v3 ^ 4294967295)
    fmt.Println(v3, v3)
    v19 := v1[("a" + "\n\t")]
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 float64 = 1.5
        _ = fi20
        _ = v21
    }
    switch "hello" {
    case "foo bar":
        v22 := []int{}
        _ = v22
    case "":
        v2 += (v2 << (((v3 | 633)) & 63))
    case "\n\t":
        var v23 uint32 = 4294967295
        _ = v23
    default:
        var v24 float32 = ((float32(0.1) / -2.25) / 1.0)
        _ = v24
    }
    if v19 {
        fmt.Println(v19)
    } else {
        v2 -= 0
    }
    var v25 float32 = float32(1e10)
    switch -9223372036854775808 {
    case -712:
        for fi26 := 0; fi26 < 19; fi26++ {
            switch 32768 {
            case 128:
                var v27 string = ("foo bar" + "\n\t")
                _ = v27
            case 1:
                var v28 uint32 = 0
                _ = v28
            }
            var v29 float64 = ((-2.25 * 0.0) - 0.0)
            v30 := GenS2{F0: 1}
            _ = fi26
            _ = v29
            _ = v30
        }
    case -642:
        v31 := GenS1{F0: "", F1: 18446744073709551615}
        _ = v31
    default:
        var v32 int32 = 321
        _ = v32
    }
    var v33 int32 = -78
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v33
}
