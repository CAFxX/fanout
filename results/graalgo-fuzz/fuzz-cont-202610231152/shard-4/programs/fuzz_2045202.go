package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int8 }
func genFn0(p0 uint8, p1 uint8) int {
    v7 := []int{2147483647, -9223372036854775808, -666}
    v8 := GenS1{F0: 628, F1: 1, F2: 104}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 int, p1 int32) uint32 {
    p0 += p0
    var v9 float64 = 1e10
    switch (("foo bar" + "a") + "") {
    case "foo bar":
        if false {
            var v10 int = (p0 + p0)
            p0 = (-(v10))
            _ = v10
        }
    case "hello":
        var v11 int = (-(p0))
        _ = v11
    }
    if (-128 >= int8(p1)) {
        fmt.Println(p1, v9)
        v12 := GenS1{F0: p0, F1: p0, F2: 100}
        p0 += p0
        _ = v12
    } else {
        var v13 bool = true
        _ = v13
    }
    switch "\n\t" {
    case "":
        var v14 uint = 862
        _ = v14
    case "hello":
        var v15 int32 = 0
        _ = v15
    case "foo bar":
        v9 = -2.25
    default:
        p0 = genFn0(1, 148)
    }
    _ = p0
    _ = p1
    _ = v9
    return 2147483648
}

func main() {
    var v1 uint8 = 255
    var v2 uint8 = (v1 ^ v1)
    var v3 uint8 = v2
    var v4 bool = false
    v5 := []int{321, 622, int(v2)}
    var v6 int8 = -128
    if v4 {
        if v4 {
            v1 = v2
            var v16 uint8 = v2
            _ = v16
        }
        if (42 >= 17) {
            v2 += v3
            var v17 int64 = -9223372036854775808
            _ = v17
        }
    } else {
        fmt.Println(v3)
    }
    v1 = v2
    v4 = false
    fmt.Println(v4, v3)
    v5 = append(v5, 2147483647)
    v4 = (53 == 42)
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := GenS1{F0: (fi19 >> ((0) & 63)), F1: 0, F2: (v6 ^ v6)}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    var v21 uint16 = 0
    for fi22 := 0; fi22 < 12; fi22++ {
        v2 = v1
        var v23 float32 = 1e10
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: 233, F1: 0, F2: v6}
    _ = len(v5)
    v25 := GenS1{F0: -646, F1: -76, F2: (v6 % (((v6 - (v6 * (v6 + 42)))) | 1))}
    var v26 bool = (4294967295 < 2147483648)
    var v27 int8 = -1
    v21 += v21
    v21 += v21
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
