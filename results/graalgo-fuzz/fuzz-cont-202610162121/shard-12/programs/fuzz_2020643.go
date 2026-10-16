package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint8 }

func main() {
    var v1 float64 = (3.14159 * 0.1)
    var v2 bool = (1 < -128)
    v3 := make(map[string]string)
    var v4 uint64 = 18446744073709551615
    if (!v2) {
        var v5 string = "\n\t"
        v6 := []int{2147483647, -904, 32768, 394}
        _ = v5
        _ = v6
    } else {
        fmt.Println(v1, v1, v1)
        v1 = v1
    }
    if true {
        for fi7 := 0; fi7 < 20; fi7++ {
            v1 += v1
            var v8 bool = v2
            for fi9 := 0; fi9 < 3; fi9++ {
                v2 = (-1 > 0)
                if (((fi9 << ((18446744073709551615) & 63)) < fi7) && (true || v2)) {
                    var v10 uint8 = 33
                    v2 = true
                    var v11 uint8 = 25
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
        var v12 int = -9223372036854775808
        _ = v12
    } else {
        v1 -= v1
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := v3["\n\t"]
        _ = fi13
        _ = v14
    }
    v15 := []int{}
    v16 := GenS1{F0: 2147483647, F1: (!true)}
    if (!v2) {
        delete(v3, ("foo bar" + "hello"))
        switch "\n\t" {
        case "a":
            _ = len(v15)
        }
        if v2 {
            v3["foo bar"] = "\n\t"
            if len(v15) > 0 {
                v15[0] = 4294967295
            }
            var v17 uint32 = 2
            _ = v17
        }
    }
    v18 := GenS2{F0: -32, F1: 710}
    v19 := GenS1{F0: -383, F1: (v4 == (v4 | v4))}
    v20 := GenS3{F0: 0, F1: 255}
    v1 = 1.5
    var v21 uint8 = 127
    v21 += v21
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
