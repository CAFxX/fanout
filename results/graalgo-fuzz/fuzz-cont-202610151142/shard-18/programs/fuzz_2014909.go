package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint }

func main() {
    v1 := make(map[string]string)
    v2 := []string{"foo bar", ""}
    var v3 int8 = 100
    var v4 float64 = 3.14159
    var v5 float32 = 0.1
    v4 += v4
    v4 -= v4
    if (42 > 42) {
        var v6 int = -2147483648
        _ = v6
    } else {
        var v7 int16 = 1000
        _ = v7
    }
    v4 -= (v4 + 0.0)
    v5 += v5
    delete(v1, ("\n\t" + "hello"))
    var v8 int16 = -386
    var v9 int64 = 9223372036854775807
    fmt.Println(v4, v3, v3)
    if (!true) {
        switch "" {
        case "hello":
            v9 = int64(v8)
        case "":
            fmt.Println(v8, v3)
        }
    } else {
        v10 := GenS1{F0: 2, F1: 1, F2: 4294967295}
        _ = len(v2)
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint32 = 2
            var v13 int16 = v8
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
    }
    v14 := GenS1{F0: 4294967295, F1: 42, F2: 4294967295}
    for fi15 := 0; fi15 < 8; fi15++ {
        v3 -= v3
        _ = fi15
    }
    v4 = float64(1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
