package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 string, p1 uint64, p2 uint) bool {
    v4 := []int{-664, 1, 478}
    var v5 uint32 = 871
    var v6 bool = true
    switch int(p2) {
    case -9223372036854775808:
        fmt.Println(v5)
    case 128:
        var v7 float64 = 0.0
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (2 <= 42)
}
func genFn1(p0 int64, p1 int32) string {
    v8 := GenS1{F0: "a"}
    p1 = (^(p1))
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint16 = 420
        _ = fi9
        _ = v10
    }
    if (0.0 >= 0.0) {
        p0 = p0
        var v11 float64 = (0.0 / 1.5)
        _ = v11
    } else {
        p0 += p0
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint32 = 2
            v14 := GenS1{F0: "foo bar"}
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v8
    return "hello"
}

func main() {
    var v1 uint = 526
    var v2 bool = false
    var v3 uint16 = 0
    for fi15 := 0; fi15 < 1; fi15++ {
        if (-128 >= 115) {
            v16 := GenS1{F0: "a"}
            _ = v16
        }
        _ = fi15
    }
    if v2 {
        v17 := []int{9223372036854775807, 4294967295, -2147483648}
        fmt.Println(v1)
        _ = v17
    }
    if v2 {
        v18 := []int{2147483647, 9223372036854775807}
        var v19 int8 = 42
        var v20 uint = 18446744073709551615
        _ = v18
        _ = v19
        _ = v20
    } else {
        var v21 int64 = 9223372036854775807
        _ = v21
    }
    v22 := []int{32767, -9223372036854775808, -1, 308}
    var v23 uint16 = (v3 | (^(v3)))
    v3 += 2
    _ = len(v22)
    _ = len(v22)
    fmt.Println(v1, v3)
    v24 := GenS1{F0: ("foo bar" + "a")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
