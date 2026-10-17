package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS2 struct { F0 int; F1 uint; F2 uint64 }
func genFn0() uint16 {
    var v4 uint64 = 18446744073709551615
    v5 := []int{128, 1, 725}
    var v6 int16 = 42
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "a"
    var v3 string = v2
    v7 := GenS1{F0: 2147483648, F1: 1, F2: -577}
    if true {
        switch "foo bar" {
        case "a":
            delete(v1, 125)
        default:
            v1[-404] = 252
        }
        v2 = v2
    } else {
        if (!false) {
            var v8 uint8 = 255
            if true {
                v3 = (((v2 + "foo bar") + v2) + "")
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 int64 = 0
                    _ = fi9
                    _ = v10
                }
            }
            _ = v8
        }
        var v11 float64 = 1.5
        _ = v11
    }
    var v12 int = 661
    switch ((v2 + "") + v2) {
    case "a":
        v12 -= (v12 & (v12 / ((v12) | 1)))
    case "hello":
        if true {
            if false {
                if ("foo bar" != "foo bar") {
                    var v13 uint = 42
                    _ = v13
                } else {
                    fmt.Println(v12, v3)
                }
            }
        } else {
            v14 := []int{(v12 - (v12 ^ 0))}
            v2 = v2
            _ = v14
        }
    default:
        v15 := GenS2{F0: v12, F1: 1, F2: 18446744073709551615}
        _ = v15
    }
    v12 = v12
    var v16 int32 = 1
    _ = len(v1)
    var v17 int32 = -406
    switch 1 {
    case -381:
        switch v2 {
        case "hello":
            v18 := GenS2{F0: 2147483647, F1: 42, F2: 492}
            _ = v18
        default:
            var v19 uint16 = 986
            _ = v19
        }
    case 0:
        var v20 uint = 0
        _ = v20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        v12 = fi21
        v17 += (-(v16))
        _ = fi21
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
}
