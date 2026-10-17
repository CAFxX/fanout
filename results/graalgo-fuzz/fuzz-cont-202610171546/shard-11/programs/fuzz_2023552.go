package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0(p0 uint32, p1 uint64) uint {
    v5 := []int{127, 2147483647}
    var v6 uint16 = 991
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 int16) int {
    switch 127 {
    case 1:
        p0 = p0
    case 854:
        var v7 bool = false
        _ = v7
    default:
        if (1000 != (p0 % ((p0) | 1))) {
            var v8 int8 = 1
            p0 = p0
            fmt.Println(v8)
            _ = v8
        }
    }
    var v9 uint16 = 513
    var v10 uint32 = 2147483647
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(fi11, v10)
        v12 := GenS1{F0: (fi11 / (((fi11 / ((fi11) | 1))) | 1)), F1: 65535}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v9
    return 127
}

func main() {
    v1 := []string{}
    var v2 float64 = 3.14159
    v3 := make(map[int]bool)
    v4 := []string{"foo bar", "foo bar"}
    fmt.Println(v2)
    var v13 uint32 = 0
    v14 := GenS1{F0: 817, F1: 568}
    if len(v1) > 0 {
        v1[0] = "a"
    }
    var v15 float64 = ((1.5 + (0.0 + (v2 + 1e10))) * v2)
    var v16 int16 = -32768
    fmt.Println(v16, v16, v15)
    switch -2147483648 {
    case -1:
        v17 := GenS1{F0: -472, F1: (v13 - 42)}
        _ = v17
    case -909:
        switch "a" {
        case "":
            if ((42 >= 0) && (981 > 209)) {
                fmt.Println(v16, v15)
                fmt.Println(v13, v15)
            }
        case "hello":
            if true {
                var v18 uint16 = 853
                _ = v18
            } else {
                v2 = v2
                var v19 uint32 = 426
                for fi20 := 0; fi20 < 11; fi20++ {
                    fmt.Println(v2, v13)
                    _ = fi20
                }
                _ = v19
            }
        case "foo bar":
            if (!(-2.25 == 1.5)) {
                v15 = 0.0
                v4 = append(v4, "\n\t")
                var v21 int16 = v16
                _ = v21
            }
        }
    case 437:
        v15 += float64((0.0 - 3.14159))
    default:
        switch "\n\t" {
        case "hello":
            var v22 uint8 = 1
            _ = v22
        case "":
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        }
    }
    v13 = (v13 * (v13 * (v13 ^ v13)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
