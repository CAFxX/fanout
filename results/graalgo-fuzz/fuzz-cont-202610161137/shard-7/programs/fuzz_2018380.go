package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int32 }
func genFn0(p0 uint8, p1 uint16, p2 uint32) string {
    var v4 float64 = (1e10 / 1.0)
    var v5 int = -1
    var v6 uint8 = p0
    v7 := []int{v5, int(p1), v5}
    v5 = (v5 % ((378) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return "hello"
}
func genFn1() int8 {
    var v8 uint = 255
    if (18446744073709551615 >= 18446744073709551615) {
        for fi9 := 0; fi9 < 4; fi9++ {
            v8 -= 255
            _ = fi9
        }
        switch "\n\t" {
        case "hello":
            var v10 uint16 = 2
            _ = v10
        }
    } else {
        switch "a" {
        case "":
            fmt.Println(v8)
        case "\n\t":
            v8 = v8
        case "foo bar":
            for fi11 := 0; fi11 < 17; fi11++ {
                v8 -= 0
                v12 := GenS1{F0: 1, F1: (fi11 % ((fi11) | 1)), F2: -459}
                _ = fi11
                _ = v12
            }
        }
    }
    _ = v8
    return -128
}

func main() {
    var v1 uint32 = 993
    var v2 uint8 = 2
    var v3 uint32 = v1
    v13 := []int{398}
    var v14 uint16 = 184
    var v15 int32 = 866
    v16 := GenS1{F0: (v2 / ((v2) | 1)), F1: 0, F2: 42}
    switch "foo bar" {
    case "hello":
        var v17 float32 = (0.0 + 3.14159)
        _ = v17
    case "foo bar":
        _ = len(v13)
    case "\n\t":
        if true {
            fmt.Println(v1)
        }
    }
    v15 = int32(v2)
    if len(v13) > 0 {
        v13[0] = -9223372036854775808
    }
    var v18 int = 32768
    v15 -= v15
    if len(v13) > 0 {
        v13[0] = v18
    }
    v1 -= 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}
