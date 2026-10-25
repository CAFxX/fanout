package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }

func main() {
    var v1 int = -728
    var v2 int = v1
    var v3 bool = false
    var v4 int64 = 53
    v1 -= (v2 / (((v1 >> ((1) & 63))) | 1))
    fmt.Println(v4, v1, v3)
    if (!((0.1 / (1e10 / 1.5)) == (1e10 * -2.25))) {
        if false {
            fmt.Println(v4)
        }
        var v5 int16 = -32768
        v3 = (2 == 90)
        _ = v5
    }
    if (-17 < 1) {
        v3 = ((!v3) && v3)
    }
    v6 := []int{127, (v2 >> ((206) & 63)), 128}
    v7 := GenS1{F0: "hello", F1: false, F2: 42}
    v8 := GenS1{F0: "a", F1: (v3 && (true && (false || v3))), F2: 42}
    v4 += 18
    fmt.Println(v4, v3)
    fmt.Println(v4)
    for fi9 := 0; fi9 < 8; fi9++ {
        v2 = (v2 & (v2 << ((4294967295) & 63)))
        _ = fi9
    }
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
