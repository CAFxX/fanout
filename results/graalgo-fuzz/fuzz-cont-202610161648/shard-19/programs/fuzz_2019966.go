package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int16; F1 string }

func main() {
    var v1 int16 = 42
    v2 := make(map[string]int)
    v3 := []int{32767, -2147483648, -2147483648, int(v1)}
    var v4 uint8 = 169
    var v5 bool = (1 != 1)
    if v5 {
        var v6 uint = 1
        _ = v6
    } else {
        var v7 uint16 = 42
        _ = v7
    }
    v2[("" + "a")] = 128
    var v8 uint8 = v4
    for fi9 := 0; fi9 < 13; fi9++ {
        if len(v3) > 0 {
            v3[0] = fi9
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            v5 = ("foo bar" == ("a" + ("a" + ("hello" + (("foo bar" + "hello") + "foo bar")))))
            v4 -= (v4 ^ (v8 << ((1) & 63)))
            _ = fi10
        }
        _ = fi9
    }
    fmt.Println(v5, v8, v8)
    v1 -= 621
    if v5 {
        fmt.Println(v4, v1, v1)
        v11 := GenS1{F0: 528, F1: v1, F2: 1}
        _ = v11
    }
    delete(v2, "")
    switch "" {
    case "a":
        fmt.Println(v8, v1, v5)
    default:
        v12 := GenS1{F0: 65535, F1: v1, F2: 2147483647}
        _ = v12
    }
    var v13 uint16 = 194
    v14 := GenS2{F0: -39, F1: 1}
    var v15 int = 127
    if ("foo bar" == (("\n\t" + "\n\t") + "\n\t")) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v13)
        v8 = (v4 + (v4 - (v8 / ((1) | 1))))
    } else {
        _ = len(v3)
        v4 += (^(v8))
        var v16 int = (v15 / (((v15 >> ((42) & 63))) | 1))
        _ = v16
    }
    var v17 float32 = 0.1
    v18 := GenS3{F0: v1, F1: "a"}
    v19 := GenS3{F0: (v1 + 0), F1: ("a" + "hello")}
    v2[("a" + "a")] = 128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
