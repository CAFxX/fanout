package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"hello", "hello"}
    v3 := make(map[uint32]bool)
    var v4 uint32 = 398
    var v5 uint = 192
    var v6 float64 = (3.14159 / 3.14159)
    v6 += v6
    for fi7 := 0; fi7 < 8; fi7++ {
        delete(v3, 0)
        v8 := GenS1{F0: v5, F1: 4294967295, F2: 42}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: v5, F1: (^(v5)), F2: 0}
    var v10 int8 = -1
    v10 += v10
    switch "foo bar" {
    case "":
        v2 = append(v2, "hello")
    }
    if (-15 > 128) {
        if ((v5 - (v5 / ((4294967295) | 1))) > 2) {
            v6 += ((v6 / 0.1) + (0.0 * float64(0.1)))
            v4 -= v4
        } else {
            v4 = (v4 | v4)
            var v11 int64 = 9223372036854775807
            var v12 int = 0
            _ = v11
            _ = v12
        }
        v10 -= (^(v10))
    } else {
        v4 -= v4
        var v13 string = ""
        _ = v13
    }
    if false {
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS1{F0: v5, F1: (v5 << (((v5 << ((v5) & 63))) & 63)), F2: (v10 | v10)}
            _ = fi14
            _ = v15
        }
        v4 = (v4 + 1)
    }
    fmt.Println(v5, v6)
    v4 = v4
    switch "\n\t" {
    case "\n\t":
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 int16 = 225
            _ = fi16
            _ = v17
        }
    case "foo bar":
        var v18 string = ""
        _ = v18
    case "":
        v5 = 2
    default:
        v19 := v3[(v4 ^ 2147483648)]
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
