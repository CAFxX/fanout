package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 string; F1 string; F2 uint }

func main() {
    v1 := []string{"\n\t", "foo bar", (("\n\t" + "foo bar") + ""), "foo bar"}
    var v2 int = 1
    var v3 float64 = float64(0.1)
    var v4 string = "hello"
    if true {
        var v5 float32 = -2.25
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v6 int = (v2 * v2)
        _ = v5
        _ = v6
    } else {
        v2 = (-(v2))
    }
    v7 := GenS1{F0: 42, F1: -128}
    var v8 int64 = 1
    var v9 uint = 1
    v1 = append(v1, ("foo bar" + ("hello" + "\n\t")))
    v4 = v4
    v10 := GenS2{F0: v4, F1: "foo bar", F2: v9}
    var v11 uint64 = 18446744073709551615
    var v12 string = ((v4 + (v4 + (("a" + v4) + (((v4 + v4) + "a") + v4)))) + v4)
    v3 += 3.14159
    v13 := GenS1{F0: 2, F1: -1}
    v1 = append(v1, "hello")
    var v14 uint32 = 2147483647
    switch (v2 | (v2 ^ 9223372036854775807)) {
    case 32768:
        v8 += v8
    case 9223372036854775807:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 2147483647:
        if false {
            fmt.Println(v14, v12, v12)
        } else {
            fmt.Println(v11, v12, v12)
            var v15 int64 = (v8 << (((v9 / ((v9) | 1))) & 63))
            v2 -= (v2 >> ((uint(v14)) & 63))
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 7; fi16++ {
            v8 = (v8 * 9223372036854775807)
            _ = fi16
        }
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS1{F0: uint32(v11), F1: 72}
        _ = fi17
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
