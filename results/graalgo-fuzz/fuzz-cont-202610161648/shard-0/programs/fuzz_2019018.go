package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint16 = 32767
    v3 := []uint32{2147483647}
    var v4 uint = (v1 - 18446744073709551615)
    v5 := make(map[int]bool)
    var v6 uint = v4
    v4 += uint(v2)
    var v7 uint64 = 4294967295
    switch ("\n\t" + "foo bar") {
    case "":
        v3 = append(v3, 899)
    case "foo bar":
        v7 += v7
    default:
        v1 = (v4 - v4)
    }
    if (1 <= -16) {
        var v8 float32 = ((0.1 / 3.14159) + 3.14159)
        _ = v8
    } else {
        _ = len(v3)
        for fi9 := 0; fi9 < 7; fi9++ {
            v7 += v7
            _ = fi9
        }
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            v1 -= (v6 + 255)
        }
    }
    if false {
        v6 = v4
    } else {
        fmt.Println(v4, v1, v6)
        if (!true) {
            v1 = v1
            switch 42 {
            case 4294967295:
                var v10 uint = v6
                _ = v10
            case -9223372036854775808:
                var v11 uint = (^(v6))
                _ = v11
            }
            v1 += v4
        }
    }
    v12 := v5[878]
    v1 -= v4
    v13 := GenS1{F0: 874, F1: v2}
    var v14 uint64 = 2
    fmt.Println(v7, v7)
    for fi15 := 0; fi15 < 5; fi15++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
