package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint16 }
func genFn0(p0 uint64, p1 int8) int {
    var v7 string = "a"
    var v8 uint16 = 42
    for fi9 := 0; fi9 < 11; fi9++ {
        fmt.Println(p0, p0)
        var v10 int = (fi9 % (((fi9 % (((fi9 | (fi9 ^ fi9))) | 1))) | 1))
        v11 := GenS1{F0: (v7 < (v7 + (v7 + v7))), F1: v7, F2: v8}
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := []int{128}
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 bool = (!false)
    var v2 uint64 = 18446744073709551615
    var v3 float32 = -2.25
    var v4 uint = 42
    var v5 int64 = 842
    var v6 string = ("foo bar" + ("a" + "a"))
    if false {
        var v13 int16 = -866
        _ = v13
    }
    v4 -= 42
    v6 = ((v6 + "foo bar") + "\n\t")
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int16 = 0
        if false {
            v2 = 137
        } else {
            v6 = v6
            v3 = (v3 / (float32(3.14159) + 3.14159))
        }
        v6 = ((("foo bar" + v6) + (v6 + (v6 + "foo bar"))) + ((v6 + "a") + "hello"))
        _ = fi14
        _ = v15
    }
    if (v4 != (v4 + (^(v4)))) {
        if (!v1) {
            var v16 int8 = -1
            var v17 uint = (v4 | 125)
            _ = v16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 float64 = 0.0
        var v20 int16 = -768
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := []int{genFn0(v2, 0), -9223372036854775808, genFn0(v2, 0)}
    switch 9223372036854775807 {
    case 989:
        if ((true || (106 <= -128)) && ((v4 < v4) && false)) {
            v6 = "foo bar"
        }
    case 0:
        if v1 {
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 int16 = 32767
                v23 += 42
                _ = fi22
                _ = v23
            }
            if len(v21) > 0 {
                v21[0] = genFn0(v2, 42)
            }
        }
    default:
        v24 := GenS1{F0: v1, F1: "a", F2: 2}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
