package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 bool; F1 int }
func genFn0(p0 int32, p1 int64) uint16 {
    var v4 int8 = 1
    var v5 int8 = (v4 * (v4 << ((144) & 63)))
    fmt.Println(p0, v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 int, p1 bool) int8 {
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6, p0)
        var v7 uint = 18446744073709551615
        _ = fi6
        _ = v7
    }
    fmt.Println(p1, p1)
    var v8 bool = ((p0 & (-(p0))) >= (p0 - p0))
    if v8 {
        p0 = 553
        v8 = v8
    } else {
        v9 := GenS1{F0: p0, F1: -1}
        var v10 bool = p1
        _ = v10
        _ = v9
    }
    var v11 int64 = 0
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return -1
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 string = "a"
    var v3 string = ((("a" + "hello") + v2) + (v2 + (v2 + "")))
    v12 := []int{}
    if true {
        v1 = v1
        if true {
            if false {
                v3 = ((v3 + (("" + v2) + v3)) + "hello")
                fmt.Println(v2, v1, v2)
                fmt.Println(v3, v1)
            }
        } else {
            v13 := GenS1{F0: 128, F1: genFn1(1, (!true))}
            fmt.Println(v1, v3)
            switch (v3 + v3) {
            case "":
                v1 -= v1
            case "foo bar":
                _ = len(v12)
            case "a":
                for fi14 := 0; fi14 < 1; fi14++ {
                    _ = len(v12)
                    if len(v12) > 0 {
                        v12[0] = -328
                    }
                    switch v2 {
                    case "foo bar":
                        v15 := GenS1{F0: fi14, F1: 55}
                        _ = v15
                    case "a":
                        v3 = ""
                    case "hello":
                        fmt.Println(fi14, v3, v1)
                    default:
                        v16 := GenS1{F0: fi14, F1: 0}
                        _ = v16
                    }
                    _ = fi14
                }
            }
            _ = v13
        }
    }
    v2 = "a"
    if len(v12) > 0 {
        v12[0] = -9223372036854775808
    }
    var v17 float64 = 0.0
    var v18 int32 = -2147483648
    v19 := GenS2{F0: (-1 <= 127), F1: 1}
    v18 = 42
    v3 = "a"
    v20 := GenS1{F0: 1, F1: genFn1(379, false)}
    for fi21 := 0; fi21 < 19; fi21++ {
        v2 = ""
        v18 += v18
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
