package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int16, p1 uint8, p2 uint32) int32 {
    if (p0 >= 955) {
        var v6 uint64 = 42
        p0 += p0
        _ = v6
    }
    v7 := []int{627, int(p0), -1, 32768}
    v8 := GenS1{F0: 421}
    p0 = (p0 + -919)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []int{}
    v2 := []int{-683, -276}
    var v3 uint32 = 835
    v4 := []uint32{}
    v5 := []string{""}
    v3 = v3
    var v9 uint16 = 895
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS2{F0: 0}
        v12 := GenS1{F0: 999}
        v9 = (v9 << ((0) & 63))
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        var v13 uint16 = 450
        fmt.Println(v3)
        fmt.Println(v9)
        _ = v13
    }
    fmt.Println(v3, v3)
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := GenS2{F0: 1}
        _ = fi14
        _ = v15
    }
    switch 32768 {
    case 1:
        v3 = (v3 % (((^(v3))) | 1))
    case -345:
        if false {
            var v16 bool = (!false)
            _ = v16
        }
    case 429:
        v1 = append(v1, 32767)
    }
    switch ("\n\t" + ("hello" + ("\n\t" + "a"))) {
    case "":
        v3 = (^(v3))
    default:
        var v17 int8 = 86
        _ = v17
    }
    if false {
        if ((v3 % ((478) | 1)) >= (v3 * (v3 * v3))) {
            v9 += v9
        }
        v18 := GenS2{F0: 403}
        _ = v18
    }
    if true {
        v9 -= v9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
