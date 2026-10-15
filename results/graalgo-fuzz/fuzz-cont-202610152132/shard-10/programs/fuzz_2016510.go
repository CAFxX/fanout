package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int }
type GenS2 struct { F0 uint16; F1 int16; F2 int16 }
func genFn0(p0 int64, p1 uint32) uint64 {
    var v5 bool = false
    v6 := GenS1{F0: 2, F1: 0, F2: 42}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int64 = 9223372036854775807
        v9 := []int{(-(fi7)), (fi7 << ((2) & 63))}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 18446744073709551615, F1: 42, F2: 1}
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int, p1 int64) uint16 {
    if true {
        fmt.Println(p0, p1)
        v11 := GenS1{F0: 18446744073709551615, F1: 2147483647, F2: -2147483648}
        _ = v11
    }
    p1 -= p1
    p0 += (p0 ^ (p0 & (p0 - 0)))
    p0 = -1
    if (!(!true)) {
        v12 := GenS1{F0: 2, F1: 232, F2: (p0 * (-(p0)))}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 int32 = -2147483648
            fmt.Println(v14)
            _ = fi13
            _ = v14
        }
        var v15 string = ("hello" + (("\n\t" + "foo bar") + ("foo bar" + "\n\t")))
        fmt.Println(p0, p0)
        _ = v15
    }
    _ = p0
    _ = p1
    return 940
}

func main() {
    var v1 float64 = 1e10
    var v2 bool = (false && false)
    v3 := make(map[string]int)
    v4 := []string{"a", (("" + "a") + (("\n\t" + "a") + "a")), "foo bar", "\n\t"}
    fmt.Println(v2)
    var v16 uint32 = 654
    var v17 int16 = -32768
    switch "\n\t" {
    case "\n\t":
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 int = (fi18 >> ((65535) & 63))
            v20 := v3["foo bar"]
            v19 += (v19 % ((fi18) | 1))
            _ = fi18
            _ = v19
            _ = v20
        }
    case "foo bar":
        v21 := GenS1{F0: 20, F1: -279, F2: 292}
        _ = v21
    default:
        if true {
            _ = len(v3)
            var v22 bool = v2
            _ = v22
        }
    }
    var v23 float64 = v1
    v24 := GenS2{F0: genFn1(535, 0), F1: -32768, F2: -32768}
    var v25 uint64 = genFn0(9223372036854775807, 746)
    var v26 string = "a"
    v23 = float64(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
