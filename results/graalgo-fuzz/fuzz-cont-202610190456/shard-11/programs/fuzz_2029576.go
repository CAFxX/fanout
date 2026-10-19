package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int16 }
func genFn0(p0 uint16) uint64 {
    v4 := GenS1{F0: -1, F1: 992, F2: 0}
    if (!false) {
        v5 := []int{4294967295}
        _ = v5
    } else {
        v6 := []int{2147483647, -534, 2147483647}
        _ = v6
    }
    v7 := []int{661, 1, 127}
    p0 -= 65535
    p0 -= p0
    _ = p0
    _ = v4
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int, p1 uint32, p2 int) int8 {
    var v8 bool = (0 < 1)
    v8 = (!((0.1 >= (0.0 * (0.1 - 1e10))) && (2 > 0)))
    if v8 {
        p2 = (p0 << ((0) & 63))
        v9 := []int{9223372036854775807, p0, 4294967295, -1}
        _ = v9
    }
    var v10 bool = v8
    switch "hello" {
    case "foo bar":
        var v11 int8 = -1
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 int64 = -9223372036854775808
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    case "":
        if (!false) {
            switch "hello" {
            case "hello":
                p1 = 2147483647
            }
        } else {
            var v15 int64 = 1
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    return -128
}

func main() {
    var v1 uint = 953
    var v2 uint32 = 767
    var v3 int32 = 0
    for fi16 := 0; fi16 < 10; fi16++ {
        fmt.Println(fi16, fi16)
        fmt.Println(v1, fi16)
        var v17 int8 = 42
        _ = fi16
        _ = v17
    }
    switch "a" {
    case "":
        v2 = (v2 | 2)
    }
    switch "" {
    case "":
        var v18 float64 = 1.5
        _ = v18
    }
    v19 := []int{32768, 32767, 0}
    if false {
        var v20 int32 = v3
        var v21 uint8 = 133
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 6; fi22++ {
            if len(v19) > 0 {
                v19[0] = fi22
            }
            for fi23 := 0; fi23 < 14; fi23++ {
                if (-29 < 42) {
                    v1 = (v1 * v1)
                }
                if len(v19) > 0 {
                    v19[0] = fi22
                }
                var v24 int = fi23
                if len(v19) > 0 {
                    v19[0] = (-(v24))
                }
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        fmt.Println(v3)
        v3 -= v3
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v1, v1)
    case "hello":
        var v25 int64 = -932
        _ = v25
    }
    var v26 bool = false
    v27 := GenS1{F0: -125, F1: 42, F2: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
