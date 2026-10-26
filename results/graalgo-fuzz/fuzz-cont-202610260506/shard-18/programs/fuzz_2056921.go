package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[int]string)
    var v3 int8 = -128
    var v4 float32 = float32((3.14159 / -2.25))
    var v5 uint32 = 2
    v6 := make(map[string]bool)
    if false {
        for fi7 := 0; fi7 < 1; fi7++ {
            fmt.Println(v5, v4, v5)
            v5 += 0
            _ = fi7
        }
        var v8 uint32 = (v5 + (^(v5)))
        var v9 uint = 4294967295
        _ = v8
        _ = v9
    } else {
        fmt.Println(v4, v5)
    }
    v3 = -118
    fmt.Println(v4)
    if ((3.14159 - (1e10 - (1.5 / (1.5 / 0.1)))) > 1.5) {
        v10 := GenS1{F0: -1}
        _ = v10
    } else {
        switch "\n\t" {
        case "hello":
            var v11 int32 = -68
            _ = v11
        default:
            var v12 uint = 0
            _ = v12
        }
        if (!false) {
            var v13 float64 = float64(v4)
            _ = v13
        } else {
            var v14 int16 = 1
            switch ("" + "foo bar") {
            case "":
                v15 := GenS2{F0: -69}
                _ = v15
            case "\n\t":
                var v16 int64 = 42
                _ = v16
            case "a":
                v6["hello"] = false
            }
            _ = v14
        }
        switch 32768 {
        case 32767:
            var v17 uint8 = 2
            _ = v17
        }
    }
    var v18 int32 = int32(v5)
    v19 := []int{128, 0, 1}
    for fi20 := 0; fi20 < 17; fi20++ {
        if false {
            var v21 float64 = -2.25
            var v22 int32 = v18
            _ = v21
            _ = v22
        }
        _ = fi20
    }
    v5 -= (v5 & v5)
    v1 -= (^(v1))
    var v23 uint64 = 18446744073709551615
    if true {
        fmt.Println(v23)
        for fi24 := 0; fi24 < 16; fi24++ {
            v5 = v5
            v1 -= v1
            _ = fi24
        }
    }
    v5 = (v5 & (^(v5)))
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 int64 = -9223372036854775808
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
