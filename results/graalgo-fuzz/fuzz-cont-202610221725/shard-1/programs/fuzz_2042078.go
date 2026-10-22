package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
func genFn0(p0 uint64) bool {
    var v7 int = -9223372036854775808
    v8 := []int{-2147483648}
    v7 = -1
    fmt.Println(p0, p0)
    _ = len(v8)
    _ = p0
    _ = v7
    _ = v8
    return (1000 > 943)
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]uint32)
    var v3 float64 = (0.0 * 0.0)
    v4 := []int{127, -4}
    v5 := []uint32{}
    var v6 int8 = -1
    if false {
        switch 9223372036854775807 {
        case -869:
            switch ("\n\t" + "") {
            case "foo bar":
                var v9 uint = 65535
                _ = v9
            default:
                var v10 int8 = v6
                _ = v10
            }
        case 128:
            var v11 uint = 255
            _ = v11
        }
        v6 -= (v6 * v6)
        v12 := GenS1{F0: (v1 >> ((4294967295) & 63)), F1: false}
        _ = v12
    }
    switch 1 {
    case -1:
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(v1, fi13, v1)
            v14 := GenS1{F0: (v1 | (v1 << ((0) & 63))), F1: false}
            _ = fi13
            _ = v14
        }
    case 9223372036854775807:
        var v15 uint = 18446744073709551615
        _ = v15
    case -2147483648:
        var v16 float32 = 1.5
        _ = v16
    default:
        if false {
            _ = len(v5)
            v3 = 3.14159
        }
    }
    var v17 uint16 = 1
    var v18 int = 391
    var v19 uint = 608
    var v20 bool = true
    fmt.Println(v18, v20)
    var v21 uint16 = 32767
    for fi22 := 0; fi22 < 16; fi22++ {
        switch ("" + (("" + "foo bar") + (("a" + "foo bar") + "foo bar"))) {
        case "foo bar":
            for fi23 := 0; fi23 < 5; fi23++ {
                if v20 {
                    v24 := v2["\n\t"]
                    if len(v5) > 0 {
                        v5[0] = (^(v24))
                    }
                    var v25 int = fi23
                    _ = v24
                    _ = v25
                }
                v20 = (v17 > v17)
                var v26 uint16 = (v17 % ((65535) | 1))
                _ = fi23
                _ = v26
            }
        case "":
            v27 := GenS1{F0: v1, F1: (true || (-1 > 42))}
            _ = v27
        case "a":
            var v28 int64 = -1
            _ = v28
        }
        _ = fi22
    }
    v1 -= v1
    v21 -= (v17 << ((v19) & 63))
    for fi29 := 0; fi29 < 2; fi29++ {
        fmt.Println(v1, v3)
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
