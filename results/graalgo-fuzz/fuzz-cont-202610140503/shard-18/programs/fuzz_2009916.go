package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    var v1 float64 = 0.0
    var v2 float64 = v1
    v3 := []int{42, 127, -9223372036854775808}
    v4 := []int{459, -105, 4294967295}
    v5 := make(map[string]bool)
    var v6 bool = true
    v6 = v6
    fmt.Println(v2)
    if v6 {
        if (42 > 18446744073709551615) {
            _ = len(v4)
            var v7 uint = 65535
            var v8 int = 9223372036854775807
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 3; fi9++ {
                _ = len(v4)
                _ = fi9
            }
            var v10 float64 = (v1 / 1.0)
            _ = v10
        }
        v3 = append(v3, -1)
    } else {
        v2 -= ((0.1 / (v2 - 0.0)) / (0.1 - float64(((float32(v2) * (1.5 * (3.14159 + float32(float64((float32((v1 - v2)) + 3.14159)))))) / 1.0))))
        switch "\n\t" {
        case "hello":
            var v11 uint64 = 424
            _ = v11
        }
        v2 += ((float64(1e10) * v2) * (v2 * 0.0))
    }
    v2 = v2
    switch ("foo bar" + ("foo bar" + ("a" + (("\n\t" + "foo bar") + "")))) {
    case "\n\t":
        v12 := GenS1{F0: 1, F1: v6}
        _ = v12
    case "foo bar":
        fmt.Println(v2, v1, v1)
    }
    var v13 int32 = 396
    var v14 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
