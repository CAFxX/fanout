package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32) bool {
    fmt.Println(p0)
    var v6 uint32 = p0
    _ = p0
    _ = v6
    return false
}

func main() {
    v1 := []string{"hello", "foo bar", "hello"}
    var v2 uint64 = 1
    v3 := []int{0}
    v4 := []string{("\n\t" + "a"), "\n\t"}
    var v5 int64 = 1
    v7 := GenS1{F0: 330, F1: 42, F2: 1}
    var v8 uint16 = 32767
    var v9 uint16 = v8
    v5 -= (-(v5))
    v9 = v8
    if true {
        if (!(!(((!(32767 == 593)) && (!true)) && ((^(v2)) <= v2)))) {
            var v10 uint8 = 2
            _ = v10
        } else {
            v9 = v8
            v9 += v9
            for fi11 := 0; fi11 < 7; fi11++ {
                v5 += v5
                _ = fi11
            }
        }
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS2{F0: (v9 + (^(v9)))}
        for fi14 := 0; fi14 < 16; fi14++ {
            v9 = v9
            _ = fi14
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            _ = len(v4)
            _ = fi15
        }
        _ = fi12
        _ = v13
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(v2, v9, fi16)
        _ = fi16
    }
    switch "\n\t" {
    case "foo bar":
        var v17 uint = 0
        _ = v17
    case "":
        v5 = (v5 ^ (v5 + v5))
    case "\n\t":
        fmt.Println(v9)
    }
    v2 += (v2 & (v2 - (v2 % ((2) | 1))))
    fmt.Println(v2, v2)
    fmt.Println(v9, v2, v5)
    var v18 int64 = 1
    var v19 int64 = 1
    v8 = (v8 << ((18446744073709551615) & 63))
    v2 = (v2 | v2)
    if (!true) {
        var v20 int16 = 1000
        v9 = (v8 - (v8 >> ((42) & 63)))
        _ = v20
    } else {
        var v21 int32 = -193
        _ = v21
    }
    if false {
        var v22 uint16 = 728
        v23 := GenS1{F0: uint(v9), F1: 934, F2: 2}
        if len(v3) > 0 {
            v3[0] = -335
        }
        _ = v22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
