package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int8 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{"foo bar"}
    v3 := make(map[string]bool)
    v4 := []int{32768}
    var v5 int32 = -1
    v6 := make(map[int]string)
    v5 = v5
    switch "\n\t" {
    case "\n\t":
        switch ("a" + ("" + "a")) {
        case "a":
            _ = len(v2)
        }
    default:
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                if false {
                    v5 += v5
                    v5 = v5
                    v5 += (-(v5))
                }
                delete(v6, fi7)
                if (9223372036854775807 <= 1) {
                    v5 = (v5 % ((772) | 1))
                }
                _ = fi8
            }
            fmt.Println(fi7, fi7)
            var v9 uint64 = 1
            _ = fi7
            _ = v9
        }
    }
    switch "" {
    case "a":
        var v10 int8 = 57
        _ = v10
    case "":
        v11 := v3["\n\t"]
        _ = v11
    }
    var v12 int32 = v5
    v5 = (v5 - 1)
    if false {
        v3["foo bar"] = (0.0 != 1.5)
    }
    v6[32768] = ("foo bar" + "")
    var v13 float64 = 1.5
    var v14 string = ("" + "")
    v12 -= v5
    if false {
        v13 = -2.25
        fmt.Println(v13)
        v2 = append(v2, v14)
    }
    _ = len(v4)
    v15 := GenS1{F0: (!(false && (!true))), F1: 1, F2: 10}
    v5 -= v5
    v16 := GenS2{F0: -1}
    delete(v1, 2147483647)
    var v17 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
