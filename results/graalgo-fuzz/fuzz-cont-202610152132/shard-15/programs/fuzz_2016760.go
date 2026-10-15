package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int16; F1 string }
func genFn0() int16 {
    var v7 float64 = (0.1 * float64(0.1))
    var v8 uint64 = 42
    v7 = -2.25
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int8 = -17
        v8 = v8
        v11 := []int{}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v7
    _ = v8
    return -834
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]int)
    var v3 float32 = float32((3.14159 * (((1.5 - 1.5) + 0.0) - -2.25)))
    var v4 int64 = 9223372036854775807
    v5 := []uint32{2}
    var v6 float64 = float64((1.5 / ((3.14159 / v3) - v3)))
    var v12 int16 = -544
    v3 = v3
    if ((v1 | v1) < (v1 >> (((v1 % (((^(v1))) | 1))) & 63))) {
        fmt.Println(v3, v4)
    }
    switch ("" + "foo bar") {
    case "foo bar":
        var v13 uint32 = 4294967295
        _ = v13
    case "hello":
        v14 := GenS1{F0: v4, F1: v4}
        _ = v14
    default:
        var v15 string = "a"
        _ = v15
    }
    v16 := GenS2{F0: (v12 >> (((v1 - 255)) & 63)), F1: (("hello" + "a") + "")}
    _ = len(v5)
    v6 = v6
    v6 -= v6
    if (!(32767 >= 248)) {
        fmt.Println(v6, v3)
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            if true {
                v6 -= v6
                _ = len(v5)
            } else {
                v3 += v3
                switch ("foo bar" + "") {
                case "a":
                    var v18 int64 = v4
                    _ = v18
                case "\n\t":
                    v4 -= v4
                }
                fmt.Println(v12, fi17)
            }
            _ = fi17
        }
        v6 = (-2.25 / v6)
    }
    _ = len(v5)
    delete(v2, 93)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
