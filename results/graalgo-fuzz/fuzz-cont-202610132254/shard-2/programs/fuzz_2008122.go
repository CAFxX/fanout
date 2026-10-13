package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint8 }

func main() {
    var v1 uint = 65535
    var v2 uint = (v1 % ((v1) | 1))
    var v3 int64 = 327
    v4 := []int{int(v1)}
    v5 := make(map[int]int)
    switch "\n\t" {
    case "a":
        delete(v5, -906)
    }
    fmt.Println(v2, v3, v3)
    var v6 int = 185
    switch ("" + "\n\t") {
    case "foo bar":
        var v7 string = "foo bar"
        _ = v7
    }
    var v8 uint8 = 1
    v6 = (^(v6))
    var v9 uint8 = v8
    v10 := GenS1{F0: -32768, F1: 1, F2: 127}
    v6 += v6
    switch ("a" + ((("" + "hello") + (("hello" + "\n\t") + "hello")) + "foo bar")) {
    case "a":
        v11 := GenS1{F0: 32767, F1: -435, F2: (v9 * 0)}
        _ = v11
    case "\n\t":
        switch "hello" {
        case "":
            v1 = (v1 / ((v1) | 1))
        default:
            v8 += 0
        }
    case "foo bar":
        v9 = 31
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 float64 = (((-2.25 / float64(1e10)) - (0.0 * 0.0)) + 1e10)
            if ((v9 + (v9 / ((v8) | 1))) > v8) {
                v9 = 240
                var v15 string = ("" + "hello")
                fmt.Println(v3, fi12, fi12)
                _ = v15
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            v2 += v2
            _ = fi13
            _ = v14
        }
        fmt.Println(v6)
        _ = fi12
    }
    switch -984 {
    case -758:
        v3 -= int64(v9)
    case -380:
        fmt.Println(v3, v3)
    default:
        v6 -= 127
    }
    fmt.Println(v8, v3)
    if len(v4) > 0 {
        v4[0] = 485
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
