package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
func genFn0(p0 uint64, p1 uint64, p2 int8) int8 {
    fmt.Println(p0, p0, p2)
    p2 = 123
    v5 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}
func genFn1(p0 int64) int16 {
    switch -51 {
    case 2147483647:
        p0 = -933
    }
    v6 := []int{2147483647, 32768, 1, -2147483648}
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]int)
    var v3 uint8 = 0
    v4 := []int{9223372036854775807, 2147483647, 127, -699}
    for fi7 := 0; fi7 < 16; fi7++ {
        v3 = (v3 - v3)
        var v8 int32 = -912
        _ = fi7
        _ = v8
    }
    switch (("foo bar" + (("a" + "foo bar") + ("hello" + "a"))) + ("a" + "hello")) {
    case "\n\t":
        v4 = append(v4, -9223372036854775808)
    default:
        delete(v2, 2)
    }
    var v9 int8 = -1
    v10 := GenS1{F0: 4294967295, F1: -9223372036854775808}
    v1 = ((v3 * 42) >= v3)
    var v11 uint = 280
    fmt.Println(v1, v9)
    switch -9223372036854775808 {
    case 32768:
        v1 = (float32(3.14159) <= -2.25)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
