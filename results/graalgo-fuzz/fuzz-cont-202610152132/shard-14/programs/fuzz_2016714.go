package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    var v1 float64 = 0.0
    var v2 uint64 = 1
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]bool)
    v5 := []uint32{2}
    var v6 uint64 = v2
    v6 = (v2 << ((488) & 63))
    if false {
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(v2, fi7, v1)
            v8 := v4[1]
            _ = fi7
            _ = v8
        }
        var v9 int16 = -1
        _ = v9
    }
    v2 -= (v2 * 0)
    v6 = (v6 & (^(v6)))
    var v10 uint64 = v6
    if false {
        _ = len(v5)
        v6 -= (v10 - v2)
    }
    var v11 float32 = (1e10 + 1e10)
    if true {
        v12 := GenS1{F0: 717, F1: 128}
        if true {
            switch "a" {
            case "":
                var v13 int32 = -41
                _ = v13
            case "foo bar":
                v14 := GenS1{F0: 65535, F1: 128}
                _ = v14
            case "\n\t":
                var v15 int32 = -2147483648
                _ = v15
            }
            v1 = ((1.5 - 0.1) - ((((v1 / (v1 * float64(v11))) - float64(3.14159)) + 1.5) * v1))
        } else {
            v1 += ((((1.5 * (v1 / 1e10)) + v1) + float64(v11)) * (v1 / v1))
            var v16 uint64 = (v6 / (((v10 >> ((1) & 63))) | 1))
            _ = v16
        }
        v6 = 18446744073709551615
        _ = v12
    }
    var v17 string = "a"
    v18 := v3[600]
    switch 127 {
    case 128:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 42:
        v18 = ""
    case 830:
        var v19 bool = (444 != 1)
        _ = v19
    }
    v6 = 457
    var v20 bool = (-1 == 9223372036854775807)
    var v21 uint8 = 69
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
