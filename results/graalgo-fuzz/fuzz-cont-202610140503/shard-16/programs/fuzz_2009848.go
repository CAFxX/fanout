package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 string }
func genFn0(p0 int, p1 bool) int16 {
    p1 = (!p1)
    var v4 float32 = 0.1
    _ = p0
    _ = p1
    _ = v4
    return -1
}
func genFn1(p0 int16, p1 uint64) uint64 {
    if true {
        p0 = (p0 + genFn0(-239, (p1 > p1)))
        var v5 bool = (!(1 == 2))
        _ = v5
    }
    p1 -= p1
    var v6 int16 = (p0 - p0)
    _ = p0
    _ = p1
    _ = v6
    return (p1 | p1)
}

func main() {
    var v1 int16 = -1
    var v2 uint8 = 2
    var v3 uint32 = 184
    fmt.Println(v3)
    v1 = v1
    v7 := []int{}
    switch "" {
    case "foo bar":
        var v8 float32 = 0.0
        _ = v8
    default:
        v9 := GenS1{F0: 42, F1: 2, F2: ("a" + ("\n\t" + "foo bar"))}
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS1{F0: 0, F1: 18446744073709551615, F2: "a"}
        _ = fi10
        _ = v11
    }
    if (true && false) {
        v12 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: ("a" + "hello")}
        _ = v12
    }
    if len(v7) > 0 {
        v7[0] = 42
    }
    var v13 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v7
}
