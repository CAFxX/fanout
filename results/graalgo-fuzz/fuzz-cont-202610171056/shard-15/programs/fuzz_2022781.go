package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint8; F1 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{-878}
    var v3 bool = true
    var v4 int64 = 9223372036854775807
    v4 -= v4
    if v3 {
        switch "\n\t" {
        case "a":
            if false {
                v3 = (148 >= 284)
            } else {
                v5 := GenS1{F0: 4294967295, F1: 755}
                v6 := GenS1{F0: 1, F1: 4294967295}
                _ = v5
                _ = v6
            }
        }
        fmt.Println(v3, v3)
        v4 = (v4 ^ 1)
    } else {
        v4 = v4
    }
    fmt.Println(v4)
    fmt.Println(v3, v3)
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(v4, v3)
        fmt.Println(fi7, fi7, fi7)
        var v8 uint8 = 2
        _ = fi7
        _ = v8
    }
    switch ("" + ("a" + ("foo bar" + (("\n\t" + "a") + "\n\t")))) {
    case "":
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS2{F0: 159, F1: 2}
            _ = fi9
            _ = v10
        }
    }
    if (659 <= 42) {
        _ = len(v2)
    } else {
        var v11 float64 = 0.1
        var v12 int32 = -2147483648
        _ = v11
        _ = v12
    }
    if (false || false) {
        var v13 int16 = 1000
        if v3 {
            v3 = ((!v3) || (!v3))
            var v14 float64 = float64((1.5 / 1.5))
            v3 = ((float64(1.5) == v14) && v3)
            _ = v14
        }
        _ = v13
    }
    v3 = (!(-1 >= 127))
    v1[2147483648] = (true && (v3 && true))
    v2 = append(v2, 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
