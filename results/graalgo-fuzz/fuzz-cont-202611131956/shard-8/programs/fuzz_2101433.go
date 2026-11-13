package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]bool)
    var v3 int16 = 1000
    var v4 bool = false
    v5 := []int{32768}
    v6 := GenS1{F0: "\n\t", F1: 2147483648}
    fmt.Println(v1, v1)
    v1 += -924
    if ((v1 + 223) < v1) {
        var v7 int8 = 19
        v4 = v4
        _ = v7
    } else {
        fmt.Println(v4)
        switch "hello" {
        case "foo bar":
            v8 := GenS2{F0: 1}
            _ = v8
        }
        var v9 bool = (v4 && (v4 || (v3 > v3)))
        _ = v9
    }
    var v10 uint = 2
    v11 := GenS2{F0: 38}
    fmt.Println(v10)
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 string = "foo bar"
        _ = fi12
        _ = v13
    }
    if (!v4) {
        var v14 string = "foo bar"
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        fmt.Println(v4)
        for fi16 := 0; fi16 < 11; fi16++ {
            v4 = (false && false)
            _ = fi16
        }
        _ = fi15
    }
    v2[131] = v4
    v10 += v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
