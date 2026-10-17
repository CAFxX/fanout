package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int8 }
func genFn0() int {
    var v4 uint32 = 2147483647
    fmt.Println(v4)
    v5 := []int{}
    v4 -= v4
    _ = v4
    _ = v5
    return -682
}
func genFn1(p0 int, p1 int64, p2 string) int16 {
    switch p2 {
    case "":
        p2 = (p2 + p2)
    case "foo bar":
        var v6 int = -544
        _ = v6
    }
    fmt.Println(p2, p0, p1)
    if false {
        v7 := []int{(p0 | (p0 % (((^(p0))) | 1))), p0, (p0 | p0)}
        v8 := GenS1{F0: 808, F1: 32767, F2: 65535}
        if (p2 != p2) {
            p0 = genFn0()
        }
        _ = v7
        _ = v8
    } else {
        switch genFn0() {
        case 32767:
            v9 := GenS1{F0: 1000, F1: 381, F2: 2147483647}
            _ = v9
        case -2147483648:
            var v10 float64 = 1e10
            _ = v10
        default:
            p1 += -601
        }
        v11 := GenS2{F0: 1}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 uint32 = 57
    var v2 uint16 = 0
    var v3 int = 4294967295
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 uint = 2
        var v14 uint64 = 4294967295
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 string = ("" + "\n\t")
    v2 = v2
    v2 = (v2 + 32767)
    v3 = (v3 << ((698) & 63))
    v2 = 0
    fmt.Println(v3, v1)
    var v16 int16 = -695
    fmt.Println(v1)
    var v17 uint = 833
    var v18 uint16 = v2
    v19 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
