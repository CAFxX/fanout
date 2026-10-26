package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int; F2 int64 }
func genFn0() uint8 {
    var v5 uint64 = 841
    if (1e10 < 3.14159) {
        var v6 bool = (!true)
        switch -606 {
        case 32767:
            v6 = (((!(1 >= 0)) || (v6 || (2147483647 < -649))) || (0 != 949))
        default:
            var v7 uint = 255
            _ = v7
        }
        _ = v6
    }
    v5 = v5
    _ = v5
    return 145
}

func main() {
    var v1 uint = 42
    v2 := []int{215, 210}
    v3 := []int{}
    var v4 float32 = (0.0 * 1e10)
    if true {
        switch "\n\t" {
        case "\n\t":
            switch 128 {
            case 42:
                v4 -= v4
            default:
                v1 -= 18446744073709551615
            }
        case "foo bar":
            v4 += (float32(0.0) / (1.5 - 1e10))
        default:
            if true {
                v4 = 1.5
                var v8 float32 = v4
                fmt.Println(v8, v8, v4)
                _ = v8
            } else {
                if true {
                    var v9 uint = v1
                    if false {
                        v4 -= v4
                        for fi10 := 0; fi10 < 17; fi10++ {
                            var v11 int8 = 114
                            _ = fi10
                            _ = v11
                        }
                    }
                    v12 := GenS1{F0: 18446744073709551615}
                    _ = v12
                    _ = v9
                } else {
                    v1 = 2
                }
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v4 = 1e10
    }
    fmt.Println(v1, v4)
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int = (fi14 / (((-(fi14))) | 1))
            _ = fi14
            _ = v15
        }
        _ = len(v2)
        _ = fi13
    }
    v3 = append(v3, -766)
    for fi16 := 0; fi16 < 10; fi16++ {
        v1 += (v1 / ((133) | 1))
        _ = fi16
    }
    fmt.Println(v4, v4)
    var v17 int8 = -83
    switch 0 {
    case 4294967295:
        v18 := GenS2{F0: "foo bar", F1: 9223372036854775807, F2: -410}
        _ = v18
    case -947:
        v17 = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
