package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 bool = ((0 == 1) || (4294967295 <= 442))
    var v2 bool = false
    v3 := []int{}
    var v4 float32 = 0.0
    var v5 uint64 = 2
    v6 := []int{-90}
    v7 := GenS1{F0: 128}
    v2 = (-217 != 2147483647)
    switch ("" + "foo bar") {
    case "":
        if (334 > -32768) {
            v6 = append(v6, 0)
            if len(v3) > 0 {
                v3[0] = 1
            }
            v2 = true
        } else {
            for fi8 := 0; fi8 < 8; fi8++ {
                v9 := GenS1{F0: 1}
                fmt.Println(v1, v2)
                v2 = (379 < 528)
                _ = fi8
                _ = v9
            }
        }
    case "hello":
        v10 := GenS2{F0: 32767, F1: 128}
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v1, fi11)
        _ = fi11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int = 32767
        _ = fi12
        _ = v13
    }
    var v14 uint64 = v5
    _ = len(v3)
    _ = len(v6)
    v14 += (v5 ^ v5)
    v1 = (!(89 < 1))
    if false {
        v14 -= 4294967295
    } else {
        var v15 int8 = -59
        _ = v15
    }
    switch -822 {
    case -820:
        v4 = v4
    case -1:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        var v16 int = -2147483648
        _ = v16
    }
    var v17 uint8 = 128
    v5 -= (v5 % ((v14) | 1))
    switch -554 {
    case 128:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    var v18 int = 42
    var v19 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
