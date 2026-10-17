package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int8 }
func genFn0(p0 uint8, p1 int64, p2 int) int32 {
    fmt.Println(p2, p1)
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int64 = (p1 % ((p1) | 1))
            if (481 <= int32(fi7)) {
                var v9 float32 = 0.0
                _ = v9
            } else {
                var v10 uint = 4294967295
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return -478
}
func genFn1(p0 int) int16 {
    v11 := []int{p0, (p0 + (p0 / (((p0 % ((42) | 1))) | 1))), p0}
    p0 += 796
    p0 += 4294967295
    if len(v11) > 0 {
        v11[0] = p0
    }
    v12 := GenS1{F0: 9223372036854775807, F1: 450, F2: -1}
    _ = p0
    _ = v11
    _ = v12
    return 32767
}

func main() {
    var v1 int8 = 127
    v2 := []int{-8}
    var v3 int = -286
    v4 := make(map[uint32]uint32)
    var v5 uint = 804
    for fi13 := 0; fi13 < 1; fi13++ {
        switch ("" + (("hello" + "hello") + ("hello" + "foo bar"))) {
        case "foo bar":
            var v14 uint = (v5 + v5)
            _ = v14
        case "\n\t":
            var v15 int16 = genFn1(v3)
            _ = v15
        }
        var v16 uint16 = 65535
        _ = fi13
        _ = v16
    }
    var v17 float32 = 0.0
    v1 += (v1 - v1)
    var v18 uint32 = 1
    var v19 int8 = (v1 / (((v1 / ((v1) | 1))) | 1))
    var v20 uint16 = 1
    switch v3 {
    case -9223372036854775808:
        var v21 int8 = 29
        _ = v21
    }
    v17 -= v17
    v3 = v3
    for fi22 := 0; fi22 < 4; fi22++ {
        v18 += v18
        _ = fi22
    }
    var v23 uint32 = v18
    var v24 uint8 = 2
    var v25 uint16 = 32767
    v5 -= (v5 | uint(v25))
    for fi26 := 0; fi26 < 8; fi26++ {
        switch "foo bar" {
        case "a":
            var v27 uint16 = v25
            _ = v27
        case "foo bar":
            v25 = 32767
        case "":
            var v28 int32 = 1
            _ = v28
        }
        var v29 int64 = 1
        _ = fi26
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
