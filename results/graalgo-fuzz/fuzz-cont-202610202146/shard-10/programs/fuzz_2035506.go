package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int64 }
type GenS2 struct { F0 bool; F1 bool }

func main() {
    v1 := []int{-184}
    v2 := make(map[int]string)
    var v3 int8 = 127
    v4 := make(map[string]bool)
    v5 := GenS1{F0: 128, F1: -96, F2: 9223372036854775807}
    fmt.Println(v3)
    v3 = (^(v3))
    v3 += (v3 + (-(v3)))
    switch 0 {
    case -9223372036854775808:
        var v6 int = 32767
        _ = v6
    case -566:
        fmt.Println(v3)
    default:
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := GenS2{F0: false, F1: false}
            v9 := GenS1{F0: 149, F1: 236, F2: -758}
            if false {
                v3 -= v3
                v3 -= v3
            } else {
                var v10 string = (((("hello" + "") + "hello") + ("a" + ("foo bar" + (("foo bar" + (("a" + "\n\t") + "a")) + "foo bar")))) + ("hello" + ("a" + "hello")))
                _ = v10
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v11 int = 127
    if len(v1) > 0 {
        v1[0] = v11
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
