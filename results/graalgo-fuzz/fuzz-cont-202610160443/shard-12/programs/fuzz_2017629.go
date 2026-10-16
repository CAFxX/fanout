package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]string)
    var v2 int = 127
    v3 := []string{}
    var v4 int8 = -46
    v2 -= v2
    var v5 bool = (519 >= 2147483648)
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 bool = v5
        v7 = true
        _ = fi6
        _ = v7
    }
    var v8 float32 = 1.5
    switch 229 {
    case -911:
        fmt.Println(v5)
    case -1:
        for fi9 := 0; fi9 < 6; fi9++ {
            if (("hello" <= "hello") || (-9223372036854775808 == -9223372036854775808)) {
                v10 := v1[fi9]
                var v11 float64 = ((((1.5 * (float64(v8) - 1.5)) + ((3.14159 * -2.25) * 1e10)) * 1.5) - (1.5 + 0.0))
                _ = v10
                _ = v11
            }
            _ = fi9
        }
    case 1:
        var v12 int64 = 42
        _ = v12
    default:
        var v13 int16 = 1
        _ = v13
    }
    switch (v2 & -233) {
    case -9223372036854775808:
        for fi14 := 0; fi14 < 6; fi14++ {
            v2 = int(v4)
            var v15 int64 = 0
            fmt.Println(v4)
            _ = fi14
            _ = v15
        }
    case -699:
        v16 := GenS1{F0: 820}
        _ = v16
    case 1:
        v5 = v5
    }
    if (false || ((false && false) && v5)) {
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 uint32 = 65535
            _ = fi17
            _ = v18
        }
    } else {
        v5 = ((!(!v5)) && v5)
        v8 = -2.25
        var v19 uint = 65535
        _ = v19
    }
    v20 := GenS1{F0: 42}
    var v21 uint = 1
    v21 = v21
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
