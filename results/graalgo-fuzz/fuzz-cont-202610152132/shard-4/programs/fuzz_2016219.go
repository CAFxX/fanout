package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint8, p1 int16, p2 int16) uint {
    var v5 uint64 = 387
    fmt.Println(p0)
    p1 = 1
    var v6 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []int{2147483647, -217, 42, 128}
    var v2 float64 = (1e10 - 0.1)
    var v3 int64 = -1
    var v4 string = ""
    v7 := GenS1{F0: 116, F1: (("foo bar" + v4) + ((("" + (v4 + ((v4 + "a") + v4))) + "") + v4))}
    v4 = v4
    fmt.Println(v4, v2)
    v8 := GenS2{F0: 1, F1: 1}
    var v9 float64 = 3.14159
    v4 = v4
    if true {
        var v10 int32 = 2147483647
        v3 += (v3 * v3)
        if len(v1) > 0 {
            v1[0] = 32767
        }
        _ = v10
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v4 = ("" + (("a" + (("" + "foo bar") + v4)) + "a"))
    v9 = float64((0.1 * (1.5 * 0.0)))
    if false {
        fmt.Println(v2)
    }
    if len(v1) > 0 {
        v1[0] = 407
    }
    switch -172 {
    case 32768:
        if (-32768 >= 566) {
            var v11 uint32 = 2147483648
            var v12 int8 = -128
            var v13 uint32 = (v11 - v11)
            _ = v11
            _ = v12
            _ = v13
        }
    case 0:
        var v14 int32 = -2147483648
        _ = v14
    default:
        v2 += v2
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint = 1
        switch "\n\t" {
        case "foo bar":
            var v17 string = v4
            _ = v17
        case "a":
            for fi18 := 0; fi18 < 19; fi18++ {
                v19 := GenS1{F0: -132, F1: "a"}
                switch v4 {
                case "\n\t":
                    v16 = v16
                }
                var v20 int8 = 92
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        v1 = append(v1, fi15)
        _ = fi15
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
