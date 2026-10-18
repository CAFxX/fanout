package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int16 }
func genFn0(p0 string, p1 int32) int8 {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := []int{2147483647}
        if ((false && true) || true) {
            fmt.Println(fi7, p0)
            if len(v8) > 0 {
                v8[0] = int(p1)
            }
        } else {
            p0 = ("" + "hello")
            p0 = p0
            p0 = "a"
        }
        v9 := GenS1{F0: 1, F1: 0, F2: 1000}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: 928, F2: 1000}
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -128
}
func genFn1(p0 uint8) uint32 {
    if true {
        var v12 float32 = float32(0.1)
        _ = v12
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            p0 = (^(p0))
            var v14 int16 = 42
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 int64 = -1
                _ = fi15
                _ = v16
            }
            _ = fi13
            _ = v14
        }
    }
    v17 := GenS1{F0: 417, F1: 1, F2: -1}
    _ = p0
    _ = v17
    return 65535
}

func main() {
    var v1 uint32 = 2147483647
    var v2 float32 = (float32(3.14159) / -2.25)
    v3 := []int{9223372036854775807}
    v4 := make(map[int]string)
    v5 := make(map[int]string)
    v6 := make(map[string]int)
    var v18 int8 = -1
    fmt.Println(v2, v1)
    v18 += (v18 ^ 93)
    v3 = append(v3, -2147483648)
    var v19 uint8 = 128
    v6[((("\n\t" + (("" + ("foo bar" + "\n\t")) + ("a" + "hello"))) + "a") + "hello")] = -1
    v19 = 0
    v20 := GenS1{F0: 492, F1: 42, F2: -53}
    var v21 int = -975
    switch "a" {
    case "foo bar":
        fmt.Println(v21, v21)
    case "hello":
        v19 -= v19
    case "\n\t":
        fmt.Println(v19, v21)
    default:
        var v22 uint = 4294967295
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
