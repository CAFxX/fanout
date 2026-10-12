package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint32, p1 string, p2 uint16) int16 {
    fmt.Println(p0)
    p0 += (p0 >> ((65535) & 63))
    v4 := GenS1{F0: -2147483648, F1: 0}
    fmt.Println(p2, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 619
}
func genFn1(p0 uint) int {
    v5 := []int{-525, 32768, 150}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 float32 = 0.1
    var v2 int32 = 42
    v3 := make(map[string]uint32)
    var v6 int16 = 1
    fmt.Println(v1)
    v7 := []int{}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(v2, v6)
    v8 := GenS1{F0: v2, F1: 18446744073709551615}
    v9 := GenS1{F0: v2, F1: 1}
    v2 -= v2
    fmt.Println(v6)
    v10 := GenS1{F0: (v2 % ((v2) | 1)), F1: 569}
    var v11 int64 = 1
    var v12 int = -10
    if ((2147483648 != 352) && false) {
        switch ("a" + "foo bar") {
        case "\n\t":
            fmt.Println(v12)
        case "":
            if len(v7) > 0 {
                v7[0] = (v12 ^ (^(v12)))
            }
        case "foo bar":
            var v13 float64 = (-2.25 / 0.1)
            _ = v13
        }
        fmt.Println(v6, v11)
    }
    _ = len(v3)
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 uint64 = 1
        v16 := GenS1{F0: (v2 - (v2 & 2147483647)), F1: 831}
        v3[(("a" + "a") + "")] = 1
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
