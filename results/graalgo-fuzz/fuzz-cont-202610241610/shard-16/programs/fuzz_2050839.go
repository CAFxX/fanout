package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint32 }
func genFn0(p0 int64, p1 bool, p2 int8) uint16 {
    switch "hello" {
    case "hello":
        v4 := []int{830, 950, -1, 128}
        _ = v4
    case "\n\t":
        v5 := []int{-1, 32768}
        _ = v5
    case "foo bar":
        switch ((("\n\t" + "") + ("a" + (("hello" + "foo bar") + "hello"))) + "hello") {
        case "foo bar":
            p2 += -1
        case "\n\t":
            p1 = p1
        case "hello":
            var v6 int16 = 1
            _ = v6
        default:
            var v7 uint = 18446744073709551615
            _ = v7
        }
    default:
        var v8 uint = 4294967295
        _ = v8
    }
    p2 = 42
    v9 := GenS1{F0: -1, F1: 101, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 360
}

func main() {
    v1 := []int{-220}
    var v2 float64 = 0.0
    var v3 int = 463
    var v10 uint = 642
    var v11 bool = (42 <= 42)
    v12 := GenS1{F0: -885, F1: 70, F2: 1}
    fmt.Println(v10, v2, v3)
    _ = len(v1)
    var v13 int64 = 42
    v11 = (!((1.5 < 1e10) || (-1 >= (v13 | v13))))
    v10 += v10
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v14 float64 = (v2 - v2)
    var v15 int32 = 2147483647
    if v11 {
        var v16 bool = v11
        var v17 uint16 = 2
        _ = v16
        _ = v17
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "" {
    case "\n\t":
        v18 := GenS1{F0: -870, F1: -1, F2: 545}
        _ = v18
    case "":
        if (42 < uint32(v10)) {
            v13 -= v13
            v10 = 65535
        }
    default:
        switch 2147483647 {
        case 128:
            fmt.Println(v15)
        case 32767:
            switch "foo bar" {
            case "foo bar":
                switch (v3 << ((v10) & 63)) {
                case 0:
                    var v19 float64 = v2
                    _ = v19
                case 42:
                    var v20 uint64 = 170
                    _ = v20
                default:
                    v21 := GenS1{F0: -473, F1: 0, F2: 0}
                    _ = v21
                }
            case "":
                if len(v1) > 0 {
                    v1[0] = (v3 | (v3 << ((42) & 63)))
                }
            case "hello":
                v15 -= v15
            }
        default:
            for fi22 := 0; fi22 < 2; fi22++ {
                fmt.Println(v3, v13, v15)
                switch "hello" {
                case "foo bar":
                    v13 = v13
                case "":
                    fmt.Println(v11)
                }
                _ = fi22
            }
        }
    }
    v13 += (v13 * v13)
    if (3.14159 != (0.1 / -2.25)) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v23 int8 = 1
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
