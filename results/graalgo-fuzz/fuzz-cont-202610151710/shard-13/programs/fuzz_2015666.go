package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint64 }

func main() {
    var v1 string = "foo bar"
    var v2 uint = 0
    var v3 uint16 = 32767
    v4 := make(map[uint32]uint32)
    v5 := []int{-875}
    var v6 uint8 = 255
    fmt.Println(v1, v2)
    _ = len(v5)
    v7 := GenS1{F0: ((42 == 4294967295) && false), F1: 0, F2: 0}
    switch 2147483647 {
    case 1:
        v3 = (v3 & 42)
    default:
        if true {
            v8 := GenS1{F0: false, F1: 888, F2: 2}
            _ = v8
        } else {
            fmt.Println(v3, v2, v2)
        }
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            if false {
                if ((-844 != 1000) && (false || (false && true))) {
                    v4[2] = 0
                } else {
                    v6 -= v6
                }
                v5 = append(v5, 127)
            }
            _ = fi10
        }
        _ = fi9
    }
    v11 := GenS1{F0: true, F1: -2147483648, F2: 437}
    var v12 uint32 = 129
    var v13 float32 = 0.1
    v6 -= v6
    v3 = (v3 * v3)
    switch 2147483647 {
    case -9223372036854775808:
        var v14 int8 = -79
        _ = v14
    default:
        var v15 string = ("foo bar" + v1)
        _ = v15
    }
    v13 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
