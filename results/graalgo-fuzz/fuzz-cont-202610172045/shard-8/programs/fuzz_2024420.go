package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    v1 := []int{-1, -287}
    var v2 uint64 = 32
    v3 := make(map[int]int)
    var v4 int32 = -42
    var v5 uint = 65535
    var v6 uint = (v5 >> ((v5) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        fmt.Println(v6, v6, v2)
        delete(v3, -623)
        v2 += v2
    } else {
        v6 = 1
    }
    if true {
        v3[42] = 9223372036854775807
        v7 := GenS1{F0: "foo bar"}
        switch "foo bar" {
        case "":
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 uint16 = 867
                _ = fi8
                _ = v9
            }
        case "a":
            v10 := GenS1{F0: ("hello" + "\n\t")}
            _ = v10
        case "\n\t":
            fmt.Println(v6)
        }
        _ = v7
    }
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    v5 -= (v5 & v5)
    for fi11 := 0; fi11 < 1; fi11++ {
        if (-1 > 9223372036854775807) {
            v3[fi11] = (fi11 << ((v6) & 63))
            fmt.Println(v2)
            var v12 float32 = -2.25
            _ = v12
        } else {
            var v13 int64 = 1
            _ = v13
        }
        v14 := GenS2{F0: v4, F1: 65535}
        _ = fi11
        _ = v14
    }
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
