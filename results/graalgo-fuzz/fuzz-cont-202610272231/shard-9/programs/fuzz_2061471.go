package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64) uint32 {
    v4 := []int{4294967295}
    for fi5 := 0; fi5 < 9; fi5++ {
        fmt.Println(fi5, p0)
        _ = fi5
    }
    p0 += p0
    _ = p0
    _ = v4
    return 2147483648
}

func main() {
    var v1 uint32 = 647
    var v2 int = -2147483648
    var v3 string = "foo bar"
    v6 := GenS1{F0: -1}
    var v7 uint32 = 185
    for fi8 := 0; fi8 < 3; fi8++ {
        v2 = fi8
        var v9 uint8 = 1
        _ = fi8
        _ = v9
    }
    v10 := []int{-578, (v2 >> ((255) & 63)), (v2 & (v2 ^ 32767)), -155}
    v1 = v1
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS2{F0: v7}
        _ = fi11
        _ = v12
    }
    if len(v10) > 0 {
        v10[0] = v2
    }
    if false {
        var v13 int = 618
        for fi14 := 0; fi14 < 11; fi14++ {
            v13 = (v13 - (fi14 << ((0) & 63)))
            v1 = 4294967295
            var v15 int64 = -359
            _ = fi14
            _ = v15
        }
        v2 = 972
        _ = v13
    }
    switch 2147483647 {
    case -9223372036854775808:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    default:
        v16 := GenS2{F0: v1}
        _ = v16
    }
    if false {
        v2 = -9223372036854775808
    }
    var v17 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
