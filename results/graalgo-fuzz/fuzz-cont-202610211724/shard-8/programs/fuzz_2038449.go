package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint8; F1 int16 }

func main() {
    var v1 int16 = -98
    var v2 float64 = (1e10 + -2.25)
    var v3 float64 = (((v2 / (((v2 - v2) - v2) / (((1e10 + (v2 + (v2 + -2.25))) - v2) + v2))) * (1e10 + (v2 * 0.1))) * 0.1)
    v1 -= 0
    for fi4 := 0; fi4 < 13; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 string = "a"
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: -476, F1: true}
        _ = fi4
        _ = v7
    }
    v3 = (v3 * 1.5)
    v8 := []int{32767}
    switch 32768 {
    case 0:
        if true {
            switch 798 {
            case 584:
                var v9 int = 127
                _ = v9
            case 0:
                v3 -= v2
            case 4294967295:
                if true {
                    v10 := GenS1{F0: -424, F1: ((^(v1)) != (v1 + (v1 * 1000)))}
                    _ = v10
                }
            default:
                _ = len(v8)
            }
        }
    default:
        var v11 uint16 = 65535
        _ = v11
    }
    var v12 int16 = v1
    fmt.Println(v3)
    switch "hello" {
    case "":
        if false {
            v1 -= (^(v1))
        }
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 string = ""
        fmt.Println(v2, v3)
        var v15 int16 = (v12 - v1)
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: false}
    v3 = (v3 * 0.0)
    v17 := GenS2{F0: 0, F1: (v1 % ((824) | 1))}
    if true {
        if (-1 == -2147483648) {
            fmt.Println(v12)
            var v18 uint32 = 2147483647
            _ = v18
        } else {
            var v19 uint64 = 2
            _ = v19
        }
        var v20 float64 = 0.1
        _ = v20
    } else {
        v3 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
