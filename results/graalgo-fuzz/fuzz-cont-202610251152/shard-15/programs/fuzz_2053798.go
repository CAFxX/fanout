package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
func genFn0() int32 {
    var v5 int32 = -1
    v5 = (v5 ^ 2147483647)
    v5 = v5
    if (!false) {
        if true {
            fmt.Println(v5)
            var v6 float32 = 1e10
            var v7 float32 = (float32((1.5 * 0.0)) * ((((v6 / v6) - (((3.14159 - (1.5 + -2.25)) + v6) + v6)) + 1e10) - v6))
            _ = v6
            _ = v7
        } else {
            v5 -= v5
        }
        var v8 float64 = 0.0
        v8 += float64((-2.25 * 1e10))
        _ = v8
    } else {
        v5 -= 348
        if ("" < "\n\t") {
            v9 := []int{}
            _ = v9
        }
    }
    _ = v5
    return v5
}
func genFn1(p0 uint8, p1 uint16) int8 {
    p0 += p0
    p0 = (p0 >> ((255) & 63))
    _ = p0
    _ = p1
    return -21
}

func main() {
    v1 := []string{"\n\t"}
    var v2 string = "a"
    var v3 uint32 = 2
    v4 := []uint32{}
    v10 := GenS1{F0: -967, F1: v2}
    switch ("a" + v2) {
    case "":
        var v11 uint64 = 0
        _ = v11
    case "foo bar":
        v2 = ("a" + (("hello" + "foo bar") + "hello"))
    }
    var v12 int16 = 6
    var v13 uint16 = 255
    var v14 int32 = 1
    v3 = 52
    if true {
        if (-528 <= -9223372036854775808) {
            for fi15 := 0; fi15 < 9; fi15++ {
                if (v2 >= "hello") {
                    var v16 float64 = 1.5
                    _ = v16
                }
                _ = fi15
            }
        }
        v2 = "foo bar"
        v2 = (((("hello" + v2) + ("\n\t" + v2)) + v2) + v2)
    }
    var v17 uint8 = 42
    for fi18 := 0; fi18 < 8; fi18++ {
        v1 = append(v1, ((v2 + "a") + (v2 + v2)))
        v19 := GenS1{F0: -235, F1: v2}
        _ = fi18
        _ = v19
    }
    v12 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
