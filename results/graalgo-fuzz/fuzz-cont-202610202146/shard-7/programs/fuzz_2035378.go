package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int8; F1 int32; F2 uint }
type GenS3 struct { F0 string; F1 uint16; F2 uint8 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    v4 := GenS1{F0: 128, F1: -128}
    v1 += 42
    var v5 int = 88
    v6 := []int{v5, v5, 980, v5}
    switch "\n\t" {
    case "foo bar":
        v3[("foo bar" + ("a" + "\n\t"))] = "foo bar"
    }
    var v7 int16 = 893
    v8 := GenS2{F0: -1, F1: -2147483648, F2: 255}
    var v9 uint = 354
    if false {
        v10 := GenS3{F0: "foo bar", F1: 255, F2: 127}
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 float64 = 3.14159
            v13 := GenS1{F0: 127, F1: 0}
            v6 = append(v6, -2147483648)
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint = 8
        _ = fi14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
