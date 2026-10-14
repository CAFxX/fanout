package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 uint32 = 0
    var v2 int16 = 42
    var v3 string = "a"
    v4 := make(map[string]int)
    v5 := v4["foo bar"]
    for fi6 := 0; fi6 < 11; fi6++ {
        v1 -= v1
        v1 = 183
        _ = fi6
    }
    v1 += 2147483647
    v7 := GenS1{F0: 0}
    var v8 float32 = ((0.1 - 0.0) + -2.25)
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS2{F0: 2, F1: 4294967295}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        if ((127 <= -128) || (18446744073709551615 > 874)) {
            delete(v4, "foo bar")
            var v12 uint32 = 928
            fmt.Println(v12)
            _ = v12
        } else {
            switch 9223372036854775807 {
            case -1:
                v8 -= v8
            }
            v13 := []int{}
            var v14 float32 = v8
            _ = v13
            _ = v14
        }
        var v15 uint16 = 0
        v16 := GenS1{F0: -128}
        _ = fi11
        _ = v15
        _ = v16
    }
    v8 = ((v8 / 0.1) - v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
