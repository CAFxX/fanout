package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }
func genFn0() uint16 {
    var v6 float32 = 0.1
    var v7 string = (("foo bar" + "a") + ("" + (((("a" + "\n\t") + "") + "foo bar") + ("hello" + "a"))))
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := []string{}
    var v2 uint8 = 1
    var v3 float64 = -2.25
    var v4 uint = 84
    v5 := make(map[uint32]bool)
    v4 -= (v4 | (v4 * 18446744073709551615))
    v8 := GenS1{F0: (499 < 312), F1: v4}
    var v9 uint = v4
    v2 -= (^(v2))
    var v10 uint8 = (v2 ^ (v2 >> ((2) & 63)))
    fmt.Println(v9, v10)
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint64 = 803
        for fi13 := 0; fi13 < 15; fi13++ {
            v10 = (v10 ^ (v2 ^ v2))
            v14 := GenS2{F0: (fi13 ^ (fi11 | fi13))}
            switch "foo bar" {
            case "foo bar":
                v3 -= float64(0.0)
            case "a":
                _ = len(v1)
            }
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    v15 := v5[65535]
    var v16 uint16 = 1
    var v17 int = -9223372036854775808
    var v18 uint64 = uint64(v4)
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := GenS3{F0: 4294967295}
        v21 := GenS4{F0: (v9 + v9)}
        fmt.Println(v9)
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 int32 = 1
    switch "\n\t" {
    case "":
        v23 := GenS3{F0: (v18 - (^(v18)))}
        _ = v23
    case "\n\t":
        v3 = (-2.25 + float64((0.1 * (-2.25 + -2.25))))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
