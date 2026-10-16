package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }

func main() {
    var v1 int64 = -501
    var v2 int = -9223372036854775808
    var v3 string = (("a" + (("foo bar" + "hello") + "a")) + "")
    v4 := GenS1{F0: 96, F1: -362}
    fmt.Println(v3)
    fmt.Println(v2)
    v5 := []int{(-(v2)), v2, (v2 ^ (v2 ^ v2))}
    v6 := GenS1{F0: 0, F1: 0}
    v1 = v1
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    v1 += (-(v1))
    fmt.Println(v2, v2, v1)
    var v7 uint16 = 549
    fmt.Println(v3, v2, v2)
    if true {
        if (v1 <= 773) {
            fmt.Println(v3, v3)
            fmt.Println(v7, v7)
        } else {
            v8 := GenS1{F0: 42, F1: -322}
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int8 = -95
        v10 -= -29
        _ = fi9
        _ = v10
    }
    v2 += (v2 & (v2 + (^(v2))))
    v11 := GenS1{F0: 65535, F1: 0}
    for fi12 := 0; fi12 < 17; fi12++ {
        v1 = v1
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
