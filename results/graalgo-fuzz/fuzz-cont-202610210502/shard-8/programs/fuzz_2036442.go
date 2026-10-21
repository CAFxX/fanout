package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 42
    v3 := []uint32{0, 924}
    if true {
        v2 = v2
        v2 += v2
    } else {
        v2 = v2
        var v4 float64 = (0.0 * 0.1)
        v2 = v2
        _ = v4
    }
    v5 := v1[2147483648]
    switch -355 {
    case -473:
        switch v5 {
        case "\n\t":
            fmt.Println(v2, v2)
        case "":
            if (-222 <= -2147483648) {
                _ = len(v3)
                var v6 int64 = 1
                _ = v6
            }
        case "a":
            fmt.Println(v5, v2)
        }
    }
    v7 := GenS1{F0: 42, F1: (-(v2))}
    for fi8 := 0; fi8 < 19; fi8++ {
        v2 += 747
        v5 = ("foo bar" + v5)
        _ = fi8
    }
    v5 = v5
    switch v2 {
    case 32767:
        v9 := v1[243]
        _ = v9
    case -444:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v10 string = "foo bar"
    v10 = "hello"
    fmt.Println(v2, v2, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v7
}
