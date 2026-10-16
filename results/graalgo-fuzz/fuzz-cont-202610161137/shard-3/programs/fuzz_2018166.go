package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32) uint32 {
    v5 := []int{32768, 743}
    var v6 int = 169
    v6 -= v6
    if false {
        v6 += -187
    } else {
        var v7 bool = (1 <= -1)
        _ = v7
    }
    p0 = p0
    _ = p0
    _ = v5
    _ = v6
    return 473
}
func genFn1(p0 uint, p1 uint64, p2 int16) uint {
    var v8 uint64 = p1
    switch 2147483647 {
    case 42:
        switch -9223372036854775808 {
        case 128:
            switch 32767 {
            case 4294967295:
                p2 = p2
            case 783:
                var v9 uint32 = 4294967295
                _ = v9
            case 32767:
                var v10 bool = false
                _ = v10
            }
        case 973:
            v11 := []int{-1, 9223372036854775807, 32767, 1}
            _ = v11
        case 822:
            p1 -= (v8 % ((v8) | 1))
        }
    case -891:
        if (4294967295 != 300) {
            var v12 uint64 = (p1 & p1)
            fmt.Println(v8)
            _ = v12
        } else {
            var v13 uint64 = v8
            _ = v13
        }
    case -1:
        switch "\n\t" {
        case "foo bar":
            var v14 int32 = -1
            _ = v14
        case "a":
            fmt.Println(v8)
        case "hello":
            var v15 int8 = -18
            _ = v15
        }
    }
    if (0.0 >= 0.0) {
        if (18446744073709551615 != 18446744073709551615) {
            var v16 uint64 = (p1 + (p1 % ((4294967295) | 1)))
            var v17 uint16 = 42
            if false {
                v18 := []int{-918, 1, 4294967295}
                _ = v18
            }
            _ = v16
            _ = v17
        }
        p0 = 4294967295
        fmt.Println(p2)
    } else {
        v19 := GenS1{F0: v8, F1: 127}
        switch 32768 {
        case 211:
            var v20 uint = p0
            _ = v20
        case -439:
            p0 += p0
        case 4294967295:
            var v21 uint = 996
            _ = v21
        default:
            p0 = 2
        }
        _ = v19
    }
    p2 = (p2 * (p2 - (p2 - (p2 & (-(p2))))))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (^(p0))
}

func main() {
    var v1 float64 = 0.1
    var v2 uint32 = 798
    v3 := make(map[int]bool)
    var v4 string = "\n\t"
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := GenS2{F0: 227}
        _ = fi22
        _ = v23
    }
    switch v4 {
    case "\n\t":
        if false {
            v4 = (v4 + v4)
        }
    case "":
        v24 := []int{127}
        _ = v24
    }
    v4 = "\n\t"
    v1 += (0.0 * 1e10)
    switch (v4 + v4) {
    case "hello":
        v4 = ("foo bar" + (v4 + "hello"))
    case "\n\t":
        v4 = "\n\t"
    }
    var v25 int32 = 42
    v1 = v1
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
