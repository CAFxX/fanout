package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint64 }
func genFn0() uint16 {
    var v4 bool = false
    v5 := GenS1{F0: 4294967295, F1: -9223372036854775808, F2: 785}
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int16 = 119
        _ = fi6
        _ = v7
    }
    fmt.Println(v4)
    var v8 string = (("\n\t" + "") + "")
    _ = v4
    _ = v5
    _ = v8
    return 1
}
func genFn1(p0 int64, p1 bool, p2 string) uint32 {
    var v9 int64 = p0
    v10 := GenS1{F0: 0, F1: v9, F2: 933}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 939
}

func main() {
    var v1 float32 = -2.25
    v2 := []string{(("hello" + (("foo bar" + "hello") + ("" + ((((("a" + ("a" + "foo bar")) + ("\n\t" + "\n\t")) + ("\n\t" + "hello")) + "\n\t") + "hello")))) + ((("hello" + ("foo bar" + "a")) + ("hello" + "\n\t")) + "foo bar"))}
    v3 := []uint32{42, 2147483648}
    if false {
        for fi11 := 0; fi11 < 1; fi11++ {
            v1 -= v1
            var v12 float64 = -2.25
            v1 -= ((-2.25 * (0.1 + ((v1 + (v1 * float32((v12 - (((float64(v1) + float64(v1)) + ((float64((0.1 * 3.14159)) / (((1e10 * (v12 - v12)) - v12) / v12)) - (v12 / (v12 * 3.14159)))) + (float64(((0.0 * (v1 / -2.25)) * 1e10)) / 1.5)))))) - (v1 - v1)))) * v1)
            _ = fi11
            _ = v12
        }
    } else {
        v13 := GenS1{F0: 1, F1: 0, F2: 429}
        var v14 int32 = 919
        _ = v13
        _ = v14
    }
    v1 -= v1
    v1 += v1
    if true {
        v1 += ((v1 / v1) - float32(-2.25))
        fmt.Println(v1)
    }
    var v15 int64 = 959
    var v16 uint = 42
    v15 = v15
    var v17 float64 = float64(v1)
    fmt.Println(v15, v15)
    var v18 uint8 = 0
    v18 = v18
    var v19 uint16 = 681
    _ = len(v3)
    switch "hello" {
    case "hello":
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    }
    v16 = 390
    var v20 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
