package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]uint32)
    v3 := []int{}
    var v4 int8 = -1
    v5 := make(map[int]int)
    v6 := GenS1{F0: 2147483647, F1: -1}
    switch ("a" + "foo bar") {
    case "\n\t":
        fmt.Println(v4, v1)
    case "foo bar":
        delete(v2, 2)
    }
    v4 = v4
    v2[4294967295] = 205
    v4 -= 0
    v4 -= (v4 * -128)
    if (((127 > 116) || false) && true) {
        for fi7 := 0; fi7 < 20; fi7++ {
            if (false && ((113 >= 42) && (164 > 579))) {
                v1 = v1
                var v8 float32 = ((3.14159 / ((1.5 + 1.5) * -2.25)) + float32((((((0.1 + 0.0) * (0.0 / ((1.5 * ((float64(1e10) * 3.14159) / ((0.0 * ((-2.25 / -2.25) * 1.5)) - 1e10))) - 0.1))) / 0.1) * 1.5) / 1.0)))
                _ = v8
            } else {
                v1 = 0
                fmt.Println(fi7)
            }
            if (!(!false)) {
                fmt.Println(v4)
                fmt.Println(v4, v1)
            }
            _ = fi7
        }
        _ = len(v2)
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 bool = true
            _ = len(v3)
            _ = fi9
            _ = v10
        }
    }
    v4 += (v4 << ((2) & 63))
    switch "hello" {
    case "\n\t":
        v1 += v1
    case "a":
        v4 += (v4 + 127)
    case "":
        v5[128] = -9223372036854775808
    }
    var v11 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
