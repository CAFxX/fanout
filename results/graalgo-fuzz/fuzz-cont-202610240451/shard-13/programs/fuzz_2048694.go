package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS5 struct { F0 int; F1 int32; F2 int16 }
type GenS6 struct { F0 int64; F1 int; F2 string }

func main() {
    var v1 int8 = 100
    var v2 int64 = -1
    v3 := make(map[int]uint32)
    var v4 string = "hello"
    v5 := []int{-233, -1}
    if false {
        v1 += v1
        switch "foo bar" {
        case "foo bar":
            v1 += v1
        }
        switch ("hello" + (v4 + v4)) {
        case "foo bar":
            if (v1 >= v1) {
                var v6 uint = 879
                v7 := GenS1{F0: v4, F1: 1}
                _ = v6
                _ = v7
            }
        default:
            for fi8 := 0; fi8 < 7; fi8++ {
                v3[fi8] = 936
                _ = fi8
            }
        }
    }
    v9 := GenS2{F0: 890}
    var v10 uint32 = 2147483647
    for fi11 := 0; fi11 < 9; fi11++ {
        if false {
            var v12 int64 = (v2 | (v2 * v2))
            var v13 uint = 816
            _ = v12
            _ = v13
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            fmt.Println(v1)
            v2 += v2
            _ = fi14
        }
        _ = fi11
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v1)
    v15 := GenS3{F0: 0}
    if false {
        var v16 uint32 = v10
        v17 := GenS4{F0: (v16 * (^(v10))), F1: (v2 - v2), F2: 992}
        fmt.Println(v1, v4)
        _ = v16
        _ = v17
    }
    switch 4294967295 {
    case -596:
        fmt.Println(v1, v4, v4)
    case 2147483647:
        switch v4 {
        case "\n\t":
            fmt.Println(v2, v2, v10)
        case "foo bar":
            _ = len(v3)
        }
    default:
        v18 := GenS5{F0: 46, F1: -173, F2: 32767}
        _ = v18
    }
    switch 127 {
    case -2147483648:
        var v19 uint = 49
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v21 := GenS6{F0: v2, F1: fi20, F2: v4}
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
