package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint8 }
func genFn0() int16 {
    switch 431 {
    case -284:
        var v4 string = "hello"
        _ = v4
    case -2147483648:
        var v5 uint64 = 798
        _ = v5
    default:
        switch ("foo bar" + ("a" + "hello")) {
        case "hello":
            var v6 uint8 = 1
            _ = v6
        case "\n\t":
            v7 := []int{32767, -813, 127, 2147483647}
            _ = v7
        case "":
            v8 := []int{-1}
            _ = v8
        }
    }
    var v9 uint8 = 241
    v10 := []int{-2147483648}
    if len(v10) > 0 {
        v10[0] = 32767
    }
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 int32 = -2147483648
    var v2 float32 = (-2.25 + float32(((1e10 + 0.1) + 1e10)))
    v3 := []int{-323, 2147483647, -136}
    v1 -= (v1 / ((2147483647) | 1))
    fmt.Println(v1)
    v11 := GenS1{F0: 2}
    var v12 uint = 174
    if (93 <= 120) {
        for fi13 := 0; fi13 < 17; fi13++ {
            fmt.Println(v1, v12, fi13)
            _ = fi13
        }
    } else {
        switch (("a" + ("\n\t" + ("" + "hello"))) + "") {
        case "hello":
            if true {
                var v14 float32 = v2
                _ = v14
            }
        case "\n\t":
            v1 = (v1 | v1)
        case "":
            if false {
                fmt.Println(v2, v2)
                var v15 uint = 472
                _ = v15
            }
        default:
            switch "foo bar" {
            case "foo bar":
                if false {
                    var v16 uint32 = 0
                    v1 = int32(v16)
                    _ = v16
                }
            case "\n\t":
                v3 = append(v3, 400)
            default:
                _ = len(v3)
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1, v12, v1)
    switch "foo bar" {
    case "foo bar":
        v17 := GenS1{F0: 0}
        _ = v17
    case "":
        v3 = append(v3, 42)
    }
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    fmt.Println(v2, v2, v2)
    v18 := GenS2{F0: "foo bar", F1: 127}
    var v19 int16 = 42
    v2 = (float32(0.1) - float32(0.1))
    var v20 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
