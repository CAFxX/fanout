package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint; F2 int16 }
func genFn0(p0 int64, p1 uint64, p2 int64) bool {
    v5 := []int{-741}
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return false
}
func genFn1() int {
    var v6 int32 = 42
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{fi7, (^(fi7)), -406, (fi7 / ((127) | 1))}
        if len(v8) > 0 {
            v8[0] = fi7
        }
        _ = fi7
        _ = v8
    }
    var v9 int16 = -32768
    fmt.Println(v9, v6)
    _ = v6
    _ = v9
    return 4294967295
}

func main() {
    var v1 int32 = 0
    v2 := []int{128, 128}
    var v3 int16 = -681
    var v4 int = 9223372036854775807
    var v10 int64 = 1
    var v11 uint = 951
    if false {
        v3 = (v3 * (^(v3)))
        var v12 uint32 = 555
        var v13 int64 = v10
        _ = v12
        _ = v13
    }
    v2 = append(v2, v4)
    switch (((("" + ("" + "hello")) + "\n\t") + (("hello" + ("\n\t" + "\n\t")) + "a")) + "foo bar") {
    case "":
        var v14 string = "hello"
        _ = v14
    case "\n\t":
        v15 := GenS1{F0: v4}
        _ = v15
    case "a":
        fmt.Println(v11, v4)
    default:
        v11 -= v11
    }
    switch ("" + "foo bar") {
    case "a":
        if (18446744073709551615 > 42) {
            _ = len(v2)
            var v16 float32 = (0.1 * 3.14159)
            _ = v16
        } else {
            switch (^(v4)) {
            case 128:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 167:
                var v17 uint = v11
                _ = v17
            case -579:
                v10 -= 0
            default:
                var v18 int64 = v10
                _ = v18
            }
            v19 := GenS1{F0: (v4 << (((^(v11))) & 63))}
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v3 += v3
        _ = fi20
    }
    var v21 int = genFn1()
    var v22 bool = (("" >= ("" + "foo bar")) || (!false))
    if len(v2) > 0 {
        v2[0] = genFn1()
    }
    v11 = 4294967295
    _ = len(v2)
    v23 := GenS2{F0: v3, F1: v11, F2: (v3 ^ 0)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
