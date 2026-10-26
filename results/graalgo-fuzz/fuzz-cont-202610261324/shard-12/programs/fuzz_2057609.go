package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
func genFn0(p0 uint8) uint64 {
    v6 := []int{32767, 359, -548}
    p0 = p0
    p0 -= p0
    _ = p0
    _ = v6
    return 42
}

func main() {
    var v1 bool = (-461 >= 0)
    var v2 uint16 = 42
    var v3 int = 1
    var v4 string = ((("foo bar" + (("" + ("hello" + "")) + "foo bar")) + "foo bar") + (("a" + ("a" + "hello")) + ""))
    var v5 float64 = 1.5
    if v1 {
        fmt.Println(v1, v4)
    }
    switch (v3 + v3) {
    case 9223372036854775807:
        v7 := []int{v3}
        _ = v7
    case 105:
        v8 := []int{0, v3}
        _ = v8
    }
    var v9 uint32 = 2147483648
    v2 -= v2
    fmt.Println(v1)
    v10 := GenS1{F0: 100, F1: 558}
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            v2 = v2
            v13 := []int{127, v3, v3, (fi11 * fi12)}
            fmt.Println(v2, fi12)
            fmt.Println(v3, v9, fi11)
            _ = len(v13)
            _ = fi12
            _ = v13
        }
        v9 += v9
        v3 += (^(fi11))
        _ = fi11
    }
    v14 := GenS1{F0: 42, F1: 2147483647}
    for fi15 := 0; fi15 < 20; fi15++ {
        fmt.Println(v3)
        _ = fi15
    }
    v16 := GenS1{F0: -128, F1: -415}
    fmt.Println(v3, v2, v1)
    v17 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
