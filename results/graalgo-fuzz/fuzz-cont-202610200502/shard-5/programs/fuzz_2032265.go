package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint }
type GenS2 struct { F0 uint; F1 int32; F2 uint32 }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]string)
    v3 := []int{}
    var v4 int = 9223372036854775807
    if ((-2147483648 == -2147483648) && true) {
        switch "foo bar" {
        case "a":
            if (2147483648 <= 42) {
                v1 += v1
            }
        case "foo bar":
            fmt.Println(v4)
        }
        var v5 uint32 = 2
        _ = v5
    }
    v6 := GenS1{F0: (^(v4)), F1: 787, F2: 2}
    if (false && true) {
        v4 -= v4
    } else {
        v1 += (v1 + v1)
    }
    v1 = v1
    v7 := GenS2{F0: 333, F1: -1, F2: 4294967295}
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS2{F0: 0, F1: -1, F2: 1}
        _ = fi8
        _ = v9
    }
    if true {
        switch "hello" {
        case "\n\t":
            var v10 float32 = 0.1
            _ = v10
        default:
            v4 = -305
        }
    } else {
        var v11 float64 = 0.1
        v11 += v11
        _ = v11
    }
    if len(v3) > 0 {
        v3[0] = (v4 + v4)
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        if false {
            v4 += 9223372036854775807
        }
        _ = fi12
    }
    var v13 string = ("a" + (("a" + "foo bar") + ("a" + "foo bar")))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
