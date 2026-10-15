package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 string, p1 uint8, p2 int8) int32 {
    switch ("a" + p0) {
    case "a":
        var v6 int64 = 982
        _ = v6
    case "foo bar":
        switch p0 {
        case "hello":
            var v7 uint16 = 1
            _ = v7
        }
    case "\n\t":
        v8 := GenS1{F0: 1, F1: -9223372036854775808, F2: (p2 ^ p2)}
        _ = v8
    default:
        if false {
            p2 -= p2
        } else {
            fmt.Println(p2)
            p1 = 98
        }
    }
    v9 := GenS2{F0: (p2 ^ 0)}
    if false {
        v10 := GenS1{F0: -32768, F1: 929, F2: (p2 + (p2 + 1))}
        var v11 float64 = 0.0
        _ = v10
        _ = v11
    } else {
        fmt.Println(p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    var v1 int32 = -810
    var v2 float64 = 0.0
    v3 := make(map[string]int)
    var v4 int = 842
    v5 := make(map[int]bool)
    switch (^(v4)) {
    case -1:
        v2 -= 1.5
    }
    var v12 float64 = 1.5
    v13 := []int{738, 42, (v4 % ((v4) | 1)), v4}
    for fi14 := 0; fi14 < 17; fi14++ {
        delete(v5, v4)
        _ = fi14
    }
    if (0 > 358) {
        v1 -= (v1 >> ((0) & 63))
        switch 42 {
        case 9223372036854775807:
            v15 := v3[""]
            _ = v15
        case 128:
            var v16 float32 = (-2.25 * (-2.25 / 1.5))
            _ = v16
        case -76:
            switch "a" {
            case "":
                var v17 uint8 = 42
                _ = v17
            case "hello":
                for fi18 := 0; fi18 < 17; fi18++ {
                    v19 := GenS1{F0: 527, F1: -1, F2: -10}
                    _ = fi18
                    _ = v19
                }
            case "foo bar":
                fmt.Println(v4, v2)
            }
        }
    }
    v13 = append(v13, (v4 >> ((2) & 63)))
    v5[4294967295] = false
    v2 -= 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
