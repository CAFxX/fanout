package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
func genFn0(p0 uint32, p1 int8) int16 {
    p0 -= (p0 / ((p0) | 1))
    var v6 int8 = p1
    v7 := GenS1{F0: 4294967295}
    v8 := GenS2{F0: -699}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1000
}

func main() {
    v1 := []string{(("foo bar" + (("\n\t" + "\n\t") + "\n\t")) + ""), "hello", "foo bar"}
    var v2 bool = false
    v3 := []int{42}
    var v4 uint64 = 18446744073709551615
    v5 := []int{-2147483648, 42, 32767}
    v9 := GenS2{F0: 128}
    v2 = v2
    v10 := GenS1{F0: 1}
    v4 = v4
    switch 32767 {
    case 32767:
        fmt.Println(v2, v4)
    default:
        v2 = v2
    }
    fmt.Println(v4)
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 int8 = -128
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 uint64 = (v4 | 218)
        var v15 int64 = 471
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = len(v5)
    v2 = v2
    v4 -= 4294967295
    v16 := GenS2{F0: 585}
    if true {
        for fi17 := 0; fi17 < 8; fi17++ {
            _ = len(v5)
            var v18 bool = false
            _ = fi17
            _ = v18
        }
        v1 = append(v1, "\n\t")
        v4 -= v4
    } else {
        v4 -= v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
