package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint32; F1 int32 }
func genFn0() int16 {
    v6 := []int{32767, -1}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint64 = 0
        v9 := GenS1{F0: 249, F1: 104}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v6
    return -32768
}

func main() {
    var v1 uint16 = 608
    v2 := make(map[uint32]bool)
    var v3 int = 349
    v4 := make(map[int]int)
    var v5 uint64 = 18446744073709551615
    v5 += v5
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint16 = 255
        fmt.Println(v1)
        var v12 int8 = 100
        _ = fi10
        _ = v11
        _ = v12
    }
    v1 -= (v1 + v1)
    for fi13 := 0; fi13 < 3; fi13++ {
        if (false && (!false)) {
            v14 := GenS2{F0: "", F1: ((!(42 <= 4294967295)) && (false && true))}
            v1 = 663
            _ = v14
        } else {
            _ = len(v2)
            var v15 uint = 0
            v15 = (v15 >> ((v15) & 63))
            _ = v15
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            switch "a" {
            case "\n\t":
                if false {
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(fi13, v1)
                    }
                    v3 = (fi13 - (v3 * 4294967295))
                    v17 := GenS2{F0: "", F1: (false || (!((42 <= 1) && (127 == 563))))}
                    _ = v17
                }
            case "":
                v18 := GenS3{F0: 846, F1: 2147483647}
                _ = v18
            default:
                v19 := GenS1{F0: 4294967295, F1: 127}
                _ = v19
            }
            _ = fi16
        }
        _ = fi13
    }
    fmt.Println(v3)
    fmt.Println(v3)
    v1 = v1
    if false {
        var v20 int64 = 1
        _ = v20
    } else {
        var v21 bool = (((1e10 * ((((0.0 + (1e10 / ((0.1 + 1e10) / 0.1))) + -2.25) + 0.0) / 0.1)) + -2.25) < (3.14159 * -2.25))
        var v22 uint16 = (^(v1))
        switch ("a" + "foo bar") {
        case "":
            var v23 bool = ((v5 & (v5 / (((v5 % ((18446744073709551615) | 1))) | 1))) == v5)
            _ = v23
        }
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
