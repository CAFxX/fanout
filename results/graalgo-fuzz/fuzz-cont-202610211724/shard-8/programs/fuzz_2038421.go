package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int = 105
    var v2 uint32 = 2
    v3 := []int{(v1 << ((165) & 63)), -977, v1}
    v4 := GenS1{F0: 1, F1: true}
    var v5 uint32 = (v2 | v2)
    var v6 int32 = 2147483647
    v7 := GenS1{F0: 42, F1: (false && ((^(v2)) <= (^(v5))))}
    var v8 uint = 1
    v8 = v8
    if true {
        switch v1 {
        case 356:
            v9 := GenS2{F0: 146}
            _ = v9
        case -9223372036854775808:
            if true {
                var v10 uint64 = uint64(v5)
                _ = v10
            } else {
                var v11 int = (v1 << (((v8 | v8)) & 63))
                _ = v11
            }
        case -2147483648:
            var v12 int16 = -595
            _ = v12
        }
    } else {
        var v13 string = ("a" + "\n\t")
        _ = v13
    }
    v6 = (v6 ^ v6)
    switch (v1 | v1) {
    case 99:
        if len(v3) > 0 {
            v3[0] = v1
        }
    case -1:
        v14 := GenS2{F0: 0}
        _ = v14
    default:
        _ = len(v3)
    }
    if true {
        var v15 int = (v1 | 484)
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := GenS1{F0: -448, F1: (("" + "") < "\n\t")}
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        fmt.Println(v8, v6)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v18 int16 = -589
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch ("foo bar" + ("hello" + "a")) {
    case "\n\t":
        var v19 uint32 = 734
        _ = v19
    }
    var v20 int32 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
