package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32; F1 bool }
func genFn0(p0 uint, p1 bool, p2 int8) int {
    v5 := []int{-2147483648, -9223372036854775808, -2147483648, -791}
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: 2, F1: 1}
        switch ("\n\t" + "foo bar") {
        case "":
            if len(v5) > 0 {
                v5[0] = 9223372036854775807
            }
        default:
            for fi8 := 0; fi8 < 5; fi8++ {
                fmt.Println(fi6, fi6)
                _ = len(v5)
                _ = fi8
            }
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}

func main() {
    v1 := []uint32{2, 2, 962, 42}
    v2 := []uint32{917, 660}
    var v3 int16 = -1
    var v4 int8 = -1
    v4 += v4
    var v9 bool = false
    switch "\n\t" {
    case "":
        switch 0 {
        case -2147483648:
            v9 = (2147483647 != 232)
        case 610:
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        }
    default:
        v10 := GenS1{F0: 65535, F1: 32768}
        _ = v10
    }
    fmt.Println(v3, v9, v9)
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int16 = 1000
            fmt.Println(fi11, v13)
            var v14 int32 = 904
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    var v15 uint = 340
    var v16 string = "hello"
    var v17 string = ("" + (v16 + v16))
    switch "" {
    case "\n\t":
        var v18 int64 = 886
        _ = v18
    default:
        v19 := GenS2{F0: 2147483647, F1: (512 == 409)}
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
