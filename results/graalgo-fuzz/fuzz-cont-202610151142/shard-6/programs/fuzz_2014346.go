package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []int{}
    var v2 uint8 = 42
    var v3 float64 = 1e10
    v4 := GenS1{F0: 42}
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(v2, v3)
        _ = fi5
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        v2 = 128
        v7 := GenS1{F0: 150}
        _ = fi6
        _ = v7
    }
    if (false || true) {
        v2 += (v2 % (((^(v2))) | 1))
        v1 = append(v1, 503)
    }
    fmt.Println(v2)
    v2 -= (^(v2))
    switch -942 {
    case -9223372036854775808:
        var v8 float32 = 0.0
        _ = v8
    default:
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 int8 = int8(v2)
            var v11 string = ((("foo bar" + (("a" + "hello") + "")) + "hello") + "foo bar")
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    v2 -= v2
    var v12 uint32 = 2147483647
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = (^(v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
