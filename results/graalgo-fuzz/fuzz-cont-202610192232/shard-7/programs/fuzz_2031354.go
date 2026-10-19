package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint }

func main() {
    var v1 int16 = 239
    v2 := []string{}
    v3 := make(map[int]uint32)
    var v4 int16 = v1
    var v5 string = ""
    var v6 int = 658
    v4 += int16(v6)
    v1 = (-(v1))
    v5 = v5
    for fi7 := 0; fi7 < 19; fi7++ {
        v4 += (v1 / (((v4 | 42)) | 1))
        _ = fi7
    }
    if (-125 > 42) {
        v8 := GenS1{F0: 127, F1: 2147483648, F2: 255}
        v6 = -233
        _ = v8
    } else {
        fmt.Println(v1)
        var v9 int16 = v4
        if ((v6 != v6) || false) {
            var v10 uint32 = 65535
            v9 = (v9 + -488)
            v11 := GenS1{F0: 1, F1: v10, F2: 623}
            _ = v10
            _ = v11
        } else {
            switch ("\n\t" + (v5 + "a")) {
            case "a":
                var v12 float64 = 0.0
                _ = v12
            }
            var v13 int16 = (^(v4))
            v4 += (v9 * v1)
            _ = v13
        }
        _ = v9
    }
    v6 = 128
    v2 = append(v2, v5)
    var v14 int64 = 1
    var v15 uint16 = 42
    if ((127 == 59) || true) {
        fmt.Println(v6, v4)
    } else {
        var v16 int = v6
        _ = v16
    }
    fmt.Println(v1)
    v15 -= (v15 + (^(v15)))
    v17 := GenS1{F0: -128, F1: 2147483648, F2: 65535}
    v18 := GenS1{F0: -11, F1: 2, F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
