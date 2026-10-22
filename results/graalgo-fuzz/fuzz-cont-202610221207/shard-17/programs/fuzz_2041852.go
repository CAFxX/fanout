package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 int8 = 100
    var v4 int32 = 2147483647
    var v5 int = 563
    var v6 uint = 4294967295
    fmt.Println(v3, v4, v3)
    v7 := []int{v5, (v5 / (((v5 ^ v5)) | 1)), -913}
    _ = len(v1)
    var v8 int64 = -1
    switch (-(v5)) {
    case -1:
        switch "" {
        case "a":
            var v9 uint64 = 4294967295
            _ = v9
        }
    case 262:
        v2[("a" + "a")] = ("\n\t" + "a")
    default:
        var v10 uint8 = 128
        _ = v10
    }
    v11 := GenS1{F0: v4, F1: (v6 ^ v6)}
    _ = len(v7)
    var v12 string = (((("\n\t" + "foo bar") + "a") + "a") + (((("a" + ("a" + "a")) + "") + "hello") + "\n\t"))
    var v13 uint64 = 18446744073709551615
    var v14 uint = 275
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS1{F0: v4, F1: (v6 >> ((v14) & 63))}
        v17 := GenS1{F0: v4, F1: v6}
        if true {
            var v18 uint32 = 4294967295
            _ = v18
        } else {
            var v19 float64 = ((1e10 - (1e10 * 1e10)) * 0.1)
            for fi20 := 0; fi20 < 1; fi20++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v21 uint16 = uint16(v13)
                _ = len(v7)
                _ = fi20
                _ = v21
            }
            _ = v19
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    v22 := GenS1{F0: int32(v13), F1: (v6 * v14)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
