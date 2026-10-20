package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0() int32 {
    v5 := []int{-967, 32767}
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 uint32 = 432
            switch "a" {
            case "hello":
                v9 := GenS1{F0: 32768, F1: 1, F2: 0}
                _ = v9
            case "":
                v5 = append(v5, (fi6 % ((fi6) | 1)))
            default:
                if false {
                    var v10 int = fi7
                    _ = v10
                }
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    if true {
        var v11 float64 = float64(-2.25)
        _ = v11
    }
    _ = v5
    return 2147483647
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 466
    var v3 uint16 = 313
    v4 := make(map[string]int)
    switch "foo bar" {
    case "hello":
        _ = len(v4)
    default:
        if (-1 != 2147483647) {
            v4["a"] = 0
        }
    }
    var v12 uint = 184
    v12 = v12
    switch "foo bar" {
    case "foo bar":
        var v13 bool = (true && (0 > 1000))
        _ = v13
    case "\n\t":
        var v14 uint16 = 2
        _ = v14
    default:
        v15 := GenS2{F0: -402, F1: 690}
        _ = v15
    }
    fmt.Println(v12)
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 float64 = (3.14159 - 1e10)
        var v18 int = (fi16 & fi16)
        v2 = 2147483647
        _ = fi16
        _ = v17
        _ = v18
    }
    switch ("hello" + ("\n\t" + ((("hello" + "hello") + "") + "\n\t"))) {
    case "foo bar":
        var v19 uint = (v12 - v12)
        _ = v19
    case "hello":
        v3 -= v3
    case "a":
        for fi20 := 0; fi20 < 12; fi20++ {
            fmt.Println(v3, v3)
            _ = fi20
        }
    }
    _ = len(v4)
    var v21 uint32 = (^(v2))
    _ = len(v4)
    fmt.Println(v2, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
