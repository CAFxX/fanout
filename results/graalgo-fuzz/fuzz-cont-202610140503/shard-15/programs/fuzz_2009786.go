package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint32 }
func genFn0() uint8 {
    var v5 float64 = 0.0
    var v6 int32 = -2147483648
    var v7 uint64 = 274
    if false {
        v8 := []int{42}
        v6 -= v6
        _ = v8
    }
    v9 := GenS1{F0: -1}
    _ = v5
    _ = v6
    _ = v7
    _ = v9
    return 127
}

func main() {
    var v1 int8 = 5
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    v4 := []string{"a", "a", "\n\t", "foo bar"}
    for fi10 := 0; fi10 < 19; fi10++ {
        switch fi10 {
        case 1:
            delete(v3, "")
        case -1:
            fmt.Println(fi10, fi10)
        default:
            var v11 float32 = (1e10 - 0.0)
            _ = v11
        }
        switch ("" + "foo bar") {
        case "hello":
            if ((!false) && true) {
                delete(v3, "a")
            } else {
                v1 -= (v1 << ((2) & 63))
                if true {
                    v1 = v1
                    fmt.Println(v1, fi10)
                    if ((float64((0.1 / 1.0)) * ((-2.25 - 1e10) / 1e10)) < -2.25) {
                        var v12 uint8 = 16
                        for fi13 := 0; fi13 < 4; fi13++ {
                            v1 += (v1 ^ v1)
                            _ = fi13
                        }
                        _ = v12
                    }
                }
            }
        case "foo bar":
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        }
        _ = fi10
    }
    v1 = (v1 << ((18446744073709551615) & 63))
    v1 = v1
    v14 := GenS2{F0: 8, F1: 0}
    for fi15 := 0; fi15 < 9; fi15++ {
        _ = len(v3)
        _ = fi15
    }
    switch "a" {
    case "hello":
        var v16 float64 = 0.0
        _ = v16
    case "":
        v1 += (-(v1))
    case "foo bar":
        _ = len(v3)
    default:
        var v17 float64 = -2.25
        _ = v17
    }
    var v18 uint32 = 42
    if false {
        var v19 float32 = 1e10
        fmt.Println(v19, v18)
        _ = v19
    }
    v18 = (v18 / (((v18 | (v18 % ((2147483648) | 1)))) | 1))
    switch 32768 {
    case 1:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 42:
        var v20 uint16 = 65535
        _ = v20
    default:
        if false {
            var v21 int = 127
            _ = v21
        } else {
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 uint = 4294967295
                fmt.Println(v23)
                _ = fi22
                _ = v23
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
