package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 bool; F2 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 219
    var v3 bool = true
    v4 := make(map[string]int)
    v5 := GenS1{F0: 1}
    if false {
        v4["foo bar"] = 9223372036854775807
        v3 = (((0 != 127) && ((1.5 > ((1.5 + -2.25) + 0.1)) && (1000 > -48))) || (v3 || v3))
    } else {
        var v6 uint8 = 178
        v7 := GenS1{F0: -1}
        v8 := []int{42}
        _ = v6
        _ = v7
        _ = v8
    }
    v3 = true
    if (!v3) {
        var v9 int = -9223372036854775808
        v10 := GenS1{F0: 1}
        _ = v10
        _ = v9
    }
    fmt.Println(v2, v2)
    for fi11 := 0; fi11 < 15; fi11++ {
        v2 = (v2 * 9223372036854775807)
        fmt.Println(v3, v3)
        _ = fi11
    }
    if v3 {
        v12 := GenS1{F0: -2147483648}
        for fi13 := 0; fi13 < 9; fi13++ {
            v3 = (1 <= 608)
            _ = fi13
        }
        fmt.Println(v2)
        _ = v12
    } else {
        fmt.Println(v3)
        _ = len(v1)
        var v14 uint8 = 2
        _ = v14
    }
    switch "" {
    case "foo bar":
        v15 := []int{0}
        _ = v15
    case "hello":
        v2 += v2
    }
    switch 9223372036854775807 {
    case -77:
        if (v3 || true) {
            for fi16 := 0; fi16 < 8; fi16++ {
                fmt.Println(fi16, fi16)
                v17 := []int{(fi16 + (fi16 << ((2) & 63))), 128}
                if (!v3) {
                    v1[2147483648] = true
                    var v18 int64 = 45
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
        } else {
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 int16 = 42
                _ = fi19
                _ = v20
            }
            v21 := []int{127, 357, 9223372036854775807, 32768}
            _ = v21
        }
    }
    v3 = false
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 int16 = -22
        _ = fi22
        _ = v23
    }
    fmt.Println(v3)
    if (-67 != -2147483648) {
        var v24 uint = 159
        var v25 bool = (("hello" + "") != "\n\t")
        _ = v24
        _ = v25
    }
    v26 := []int{9223372036854775807, 9223372036854775807, 128}
    for fi27 := 0; fi27 < 19; fi27++ {
        _ = len(v26)
        fmt.Println(fi27, v2, fi27)
        _ = fi27
    }
    v28 := v1[65535]
    v29 := GenS2{F0: -9223372036854775808, F1: (!v3), F2: ((1e10 - 0.1) <= 3.14159)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
