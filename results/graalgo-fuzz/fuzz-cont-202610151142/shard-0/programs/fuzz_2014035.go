package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    var v1 string = "a"
    v2 := []int{2147483647, 418}
    var v3 float64 = ((3.14159 * (((1.5 / 1e10) - ((0.1 - 1e10) / 1.0)) - (float64((-2.25 + 0.0)) + -2.25))) / (1.5 * -2.25))
    var v4 uint8 = 8
    var v5 string = ""
    var v6 int = 4294967295
    var v7 uint8 = v4
    fmt.Println(v1)
    for fi8 := 0; fi8 < 12; fi8++ {
        v7 += v4
        _ = fi8
    }
    _ = len(v2)
    switch v1 {
    case "foo bar":
        var v9 bool = (false || true)
        _ = v9
    case "a":
        v3 += v3
    case "":
        v1 = v5
    }
    v10 := GenS1{F0: v4, F1: -9223372036854775808}
    if true {
        switch v6 {
        case -1:
            var v11 uint8 = v7
            _ = v11
        }
        if false {
            var v12 uint16 = 1
            v13 := GenS1{F0: (v4 | v7), F1: (v6 / ((v6) | 1))}
            _ = v12
            _ = v13
        }
        fmt.Println(v4, v7)
    } else {
        fmt.Println(v3, v5, v4)
        var v14 uint32 = 34
        _ = v14
    }
    var v15 float32 = (1e10 - -2.25)
    var v16 int32 = -1
    if (v16 <= (v16 >> ((150) & 63))) {
        var v17 uint16 = 72
        _ = v17
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v6 += v6
    fmt.Println(v6)
    v18 := GenS1{F0: v4, F1: v6}
    v19 := GenS1{F0: (^(v4)), F1: 128}
    v20 := GenS1{F0: v7, F1: 9223372036854775807}
    if (2 > 18446744073709551615) {
        if true {
            v5 = ""
            var v21 uint32 = 2
            _ = v21
        }
        v2 = append(v2, (-(v6)))
        v22 := GenS1{F0: v4, F1: (v6 + v6)}
        _ = v22
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
