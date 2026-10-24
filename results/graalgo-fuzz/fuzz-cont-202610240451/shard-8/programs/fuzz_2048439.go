package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 bool; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 0.1
    var v3 int32 = 2147483647
    var v4 int = 4294967295
    if (0 == 128) {
        v1[128] = true
    } else {
        v5 := v1[v4]
        v6 := GenS1{F0: v5, F1: 1}
        _ = v5
        _ = v6
    }
    var v7 int8 = -125
    if (!(true || (v2 < -2.25))) {
        v2 = 1e10
        var v8 uint = 42
        _ = v8
    }
    var v9 uint16 = 972
    v10 := []int{v4, 1}
    v7 -= (v7 * v7)
    switch (^(v4)) {
    case -974:
        var v11 float64 = 1.5
        _ = v11
    case 32768:
        _ = len(v10)
    default:
        v2 = (-2.25 / (1e10 + ((v2 + v2) / (((v2 - 1e10) * 1.5) + v2))))
    }
    v2 += v2
    var v12 int64 = 551
    v13 := GenS2{F0: true, F1: v7}
    if (-2147483648 < v3) {
        v7 -= v7
    } else {
        v1[0] = false
    }
    if true {
        v1[-1] = ("foo bar" < ((("\n\t" + "a") + ("" + ("\n\t" + "\n\t"))) + (((("foo bar" + ("foo bar" + "a")) + ("hello" + ("" + ("foo bar" + "\n\t")))) + "foo bar") + (("hello" + "hello") + (("hello" + "foo bar") + "foo bar")))))
        for fi14 := 0; fi14 < 4; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                v10 = append(v10, 32768)
                _ = fi15
            }
            v7 -= (^(v7))
            var v16 uint = 18446744073709551615
            _ = fi14
            _ = v16
        }
        if len(v10) > 0 {
            v10[0] = 166
        }
    }
    var v17 int16 = -74
    var v18 float64 = (v2 - ((3.14159 / (v2 - 0.0)) * v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
