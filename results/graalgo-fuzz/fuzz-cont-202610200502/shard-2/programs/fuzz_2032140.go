package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int, p1 int) int8 {
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := []int{128}
            _ = fi8
            _ = v9
        }
        var v10 uint = 18446744073709551615
        var v11 uint8 = 127
        _ = fi7
        _ = v10
        _ = v11
    }
    var v12 int16 = -166
    v13 := GenS1{F0: 32768}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 1
}
func genFn1(p0 bool, p1 int32, p2 uint32) int {
    var v14 float64 = -2.25
    v14 = (1e10 + ((1.5 / v14) + v14))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return -9223372036854775808
}

func main() {
    v1 := []string{"foo bar", "hello", ("foo bar" + ((("" + "\n\t") + "hello") + "\n\t"))}
    v2 := make(map[uint32]int)
    var v3 bool = (!false)
    var v4 uint64 = 0
    var v5 float32 = -2.25
    var v6 uint = 733
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS1{F0: fi15}
        var v17 int16 = -32768
        _ = fi15
        _ = v16
        _ = v17
    }
    if ((false || (v3 || (false || false))) || v3) {
        var v18 int64 = -1
        v3 = v3
        _ = v18
    }
    var v19 uint = uint(v4)
    for fi20 := 0; fi20 < 7; fi20++ {
        v6 = (v19 * (^(v19)))
        _ = fi20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        if v3 {
            var v22 uint8 = 0
            _ = v22
        }
        fmt.Println(v6, v6, v3)
        fmt.Println(v3)
        _ = fi21
    }
    v23 := GenS1{F0: -206}
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 int = (fi24 & -9223372036854775808)
        _ = fi24
        _ = v25
    }
    v19 -= v19
    v6 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
