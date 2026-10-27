package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint64 }
func genFn0(p0 uint16) uint16 {
    fmt.Println(p0)
    var v7 uint64 = 0
    if true {
        v7 += 986
        v8 := GenS1{F0: 594, F1: v7, F2: v7}
        v9 := []int{42, 972, 1, -2147483648}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            fmt.Println(v7)
            _ = fi11
        }
        _ = fi10
    }
    _ = p0
    _ = v7
    return p0
}

func main() {
    var v1 string = "foo bar"
    var v2 uint16 = 736
    var v3 uint32 = 2147483647
    var v4 int = 127
    var v5 int32 = 42
    v6 := make(map[uint32]uint32)
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint = 1
        v14 := []int{(^(v4)), fi12}
        _ = len(v6)
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 0, F1: 4294967295, F2: 0}
    for fi16 := 0; fi16 < 9; fi16++ {
        v1 = "a"
        var v17 float32 = float32(float64(float32((3.14159 / (-2.25 / 1e10)))))
        _ = fi16
        _ = v17
    }
    var v18 float64 = ((0.1 * 1.5) + 0.0)
    v19 := GenS1{F0: 852, F1: 1, F2: 847}
    v1 = v1
    for fi20 := 0; fi20 < 19; fi20++ {
        fmt.Println(v5)
        _ = fi20
    }
    switch "" {
    case "foo bar":
        v6[v3] = v3
    }
    v5 -= -2147483648
    v5 = (-(v5))
    if (32767 >= (v4 ^ (v4 & v4))) {
        v21 := GenS1{F0: 0, F1: 2, F2: 971}
        _ = v21
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
