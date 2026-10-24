package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int; F2 int }

func main() {
    v1 := []string{"", "hello", "foo bar", ""}
    var v2 float64 = 3.14159
    var v3 int = 123
    v4 := make(map[int]string)
    var v5 uint8 = 42
    v2 = -2.25
    v2 = 0.0
    fmt.Println(v2, v5, v2)
    if (-1 <= 1) {
        v3 += v3
    } else {
        _ = len(v4)
        if (true || (false && ((0 > 1) && ("a" <= "foo bar")))) {
            v2 += (3.14159 + v2)
            var v6 uint32 = 0
            _ = v6
        } else {
            fmt.Println(v2, v5)
        }
        v3 += v3
    }
    v7 := GenS1{F0: -1, F1: v3}
    for fi8 := 0; fi8 < 5; fi8++ {
        v5 = 128
        _ = fi8
    }
    v9 := GenS2{F0: true}
    v2 -= v2
    v10 := GenS1{F0: -128, F1: v3}
    for fi11 := 0; fi11 < 13; fi11++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi11
    }
    v12 := GenS1{F0: 85, F1: v3}
    var v13 int8 = -30
    for fi14 := 0; fi14 < 10; fi14++ {
        v13 = (v13 ^ v13)
        v1 = append(v1, "hello")
        _ = fi14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int8 = v13
        v17 := GenS3{F0: 924, F1: fi15, F2: v3}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 string = "a"
    v2 += float64(-2.25)
    v18 = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
