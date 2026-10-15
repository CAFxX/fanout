package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int8; F1 int; F2 uint64 }

func main() {
    var v1 uint = 533
    v2 := make(map[uint32]int)
    v3 := []uint32{}
    var v4 int64 = -821
    v4 = 42
    var v5 int16 = 32767
    switch "foo bar" {
    case "a":
        var v6 int64 = (-(v4))
        _ = v6
    case "hello":
        if (-32768 == v5) {
            var v7 int = 4294967295
            for fi8 := 0; fi8 < 18; fi8++ {
                for fi9 := 0; fi9 < 13; fi9++ {
                    v10 := GenS1{F0: v1, F1: v1}
                    _ = fi9
                    _ = v10
                }
                v11 := v2[0]
                fmt.Println(v4)
                _ = fi8
                _ = v11
            }
            switch -2147483648 {
            case 128:
                v7 -= (-(v7))
            case 9223372036854775807:
                fmt.Println(v5, v7, v5)
            case -189:
                var v12 string = ""
                _ = v12
            }
            _ = v7
        } else {
            v1 += 2
            fmt.Println(v5, v1, v1)
        }
    default:
        v13 := v2[0]
        _ = v13
    }
    v14 := GenS1{F0: 0, F1: 96}
    v15 := GenS1{F0: v1, F1: v1}
    switch 9223372036854775807 {
    case 32768:
        v16 := GenS1{F0: v1, F1: (v1 << ((v1) & 63))}
        _ = v16
    case 127:
        v17 := GenS2{F0: 127, F1: 1, F2: 4294967295}
        _ = v17
    }
    var v18 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
