package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int64, p1 int32) int64 {
    v5 := []int{858}
    var v6 int = 2147483647
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p0 * p0)
}
func genFn1(p0 string) int32 {
    var v7 string = p0
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int = 32768
        fmt.Println(p0, v9, p0)
        p0 = ("hello" + ((p0 + "foo bar") + "\n\t"))
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: -108}
    v11 := GenS1{F0: -9223372036854775808}
    fmt.Println(v7)
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    return -508
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 5
    var v3 int = -2147483648
    var v4 uint8 = 42
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := GenS1{F0: v3}
        _ = fi12
        _ = v13
    }
    var v14 float32 = (((((0.0 * 0.1) - (0.1 / 1.5)) + 1e10) + (((1e10 - ((0.1 * 0.0) * -2.25)) + -2.25) / -2.25)) * 1e10)
    v2 = uint64(v4)
    v15 := GenS1{F0: v3}
    v1[v3] = 2147483647
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 uint8 = v4
        _ = fi16
        _ = v17
    }
    v14 = v14
    v3 = -199
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
