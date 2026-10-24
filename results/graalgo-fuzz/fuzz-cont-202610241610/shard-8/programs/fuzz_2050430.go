package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 string }
type GenS2 struct { F0 int8; F1 uint16; F2 string }
func genFn0(p0 uint) int {
    p0 = 545
    var v6 int16 = 42
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int32 = 42
        v9 := GenS1{F0: false, F1: 9223372036854775807, F2: (((((("foo bar" + "a") + (("" + "") + "foo bar")) + "a") + "\n\t") + "foo bar") + "a")}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p0, p0)
    var v10 uint = p0
    _ = p0
    _ = v10
    _ = v6
    return 0
}

func main() {
    var v1 int = 655
    v2 := make(map[int]bool)
    var v3 int64 = -1
    v4 := make(map[int]int)
    var v5 uint32 = 1
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 string = ""
        _ = fi11
        _ = v12
    }
    switch (((("\n\t" + ("foo bar" + (("a" + "") + "foo bar"))) + "") + "hello") + "") {
    case "foo bar":
        var v13 int64 = v3
        _ = v13
    case "\n\t":
        var v14 float64 = 3.14159
        _ = v14
    case "a":
        fmt.Println(v3, v5, v1)
    default:
        if false {
            v1 = (-(v1))
            var v15 int32 = 1
            _ = v15
        } else {
            for fi16 := 0; fi16 < 7; fi16++ {
                v3 = (v3 << ((199) & 63))
                _ = fi16
            }
            v3 = (v3 * (v3 >> ((931) & 63)))
        }
    }
    v17 := []int{v1, (-(v1))}
    if len(v17) > 0 {
        v17[0] = 128
    }
    if len(v17) > 0 {
        v17[0] = v1
    }
    v3 += (v3 | -1)
    var v18 int8 = -1
    v5 = 65535
    v1 -= 42
    v19 := GenS1{F0: (false && (false && false)), F1: (v3 ^ (v3 >> ((946) & 63))), F2: "foo bar"}
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 float64 = (1.5 + (-2.25 / (float64(1e10) * 0.1)))
        v18 = v18
        _ = fi20
        _ = v21
    }
    v22 := GenS2{F0: v18, F1: 1, F2: "foo bar"}
    var v23 float32 = ((0.0 / (((-2.25 + -2.25) + (-2.25 + (1.5 / (((float32(3.14159) / (0.0 - ((1e10 - 0.1) / 1.5))) + (0.0 - 0.1)) * (1.5 - (0.0 - 1.5)))))) * (0.0 - float32(-2.25)))) - float32(-2.25))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
