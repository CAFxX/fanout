package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 bool }
func genFn0(p0 int8, p1 uint32) uint64 {
    fmt.Println(p0)
    p0 += p0
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 2
}

func main() {
    v1 := []string{("a" + "hello")}
    var v2 int32 = -2147483648
    v3 := make(map[int]int)
    var v4 int8 = 0
    var v5 uint8 = 1
    if (((("foo bar" + "foo bar") + "hello") + "") < "a") {
        var v6 float64 = 1e10
        for fi7 := 0; fi7 < 13; fi7++ {
            delete(v3, (fi7 & (fi7 % ((fi7) | 1))))
            v2 += (v2 | (v2 ^ v2))
            var v8 uint64 = 2
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    var v9 bool = false
    var v10 int8 = (v4 | int8(v5))
    if v9 {
        v2 -= -384
    }
    switch (("foo bar" + "\n\t") + "") {
    case "":
        var v11 int64 = 0
        _ = v11
    case "foo bar":
        for fi12 := 0; fi12 < 12; fi12++ {
            _ = len(v1)
            for fi13 := 0; fi13 < 6; fi13++ {
                v5 -= (v5 >> ((0) & 63))
                var v14 int64 = -9223372036854775808
                fmt.Println(v2, fi12, v5)
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    case "a":
        v10 -= -36
    }
    v15 := GenS1{F0: 881, F1: v4}
    var v16 uint32 = 2147483648
    if v9 {
        var v17 int = 4294967295
        _ = v17
    } else {
        v16 = v16
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v16, v9)
    case "\n\t":
        v18 := GenS1{F0: 682, F1: v4}
        _ = v18
    case "hello":
        v10 -= -98
    }
    v19 := GenS2{F0: (9223372036854775807 != -1)}
    v5 += v5
    fmt.Println(v9)
    var v20 int32 = (v2 << ((466) & 63))
    switch "foo bar" {
    case "hello":
        if v9 {
            switch 1 {
            case 128:
                var v21 int8 = 100
                _ = v21
            }
        } else {
            fmt.Println(v20, v20)
            v5 -= (v5 & (v5 ^ v5))
        }
    default:
        v10 = v4
    }
    v1 = append(v1, ("a" + (("a" + (("foo bar" + "a") + "")) + "hello")))
    var v22 int32 = 2147483647
    fmt.Println(v22, v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
