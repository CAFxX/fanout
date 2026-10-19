package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint = 1
    v2 := make(map[string]bool)
    var v3 float32 = float32(3.14159)
    var v4 float64 = 0.1
    if false {
        var v5 uint16 = 65535
        _ = v5
    } else {
        v6 := []int{128, -2147483648}
        fmt.Println(v4, v3)
        for fi7 := 0; fi7 < 17; fi7++ {
            fmt.Println(fi7)
            if ((0 > int64(v1)) && (4294967295 > 424)) {
                var v8 uint32 = 2147483647
                _ = v8
            }
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 uint = (v1 - 1)
                _ = fi9
                _ = v10
            }
            _ = fi7
        }
        _ = v6
    }
    switch "" {
    case "a":
        var v11 uint8 = 24
        _ = v11
    case "hello":
        var v12 bool = (128 > 42)
        _ = v12
    case "foo bar":
        var v13 uint8 = 27
        _ = v13
    }
    v2["foo bar"] = true
    for fi14 := 0; fi14 < 12; fi14++ {
        v1 = v1
        _ = fi14
    }
    v3 += ((float32((float64(1.5) - (v4 * v4))) + v3) / (v3 / v3))
    var v15 int8 = 0
    var v16 uint32 = 0
    if (42 == -410) {
        v1 += (v1 + (v1 / ((65535) | 1)))
    } else {
        var v17 uint32 = 2147483648
        fmt.Println(v17, v15, v3)
        if (-803 > 0) {
            var v18 uint64 = 42
            switch "a" {
            case "\n\t":
                v19 := GenS1{F0: 127}
                _ = v19
            case "hello":
                v20 := []int{1, 1}
                _ = v20
            case "":
                v21 := GenS1{F0: 1}
                _ = v21
            }
            _ = v18
        } else {
            v22 := GenS2{F0: 1}
            _ = v22
        }
        _ = v17
    }
    v16 -= 1
    var v23 uint = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
