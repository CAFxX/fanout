package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0() uint64 {
    var v7 int64 = 1
    var v8 uint = 0
    _ = v7
    _ = v8
    return 296
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -1
    var v3 float64 = 1e10
    var v4 uint16 = 1
    var v5 int16 = (^(v2))
    var v6 int16 = v2
    var v9 float32 = ((3.14159 / (3.14159 + 1e10)) + -2.25)
    v10 := GenS1{F0: 757, F1: 0}
    for fi11 := 0; fi11 < 11; fi11++ {
        v4 -= (v4 ^ v4)
        var v12 float64 = (-2.25 * (v3 / 1.0))
        v9 -= v9
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 496, F1: 897}
    var v14 int16 = 1000
    for fi15 := 0; fi15 < 2; fi15++ {
        switch (fi15 + int(v4)) {
        case -9223372036854775808:
            var v16 float64 = (v3 - (v3 + ((v3 / v3) + v3)))
            _ = v16
        default:
            if false {
                fmt.Println(v4)
            } else {
                v17 := GenS1{F0: 0, F1: 42}
                _ = v17
            }
        }
        _ = fi15
    }
    v18 := []int{1, -607, -586, -2147483648}
    fmt.Println(v3)
    var v19 string = ("foo bar" + "hello")
    for fi20 := 0; fi20 < 11; fi20++ {
        v6 = -114
        v19 = v19
        fmt.Println(v6, v2)
        _ = fi20
    }
    switch "hello" {
    case "":
        for fi21 := 0; fi21 < 20; fi21++ {
            if (false || false) {
                var v22 uint = 18446744073709551615
                _ = v22
            } else {
                v6 += (v14 + 0)
                v5 = (v14 ^ v14)
            }
            _ = fi21
        }
    case "a":
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    case "foo bar":
        v23 := GenS1{F0: 42, F1: 2}
        _ = v23
    default:
        v1[42] = (("foo bar" + "foo bar") + "\n\t")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
