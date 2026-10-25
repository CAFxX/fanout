package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int64 = 1
    v2 := make(map[string]int)
    var v3 int16 = 247
    if ((!(!(2147483647 > 928))) && false) {
        v1 -= (v1 * v1)
    }
    v4 := GenS1{F0: 869}
    v5 := []int{128, 128, -1}
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint = 65535
        _ = fi6
        _ = v7
    }
    v3 = (v3 | (-(v3)))
    switch 9223372036854775807 {
    case 798:
        var v8 float32 = 0.0
        _ = v8
    }
    fmt.Println(v3, v1)
    v1 = v1
    for fi9 := 0; fi9 < 16; fi9++ {
        switch "\n\t" {
        case "foo bar":
            v1 += v1
        case "":
            var v10 int16 = (v3 * (v3 * (v3 / ((v3) | 1))))
            _ = v10
        case "\n\t":
            fmt.Println(v1)
        }
        if true {
            v11 := GenS1{F0: 2}
            v2["foo bar"] = (fi9 * -2147483648)
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS1{F0: 0}
                var v14 int = (fi12 | (fi12 - 2147483647))
                fmt.Println(v1)
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
        switch fi9 {
        case 9223372036854775807:
            v1 = v1
        case 2147483647:
            var v15 float32 = 3.14159
            _ = v15
        default:
            v16 := GenS1{F0: 2}
            _ = v16
        }
        _ = fi9
    }
    v17 := GenS1{F0: 57}
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
