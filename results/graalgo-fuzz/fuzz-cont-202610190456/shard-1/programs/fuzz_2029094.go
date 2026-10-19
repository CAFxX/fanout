package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := []uint32{61, 638}
    v2 := []int{127, 2147483647}
    v3 := make(map[string]bool)
    var v4 string = (("foo bar" + ("foo bar" + ((((("hello" + "hello") + "foo bar") + "hello") + (("" + "foo bar") + "a")) + "a"))) + "foo bar")
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint16 = 2
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 977, F1: 1000}
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int16 = -978
        v4 = "a"
        var v11 int32 = 1
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: -128}
    var v13 int = 0
    delete(v5, 969)
    var v14 int32 = -573
    var v15 bool = (42 > -546)
    if false {
        if len(v1) > 0 {
            v1[0] = 1
        }
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        v4 = v4
        v13 = fi16
        var v17 int = 72
        _ = fi16
        _ = v17
    }
    v15 = ((!(42 == 529)) && v15)
    v4 = "hello"
    switch ("\n\t" + v4) {
    case "\n\t":
        v13 -= v13
    case "foo bar":
        var v18 float64 = ((0.0 + (0.1 / 1.0)) + (((1e10 * 1e10) + 1e10) + 0.0))
        _ = v18
    case "a":
        v19 := GenS1{F0: 2147483648, F1: int16(v13)}
        _ = v19
    }
    var v20 bool = v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
