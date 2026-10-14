package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int16 }
type GenS2 struct { F0 int; F1 uint16; F2 uint16 }
func genFn0(p0 string, p1 int64, p2 int64) string {
    if false {
        var v6 uint8 = 253
        _ = v6
    }
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = 2147483647
    }
    if len(v7) > 0 {
        v7[0] = -1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ((p0 + (p0 + ((p0 + (("foo bar" + "") + (p0 + "foo bar"))) + ("\n\t" + p0)))) + "")
}

func main() {
    var v1 uint = 491
    var v2 float32 = 1e10
    v3 := []int{}
    var v4 float64 = (0.1 * (float64((float32(0.0) * 3.14159)) + (1.5 / (0.1 - ((0.0 + 3.14159) + 3.14159)))))
    var v5 int32 = 2147483647
    if len(v3) > 0 {
        v3[0] = 1
    }
    v8 := GenS1{F0: 1, F1: -128, F2: 42}
    for fi9 := 0; fi9 < 2; fi9++ {
        v5 = v5
        var v10 string = ""
        v5 = 1
        _ = fi9
        _ = v10
    }
    var v11 int64 = -183
    if (520 != 0) {
        var v12 float32 = v2
        _ = v12
    }
    var v13 uint16 = 0
    v14 := GenS2{F0: -174, F1: v13, F2: v13}
    fmt.Println(v1)
    var v15 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
