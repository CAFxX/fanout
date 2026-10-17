package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
type GenS2 struct { F0 int }
func genFn0(p0 string, p1 int8) string {
    v6 := GenS1{F0: (p1 * p1), F1: -584, F2: 32767}
    switch ("hello" + p0) {
    case "":
        v7 := []int{-490, 385, 632, 488}
        _ = v7
    case "foo bar":
        p0 = (p0 + (p0 + "foo bar"))
    }
    _ = p0
    _ = p1
    _ = v6
    return ""
}

func main() {
    var v1 int8 = 127
    var v2 float64 = 1.5
    var v3 int32 = 243
    var v4 int = 1
    var v5 uint32 = 2147483647
    var v8 string = ("a" + "foo bar")
    fmt.Println(v5, v1, v4)
    if (942 <= 18446744073709551615) {
        v4 = v4
        var v9 int = (v4 * int(v5))
        var v10 float64 = float64(0.0)
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v5, v8)
        _ = fi11
    }
    fmt.Println(v5, v3, v3)
    if false {
        switch 127 {
        case 4294967295:
            var v12 int16 = 0
            _ = v12
        case -481:
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := GenS1{F0: (^(v1)), F1: 1, F2: v4}
                _ = fi13
                _ = v14
            }
        case -22:
            v15 := []int{}
            _ = v15
        default:
            v4 += v4
        }
        v16 := GenS1{F0: 0, F1: -1, F2: v4}
        _ = v16
    } else {
        v17 := GenS1{F0: 0, F1: 1000, F2: -2147483648}
        var v18 int = (v4 - 706)
        _ = v17
        _ = v18
    }
    fmt.Println(v2, v5, v3)
    v19 := GenS2{F0: (^(v4))}
    fmt.Println(v8, v2, v5)
    v4 -= -806
    v8 = ("a" + (v8 + ("foo bar" + (v8 + (("foo bar" + v8) + "foo bar")))))
    var v20 uint32 = (v5 * (v5 * (v5 | 4294967295)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
