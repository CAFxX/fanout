package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int8 }
func genFn0(p0 int64, p1 int32) int8 {
    v6 := GenS1{F0: 1}
    var v7 float64 = -2.25
    fmt.Println(p0, p0, v7)
    v8 := GenS1{F0: 0}
    v9 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 100
}

func main() {
    var v1 int32 = 1
    v2 := []int{}
    v3 := make(map[uint32]uint32)
    var v4 int8 = 57
    v5 := make(map[int]int)
    v10 := GenS2{F0: 590}
    switch 9223372036854775807 {
    case 657:
        v4 -= 100
    default:
        fmt.Println(v1)
    }
    var v11 int16 = -769
    v4 = genFn0(0, (v1 % ((v1) | 1)))
    var v12 int = 127
    v5[v12] = 9223372036854775807
    v11 -= 32767
    var v13 float32 = 1.5
    v11 += (v11 + (v11 - v11))
    if true {
        if (("" + ((((("hello" + ("a" + "hello")) + ("foo bar" + "hello")) + "\n\t") + "a") + ("hello" + ("hello" + (("a" + ((("hello" + "foo bar") + "hello") + (("\n\t" + "a") + "a"))) + "\n\t"))))) <= ("a" + "a")) {
            v13 -= (0.0 * float32(1e10))
            v1 = (v1 & -406)
        }
        fmt.Println(v1)
    } else {
        switch -2147483648 {
        case -2147483648:
            for fi14 := 0; fi14 < 12; fi14++ {
                if len(v2) > 0 {
                    v2[0] = v12
                }
                v11 = (v11 * (-(v11)))
                _ = fi14
            }
        }
        v11 = 269
    }
    var v15 int16 = v11
    v16 := GenS1{F0: 1}
    var v17 bool = false
    if false {
        var v18 int16 = (v15 / ((32767) | 1))
        var v19 int64 = 185
        _ = v18
        _ = v19
    }
    v13 = ((v13 - ((0.1 - v13) - v13)) + float32(((0.0 * (3.14159 * ((0.0 / 1e10) - 0.1))) - (0.1 - 1e10))))
    fmt.Println(v12, v17)
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := GenS3{F0: v17, F1: (^(v4))}
        if ((((18446744073709551615 < 1) && (255 == 908)) || ((v12 + (fi20 / ((fi20) | 1))) == v12)) || true) {
            v15 -= -730
            var v22 int = v12
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    v12 += v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
