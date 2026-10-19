package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 198
    v3 := []uint32{1, 2147483647, 65535}
    v4 := make(map[string]bool)
    var v5 int8 = 77
    v6 := []uint32{65535, 0, 2147483648}
    v5 -= 127
    fmt.Println(v2, v2)
    if (2 < 42) {
        v5 -= v5
        fmt.Println(v5)
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 uint8 = 127
            _ = fi7
            _ = v8
        }
    } else {
        delete(v4, "\n\t")
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            _ = len(v3)
            switch "\n\t" {
            case "\n\t":
                var v11 float32 = 3.14159
                _ = v11
            case "hello":
                v2 += -2147483648
                var v12 bool = false
                v3 = append(v3, 0)
                var v13 int64 = 9223372036854775807
                _ = v12
                _ = v13
            default:
                v2 = (v2 / (((v2 / ((978) | 1))) | 1))
            }
            _ = fi10
        }
        var v14 float64 = -2.25
        v15 := GenS1{F0: 666, F1: 356}
        _ = fi9
        _ = v14
        _ = v15
    }
    var v16 float64 = 1.5
    var v17 int32 = (v2 ^ (v2 % ((v2) | 1)))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v18 := GenS1{F0: 2147483648, F1: 42}
    var v19 int8 = (v5 + (v5 + v5))
    delete(v1, "foo bar")
    var v20 int8 = (-(v5))
    switch 2147483647 {
    case 128:
        if (!true) {
            fmt.Println(v16)
        }
    default:
        v20 += (v19 * (v5 % ((0) | 1)))
    }
    v19 = v19
    var v21 uint32 = 65535
    var v22 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
