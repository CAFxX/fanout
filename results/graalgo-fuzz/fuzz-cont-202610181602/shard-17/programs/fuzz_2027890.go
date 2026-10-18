package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int }
func genFn0(p0 uint64) uint32 {
    v4 := []int{}
    p0 += 0
    v5 := GenS1{F0: -1, F1: 65535}
    _ = p0
    _ = v4
    _ = v5
    return 2147483648
}

func main() {
    var v1 string = "foo bar"
    v2 := []int{}
    var v3 uint32 = 0
    var v6 uint = 0
    if (611 < 932) {
        v1 = v1
        v7 := GenS1{F0: 100, F1: 2}
        for fi8 := 0; fi8 < 7; fi8++ {
            if len(v2) > 0 {
                v2[0] = fi8
            }
            if true {
                fmt.Println(v1, v3)
            } else {
                v6 = v6
            }
            _ = len(v2)
            _ = fi8
        }
        _ = v7
    } else {
        v3 -= (v3 + v3)
    }
    v9 := GenS1{F0: 1, F1: 65535}
    var v10 string = (v1 + "")
    switch -411 {
    case -240:
        v11 := GenS2{F0: 42}
        _ = v11
    case 587:
        var v12 uint8 = 10
        _ = v12
    default:
        if (-128 > -128) {
            var v13 int16 = 351
            v3 += (v3 * v3)
            _ = v13
        }
    }
    var v14 uint8 = 2
    var v15 int8 = 100
    v15 = (-(v15))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
