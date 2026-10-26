package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }
func genFn0(p0 uint8, p1 int, p2 uint8) uint32 {
    if (-1 >= 100) {
        p1 += 9223372036854775807
        if (0.1 <= 3.14159) {
            var v5 int16 = 1000
            var v6 int32 = 42
            var v7 int32 = v6
            _ = v5
            _ = v6
            _ = v7
        }
        if false {
            p2 -= p2
        }
    }
    p2 = 92
    _ = p0
    _ = p1
    _ = p2
    return 687
}
func genFn1() int64 {
    if (1.5 < 3.14159) {
        var v8 int = 620
        switch (v8 - v8) {
        case -9223372036854775808:
            v9 := []int{(v8 << ((2) & 63)), (v8 ^ v8), v8}
            _ = v9
        default:
            v10 := []int{}
            _ = v10
        }
        v11 := []int{v8, (v8 >> ((255) & 63)), v8, v8}
        _ = v11
        _ = v8
    }
    if (!true) {
        var v12 float64 = float64(0.0)
        var v13 int64 = 821
        var v14 uint64 = 42
        _ = v12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 4294967295, F1: 42, F2: -1}
    v16 := []int{127, -495, 42, 127}
    _ = v15
    _ = v16
    return 9223372036854775807
}

func main() {
    var v1 uint8 = 7
    var v2 int16 = 1000
    v3 := []string{}
    var v4 int = -873
    v4 += (v4 + 32767)
    var v17 uint64 = 0
    fmt.Println(v4)
    var v18 uint = 1
    switch "a" {
    case "foo bar":
        v17 += v17
    case "a":
        for fi19 := 0; fi19 < 17; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := GenS1{F0: v18, F1: (^(v2)), F2: (v4 >> ((909) & 63))}
                v3 = append(v3, ("foo bar" + "\n\t"))
                var v22 uint8 = 2
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi19
        }
    case "\n\t":
        var v23 int32 = int32(v2)
        _ = v23
    default:
        v18 = v18
    }
    v1 = 2
    _ = len(v3)
    fmt.Println(v18, v4, v18)
    _ = len(v3)
    _ = len(v3)
    var v24 int8 = -128
    switch ("\n\t" + (("hello" + ((("\n\t" + "") + "a") + "hello")) + "foo bar")) {
    case "":
        v2 += -266
    case "a":
        var v25 bool = true
        _ = v25
    default:
        fmt.Println(v18, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
