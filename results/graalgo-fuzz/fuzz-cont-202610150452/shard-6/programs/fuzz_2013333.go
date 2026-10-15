package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 bool; F1 int64 }
func genFn0(p0 int8, p1 uint32, p2 bool) uint8 {
    p0 = p0
    if p2 {
        var v5 bool = ((42 <= 255) || (0 == -892))
        _ = v5
    }
    var v6 float64 = (1.5 - 0.0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 127
}

func main() {
    var v1 string = ""
    var v2 uint64 = 1
    var v3 uint = 538
    var v4 uint32 = 279
    var v7 int = -782
    switch (-(v7)) {
    case -2147483648:
        v1 = (v1 + "")
    }
    var v8 float64 = 1.5
    v9 := GenS1{F0: (v2 | v2), F1: v3}
    v10 := []int{(v7 << ((0) & 63)), (-(v7))}
    v4 = v4
    _ = len(v10)
    v1 = v1
    for fi11 := 0; fi11 < 10; fi11++ {
        v1 = v1
        for fi12 := 0; fi12 < 8; fi12++ {
            if false {
                v2 = (^(v2))
            }
            v13 := GenS2{F0: (1e10 <= (((-2.25 - -2.25) / ((0.0 - 0.0) + -2.25)) * 1e10)), F1: 9223372036854775807}
            v3 -= uint(v2)
            _ = fi12
            _ = v13
        }
        v1 = (("hello" + "hello") + ("hello" + "foo bar"))
        _ = fi11
    }
    v3 = (^(v3))
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
