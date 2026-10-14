package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int64 = -113
    var v2 int16 = 32767
    v3 := make(map[int]bool)
    var v4 string = ("foo bar" + ("\n\t" + ("" + ("foo bar" + ("a" + "a")))))
    v1 = v1
    if true {
        if false {
            var v5 uint64 = 4294967295
            for fi6 := 0; fi6 < 7; fi6++ {
                var v7 float32 = (float32((0.1 / 1e10)) / 1.5)
                var v8 float64 = float64(v7)
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = v5
        } else {
            var v9 bool = true
            if v9 {
                if (!(!v9)) {
                    v4 = "\n\t"
                }
                if v9 {
                    v1 += v1
                    v10 := GenS1{F0: false, F1: v2, F2: -2147483648}
                    _ = v10
                } else {
                    v11 := GenS1{F0: (!true), F1: -12, F2: -3}
                    _ = v11
                }
            }
            _ = v9
        }
        v2 += (v2 / (((-(v2))) | 1))
        v2 = 0
    } else {
        if (("hello" + "") >= v4) {
            var v12 float64 = (0.0 - (1.5 + 1e10))
            v12 -= (v12 / v12)
            _ = v12
        }
        fmt.Println(v2, v4)
    }
    v13 := GenS2{F0: 0, F1: 835, F2: 2147483647}
    fmt.Println(v1, v1)
    switch (("foo bar" + ("\n\t" + (("a" + "hello") + v4))) + v4) {
    case "\n\t":
        var v14 float64 = 0.0
        _ = v14
    case "a":
        v3[9223372036854775807] = (false && (!false))
    case "foo bar":
        v15 := GenS1{F0: (42 <= 2), F1: (-(v2)), F2: -1}
        _ = v15
    default:
        fmt.Println(v1)
    }
    v16 := GenS3{F0: (v1 % ((1) | 1))}
    switch "foo bar" {
    case "\n\t":
        v1 = v1
    }
    v1 = v1
    if false {
        var v17 float32 = 0.0
        _ = v17
    }
    v18 := GenS1{F0: true, F1: 0, F2: -1}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
