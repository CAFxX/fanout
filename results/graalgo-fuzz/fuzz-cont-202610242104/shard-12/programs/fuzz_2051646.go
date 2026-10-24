package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int32 }
func genFn0() int64 {
    v5 := []int{}
    switch "foo bar" {
    case "a":
        v6 := GenS1{F0: ("a" > ""), F1: 507, F2: 42}
        _ = v6
    default:
        if len(v5) > 0 {
            v5[0] = 0
        }
    }
    var v7 string = ("\n\t" + "\n\t")
    _ = v5
    _ = v7
    return 1
}
func genFn1(p0 int8) int32 {
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: false, F1: 65535, F2: 0}
        var v10 string = "\n\t"
        var v11 uint32 = 4294967295
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := []int{}
    var v13 uint = 0
    _ = p0
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int64 = -9223372036854775808
    var v3 float64 = ((1.5 / ((1e10 + (float64((float32(float64(1e10)) + 3.14159)) / (-2.25 / float64((1e10 * 1.5))))) / (-2.25 - ((0.1 - (0.1 + 0.1)) + 0.1)))) + -2.25)
    v4 := make(map[int]uint32)
    v3 -= -2.25
    v3 -= (v3 * -2.25)
    var v14 int64 = genFn0()
    switch 298 {
    case 2147483647:
        var v15 uint16 = 2
        _ = v15
    case 32767:
        switch 42 {
        case 1:
            v1 = (v1 * 4294967295)
        default:
            v16 := GenS1{F0: true, F1: 65535, F2: -597}
            _ = v16
        }
    case 127:
        v14 -= 42
    default:
        if true {
            var v17 string = (("foo bar" + "a") + "a")
            switch "foo bar" {
            case "\n\t":
                var v18 uint64 = 2
                _ = v18
            case "":
                v19 := GenS1{F0: false, F1: 1, F2: genFn1(106)}
                _ = v19
            default:
                if ((v17 + v17) >= "a") {
                    fmt.Println(v3, v3, v17)
                } else {
                    v17 = (("" + ("a" + "a")) + ((v17 + (v17 + "foo bar")) + v17))
                    if true {
                        var v20 string = v17
                        v21 := GenS1{F0: false, F1: 65535, F2: -2147483648}
                        v22 := GenS1{F0: true, F1: 32767, F2: genFn1(-1)}
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                }
            }
            _ = v17
        }
    }
    var v23 string = "foo bar"
    fmt.Println(v14, v1)
    v24 := []int{876, 32768}
    v25 := GenS1{F0: true, F1: 0, F2: -1}
    _ = len(v4)
    v4[9223372036854775807] = 0
    if (false && false) {
        switch 32767 {
        case 42:
            _ = len(v4)
        case 601:
            var v26 uint32 = 42
            _ = v26
        default:
            v24 = append(v24, 137)
        }
        v23 = (v23 + v23)
        var v27 int64 = (v14 / (((v14 | (v2 / ((v2) | 1)))) | 1))
        _ = v27
    }
    fmt.Println(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
