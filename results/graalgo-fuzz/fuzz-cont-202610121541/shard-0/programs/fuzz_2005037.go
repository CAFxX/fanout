package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]uint32)
    v3 := []string{}
    var v4 string = ""
    var v5 float32 = float32(0.0)
    var v6 int64 = 0
    for fi7 := 0; fi7 < 17; fi7++ {
        if (v5 <= (float32(-2.25) / v5)) {
            var v8 uint8 = v1
            var v9 float64 = (float64(1.5) + -2.25)
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 2; fi10++ {
            fmt.Println(v5)
            var v11 bool = ((3.14159 / v5) == (v5 * 3.14159))
            fmt.Println(v1, v1, v4)
            _ = fi10
            _ = v11
        }
        if (2 <= 376) {
            v6 = -9223372036854775808
        }
        _ = fi7
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (-1 != 0) {
            v4 = v4
        }
        v5 = (v5 - 3.14159)
    } else {
        v4 = v4
        var v12 int8 = 77
        _ = v12
    }
    v6 = (v6 & 42)
    v13 := GenS1{F0: 0, F1: v4}
    var v14 uint = 0
    var v15 uint32 = 2147483647
    v15 = 431
    for fi16 := 0; fi16 < 14; fi16++ {
        v17 := GenS2{F0: 128, F1: 255}
        var v18 int8 = 42
        delete(v2, "foo bar")
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v6)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 -= (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
