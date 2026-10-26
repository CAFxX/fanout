package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }

func main() {
    var v1 int64 = 0
    v2 := []string{""}
    var v3 string = "foo bar"
    var v4 uint8 = 42
    v5 := make(map[string]bool)
    var v6 uint = 553
    for fi7 := 0; fi7 < 17; fi7++ {
        switch fi7 {
        case 907:
            var v8 uint = v6
            _ = v8
        }
        _ = fi7
    }
    fmt.Println(v3)
    var v9 uint32 = 0
    if (-544 < 127) {
        fmt.Println(v6, v3)
        var v10 int64 = 0
        var v11 bool = (1 > 4294967295)
        _ = v10
        _ = v11
    }
    if (v1 > (v1 ^ 0)) {
        var v12 uint32 = v9
        v1 = 42
        _ = v12
    } else {
        v13 := GenS1{F0: (v6 << (((v6 * v6)) & 63)), F1: 0, F2: (v6 - v6)}
        v4 = (v4 >> ((v6) & 63))
        _ = v13
    }
    v1 -= 1
    v6 = 255
    v6 = uint(v4)
    if (1e10 == (float64((3.14159 + 0.1)) - 1e10)) {
        v14 := GenS1{F0: uint(v4), F1: -2147483648, F2: (v6 & (^(v6)))}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v14
    }
    v4 += v4
    var v15 uint32 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
