package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    var v1 int = -9223372036854775808
    var v2 string = (("hello" + "hello") + "foo bar")
    v3 := make(map[int]int)
    var v4 uint64 = 18446744073709551615
    switch (v1 - 274) {
    case 1:
        v5 := GenS1{F0: 9223372036854775807, F1: -925, F2: 18446744073709551615}
        _ = v5
    case -9223372036854775808:
        v6 := []int{(-(v1)), (v1 << ((42) & 63)), (v1 - v1), 128}
        _ = v6
    default:
        v1 -= v1
    }
    v7 := GenS1{F0: 0, F1: 9223372036854775807, F2: 2}
    v8 := []int{}
    for fi9 := 0; fi9 < 20; fi9++ {
        if false {
            fmt.Println(v1)
            v10 := GenS2{F0: ("" < (("" + v2) + (((v2 + "hello") + (v2 + "a")) + (v2 + v2)))), F1: (v4 + 42)}
            switch v1 {
            case -9223372036854775808:
                fmt.Println(v1, v4, fi9)
            case 459:
                v3[(-(v1))] = 128
            default:
                v2 = v2
            }
            _ = v10
        }
        var v11 string = v2
        fmt.Println(v2, fi9)
        _ = fi9
        _ = v11
    }
    var v12 int = (v1 & 898)
    v8 = append(v8, (-(v12)))
    switch ((("hello" + "\n\t") + "\n\t") + v2) {
    case "\n\t":
        if true {
            if len(v8) > 0 {
                v8[0] = (-(v1))
            }
            for fi13 := 0; fi13 < 8; fi13++ {
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 int16 = 42
                    v1 = int(v4)
                    fmt.Println(v12)
                    v2 = ("a" + v2)
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            _ = len(v3)
        } else {
            if (557 < 42) {
                v16 := GenS1{F0: 0, F1: -951, F2: 707}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                v3[v1] = 4294967295
                _ = v16
            }
        }
    }
    if false {
        if true {
            fmt.Println(v12)
            v4 = 11
        } else {
            v8 = append(v8, v12)
        }
    }
    var v17 uint64 = v4
    v17 += v17
    fmt.Println(v12, v17, v17)
    delete(v3, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
